package vo

// 用户登陆视图
type UserLoginVo struct {
	Name     string `json:"name" form:"name"  vd:"($!='')"`
	PassWord string `json:"passWord" form:"password"  vd:"($!='')"`
}
type UserRegisterVo struct {
	Name     string `form:"name"  vd:"($!='')"`
	PassWord string `form:"password"  vd:"($!='')"`
	Identity string `form:"Identity"  vd:"($!='')"`
}

// 添加好友时用的视图
type FriendVo struct {
	UserId     uint   `json:"userId,omitempty"`
	TargetName string `json:"targetName,omitempty"`
}
