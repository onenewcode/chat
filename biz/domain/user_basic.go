package domain

import (
	"context"
	"time"

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

type UserBasicI interface {
	// 注意 pageIdx<0 时，返回所有数据
	GetList(ctx context.Context, name, phone, email string, pageIdx, pageSize int) *[]UserBasic
	FindByNameAndPwd(ctx context.Context, name string, password string) UserBasic
	// 非模糊查询一般
	FindByName(ctx context.Context, name string) UserBasic
	FindByPhone(ctx context.Context, phone string) UserBasic
	FindByEmail(ctx context.Context, email string) UserBasic
	Create(ctx context.Context, user UserBasic) error
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, user UserBasic) error
	FindByID(ctx context.Context, id uint) UserBasic
}
