package vo

// 用户登陆视图
type UserLoginVo struct {
	Name     string `json:"name,omitempty" form:"name" vd:"($!='')" gorm:"name,type:;not null;"`
	PassWord string `json:"passWord,omitempty" form:"password" vd:"($!='')" gorm:"pass_word,type:;not null;"`
}
type UserRegisterVo struct {
	Name     string `form:"name" vd:"($!='')" json:"name,omitempty" gorm:"name,type:;not null;"`
	PassWord string `form:"password" vd:"($!='')" json:"pass_word,omitempty" gorm:"pass_word,type:;not null;"`
	Identity string `form:"Identity" vd:"($!='')" json:"identity,omitempty" gorm:"identity,type:;not null;"`
}

// 添加好友时用的视图
type FriendVo struct {
	UserId     uint   `form:"userId,omitempty" vd:"($!=0)" json:"user_id,omitempty" gorm:"user_id,type:;not null;"`
	TargeId    uint   `form:"targeId,omitempty" vd:"($!=0)" json:"targe_id,omitempty" gorm:"targe_id,type:;not null;"`
	TargetName string `form:"targetName,omitempty" vd:"($!='')" json:"target_name,omitempty" gorm:"target_name,type:;not null;"`
}

// 用户更新视图
type UserUpdate struct {
	ID     uint   `form:"id" json:"id,omitempty" gorm:"id,type:;not null;"`
	Name   string `form:"name" vd:"($!='')" json:"name,omitempty" gorm:"name,type:;not null;"`
	Avatar string `form:"icon" json:"avatar,omitempty" gorm:"avatar,type:;not null;"`
}

type AddContactVo struct {
	TargetId uint
	Type     int
}
