package main

import (
	"chat/config"
	"chat/router"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func init() {
	config.InitConfig()
}
func main() {
	h := server.New(
		server.WithHostPorts(config.Port.Server),
	)
	router.Router(h)
	h.Use(recovery.Recovery()) // 可确保即使在处理请求过程中发生未预期的错误或异常，服务也能维持运行状态
	h.Spin()
}
