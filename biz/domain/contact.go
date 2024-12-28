package domain

import (
	"context"

	"gorm.io/gorm"
)

// 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint   `json:"owner_id,omitempty" gorm:"owner_id,type:;not null;" valid:"owner_id"`    //谁的关系信息
	TargetId uint   `json:"target_id,omitempty" gorm:"target_id,type:;not null;" valid:"target_id"` //对应的谁 /群 ID
	Type     int    `json:"type,omitempty" gorm:"type,type:;not null;" valid:"type"`                //对应的类型  1好友  2群  3xx
	Desc     string `json:"desc,omitempty" gorm:"desc,type:;not null;" valid:"desc"`
}

func (table *Contact) TableName() string {
	return "contact"
}

type ContactI interface {
	SearchFriend(ctx context.Context, userId uint) []UserBasic
	// 添加好友   自己的ID  ， 好友的ID
	AddFriend(ctx context.Context, userId uint, targetName string) (int, string)

	SearchUserByGroupId(ctx context.Context, communityId uint) []uint
}
