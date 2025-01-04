package home

import (
	"chat/service"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterPageRouter(h *server.Hertz) {
	h.GET("/", service.GetIndex)             // 主页
	h.GET("/index", service.GetIndex)        // 主页与"/"显示的界面相同
	h.GET("/toRegister", service.ToRegister) // 注册界面
	h.GET("/toChat", service.ToChat)         // 聊天主页
}
