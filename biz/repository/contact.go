package repository

import (
	"chat/biz/domain"
	"context"

	"gorm.io/gorm"
)

type ContactRepo struct {
	db *gorm.DB
}

func (c ContactRepo) SearchFriend(ctx context.Context, userId uint) *[]domain.UserBasic {
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

// TODO
