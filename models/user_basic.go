package models

import (
	"chat/utils"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string    `json:"name,omitempty" form:"name" gorm:"name,type:;not null;"`
	PassWord      string    `json:"passWord,omitempty" form:"password" gorm:"pass_word,type:;not null;"`
	Phone         string    `json:"phone,omitempty" gorm:"phone,type:;not null;"` // vd:"regexp('^1[3-9]{1}\\d{9}$')"
	Email         string    `json:"email,omitempty" gorm:"email,type:;not null;"`
	Avatar        string    `json:"avatar,omitempty" gorm:"avatar,type:;not null;"` //头像
	Identity      string    `form:"Identity" json:"identity,omitempty" gorm:"identity,type:;not null;"`
	ClientIp      string    `json:"clientIp,omitempty" gorm:"client_ip,type:;not null;"`
	ClientPort    string    `json:"clientPort,omitempty" gorm:"client_port,type:;not null;"`
	Salt          string    `json:"salt,omitempty" gorm:"salt,type:;not null;"`
	LoginTime     time.Time `json:"loginTime,omitempty" gorm:"login_time,type:;not null;"`
	HeartbeatTime time.Time `json:"heartbeatTime,omitempty" gorm:"heartbeat_time,type:;not null;"`
	LoginOutTime  time.Time `json:"loginOutTime,omitempty" gorm:"column:login_out_time,type:;not null;"`
	IsLogout      bool      `json:"isLogout,omitempty" gorm:"is_logout,type:;not null;"`
	DeviceInfo    string    `json:"deviceInfo,omitempty" gorm:"device_info,type:;not null;"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

// 用户列表查询
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		hlog.Info(v)
	}
	return data
}

// 用作用户登陆
func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and pass_word=?", name, password).First(&user)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}
func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("Phone = ?", phone).First(&user)
}
func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}
func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Select("*").Create(&user)
}
func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(&user)
}

// 查找某个用户
func FindByID(id uint) UserBasic {
	user := UserBasic{}
	utils.DB.Where("id = ?", id).First(&user)
	return user
}
