package repository

import (
	"chat/biz/domain"
	"chat/common"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type MessageRepo struct {
	db *gorm.DB
}

func (m MessageRepo) Save(ctx context.Context, msg domain.Message) error {
	db := m.db.WithContext(ctx)
	if err := db.Create(&msg).Error; err != nil {
		return fmt.Errorf(common.DBError, err)
	}
	return nil
}

func (m MessageRepo) ListUserId(ctx context.Context, userId, targetId int64, pageIdx, pageSize int) *[]domain.Message {
	db := m.db.WithContext(ctx)
	var data []domain.Message
	db.Where("user_id=?", userId).Or("target_id=?", targetId).Limit(pageSize).Offset((pageIdx - 1) * pageSize).Find(&data)
	return &data
}
