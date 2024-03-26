package common

const (
	UserNameExist            = "用户名已注册！"
	UserNamePassWordISEmpty  = "用户名或密码不能为空！"
	UserPasswordInconsistent = "两次密码不一致！"
	UserDeletedSucceed       = "用户删除成功"
	UserISEmpty              = "该用户不存在"
	UserPasswordError        = "用户密码错误"
	UserLoginSucceed         = "用户登录成功"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
