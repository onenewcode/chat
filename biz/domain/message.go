package domain

import (
	"context"

	"gorm.io/gorm"
)

// 消息
type Message struct {
	gorm.Model
	UserId     int64  `json:"userId,omitempty" gorm:"user_id,type:;not null;"`          //发送者
	TargetId   int64  `json:"target_id,omitempty" gorm:"target_id,type:;not null;"`     //接受者
	Type       int    `json:"type,omitempty" gorm:"type,type:;not null;"`               //发送类型  1私聊  2群聊  3心跳
	Media      int    `json:"media,omitempty" gorm:"media,type:;not null;"`             //消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string `json:"content,omitempty" gorm:"content,type:;not null;"`         //消息内容
	CreateTime uint64 `json:"create_time,omitempty" gorm:"create_time,type:;not null;"` //创建时间
	ReadTime   uint64 `json:"read_time,omitempty" gorm:"read_time,type:;not null;"`     //读取时间
	Pic        string `json:"pic,omitempty" gorm:"pic,type:;not null;"`
	Url        string `json:"url,omitempty" gorm:"url,type:;not null;"`
	Desc       string `json:"desc,omitempty" gorm:"desc,type:;not null;"`
	Amount     int    `json:"amount,omitempty" gorm:"amount,type:;not null;"` //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type MessageI interface {
	// 存储消息
	Save(ctx context.Context, msg Message) error
	// 通过user_id获取消息列表
	ListUserId(ctx context.Context, userId, targetId int64, pageIdx, pageSize int) *[]Message
}
