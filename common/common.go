package common

const (
	UserNameExist            = "用户名已注册！"
	UserNamePassWordISEmpty  = "用户名或密码不能为空！"
	UserPasswordInconsistent = "两次密码不一致！"
	UserDeletedSucceed       = "该用户不存在"
	User
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
