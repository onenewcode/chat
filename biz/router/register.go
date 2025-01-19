package router

import (
	"chat/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRouter(h *server.Hertz) {
	// 注册测试 swagger 路由
	h.GET("/ping", handler.PingHandler)

	// TODO 未测试
	//静态资源
	h.Static("/asset", ".")
	// 为单个文件提供映射
	h.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	//	设置 html 模板路径
	h.LoadHTMLGlob("views/**/*")
}
