package router

import (
	"chat/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRouter(h *server.Hertz) {
	// 注册测试swagger路由
	h.GET("/ping", handler.PingHandler)

}
