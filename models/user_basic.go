package models

import (
	"chat/utils"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model    `json:"gorm.Model"`
	Name          string    `json:"name" vd:"($!='')"`
	PassWord      string    `json:"passWord" vd:"($!='')"`
	Phone         string    `json:"phone" vd:"regexp('(^1[3-9]{1}\\d{9}$')"`
	Email         string    `json:"email" vd:"email"`
	Avatar        string    `json:"avatar,omitempty"` //头像
	Identity      string    `json:"identity,omitempty" vd:"($!='')"`
	ClientIp      string    `json:"clientIp,omitempty"`
	ClientPort    string    `json:"clientPort,omitempty"`
	Salt          string    `json:"salt,omitempty"`
	LoginTime     time.Time `json:"loginTime"`
	HeartbeatTime time.Time `json:"heartbeatTime"`
	LoginOutTime  time.Time `json:"loginOutTime" gorm:"column:login_out_time" json:"login_out_time" `
	IsLogout      bool      `json:"isLogout,omitempty"`
	DeviceInfo    string    `json:"deviceInfo,omitempty"`
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
	return utils.DB.Create(&user)
}
func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email, Avatar: user.Avatar})
}

// 查找某个用户
func FindByID(id uint) UserBasic {
	user := UserBasic{}
	utils.DB.Where("id = ?", id).First(&user)
	return user
}
