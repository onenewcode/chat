package models

import "gorm.io/gorm"

// 群信息
type GroupBasic struct {
	gorm.Model
	Name    string `json:"name,omitempty" gorm:"name,type:;not null;"`
	OwnerId uint   `json:"owner_id,omitempty" gorm:"owner_id,type:;not null;"`
	Icon    string `json:"icon,omitempty" gorm:"icon,type:;not null;"`
	Type    int    `json:"type,omitempty" gorm:"type,type:;not null;"`
	Desc    string `json:"desc,omitempty" gorm:"desc,type:;not null;"`
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
