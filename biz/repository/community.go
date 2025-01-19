package repository

import (
	"chat/biz/domain"
	"chat/common"
	"chat/utils"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type communityRepo struct {
	db *gorm.DB // 使用构造函数注入 DB 依赖
}

func (c communityRepo) Create(ctx context.Context, community domain.Community) error {
	tx := c.db.WithContext(ctx).Begin()
	//事务一旦开始，不论什么异常最终都会 Rollback
	defer func() {
		// 防止出现 panic
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := c.db.Create(&community).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return fmt.Errorf(common.DBError, err)
	}
	contact := domain.Contact{}
	contact.OwnerId = community.OwnerId
	contact.TargetId = community.ID
	contact.Type = 2 //群关系
	if err := utils.DB.Create(&contact).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf(common.DBError, err)
	}

	tx.Commit()
	return nil
}

// TODO
func (c communityRepo) Load(ctx context.Context, ownerId uint) (*[]domain.Community, error) {
	db := c.db.WithContext(ctx)
	contacts := make([]domain.Contact, 0)
	objIds := make([]uint64, 0)
	db.Where("owner_id = ? and type=2", ownerId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}

	data := make([]domain.Community, 10)
	db.Where("id in ?", objIds).Find(&data)
	return &data, nil
}
