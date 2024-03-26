package vo

// 用户登陆视图
type UserLoginVo struct {
	Name     string `json:"name" form:"name"  vd:"($!='')"`
	PassWord string `json:"passWord" form:"password"  vd:"($!='')"`
}
