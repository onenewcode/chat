package models

import (
	"chat/utils"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name    string `form:"name,omitempty" vd:"($!='')"`
	OwnerId uint   `form:"ownerId,omitempty" vd:"($!=0)"`
	Img     string `form:"icon,omitempty"`
	Desc    string `form:"desc,omitempty"`
}

func CreateCommunity(community Community) (int, string) {
	tx := utils.DB.Begin()
	//事务一旦开始，不论什么异常最终都会 Rollback
	defer func() {
		// 防止出现panic
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := utils.DB.Create(&community).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return -1, "建群失败"
	}
	contact := Contact{}
	contact.OwnerId = community.OwnerId
	contact.TargetId = community.ID
	contact.Type = 2 //群关系
	if err := utils.DB.Create(&contact).Error; err != nil {
		tx.Rollback()
		return -1, "添加群关系失败"
	}

	tx.Commit()
	return 0, "建群成功"

}

func LoadCommunity(ownerId uint) ([]*Community, string) {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type=2", ownerId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}

	data := make([]*Community, 10)
	utils.DB.Where("id in ?", objIds).Find(&data)
	for _, v := range data {
		hlog.Info(v)
	}
	//utils.DB.Where()
	return data, "查询成功"
}
