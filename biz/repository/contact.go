package repository

import (
	"chat/biz/domain"
	"chat/common"
	"chat/utils"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type contactRepo struct {
	db *gorm.DB
}

func (c contactRepo) SearchFriend(ctx context.Context, userId uint) *[]domain.UserBasic {
	db := c.db.WithContext(ctx)
	contacts := make([]domain.Contact, 0)
	objIds := make([]uint64, 0)
	db.Where("owner_id = ? and type=1", userId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]domain.UserBasic, 0)
	db.Where("id in ?", objIds).Find(&users)
	return &users
}

// 添加好友   自己的 ID，好友的 ID
func (c contactRepo) AddFriend(ctx context.Context, userId, targetId uint) error {
	db := c.db.WithContext(ctx)
	tx := db.Begin()
	//事务一旦开始，不论什么异常最终都会 Rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	contact := domain.Contact{}
	contact.OwnerId = userId
	contact.TargetId = targetId
	contact.Type = 1
	if err := utils.DB.Create(&contact).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf(common.DBError, err)
	}
	contact1 := domain.Contact{}
	contact1.OwnerId = targetId
	contact1.TargetId = userId
	contact1.Type = 1
	if err := utils.DB.Create(&contact1).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf(common.DBError, err)
	}
	tx.Commit()
	return nil

}

func (c contactRepo) SearchUserByGroupId(ctx context.Context, communityId uint, targetId int) []uint {
	contacts := make([]domain.Contact, 0)
	objIds := make([]uint, 0)
	utils.DB.Where("target_id = ? and type= ?", communityId, targetId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint(v.OwnerId))
	}
	return objIds
}

func (c contactRepo) Load(ctx context.Context, ownerId uint) (*[]domain.Community, error) {
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
