package common

const (
	UserNameExist            = "用户名已注册！"
	UserNamePassWordISEmpty  = "用户名或密码不能为空！"
	UserPasswordInconsistent = "两次密码不一致！"
	UserISEmpty              = "该用户不存在"
	UserPasswordError        = "用户密码错误"
	UserParamError           = "修改参数不匹配！"
	UserUpdateSucceed        = "用户修改信息成功"
	UserCantAddYourself      = "不能加自己"
	UserCantAddRepeat        = "不能重复添加"
	UserAddFriendError       = "添加好友失败"
	FindNotUser              = "没有找到此用户"
)
const (
	UserDeletedSucceed   = "用户删除成功"
	UserLoginSucceed     = "用户登录成功"
	UserAddFriendSucceed = "添加好友成功"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
