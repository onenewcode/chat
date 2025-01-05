package usr

import (
	"chat/biz/handler/usr"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterPageRouter(h *server.Hertz) {
	h.GET("/", usr.GetIndex)             // 主页
	h.GET("/index", usr.GetIndex)        // 主页与"/"显示的界面相同
	h.GET("/toRegister", usr.ToRegister) // 注册界面
	h.GET("/toChat", usr.ToChat)         // 聊天主页
}
