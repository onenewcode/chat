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
	UserId     uint   `form:"userId,omitempty" vd:"($!=0)"`
	TargetName string `form:"targetName,omitempty" vd:"($!='')"`
}
type UserUpdate struct {
	ID     uint   `form:"id"`
	Name   string `form:"name"  vd:"($!='')"`
	Avatar string `form:"icon"`
}
