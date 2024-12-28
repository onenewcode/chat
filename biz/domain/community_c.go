package domain

import (
	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name    string `form:"name,omitempty" vd:"($!='')" json:"name,omitempty" gorm:"name,type:;not null;"`
	OwnerId uint   `form:"ownerId,omitempty" vd:"($!=0)" json:"owner_id,omitempty" gorm:"owner_id,type:;not null;"`
	Img     string `form:"icon,omitempty" json:"img,omitempty" gorm:"img,type:;not null;"`
	Desc    string `form:"desc,omitempty" json:"desc,omitempty" gorm:"desc,type:;not null;"`
}

func (table *Community) TableName() string {
	return "communit"
}

type CommunityI interface {
}
