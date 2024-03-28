package main

import (
	"chat/config"
	"chat/router"
	"chat/utils"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func init() {
	config.InitConfig()
}
func main() {
	// 初始化一些必要项目
	utils.InitMySQL()
	utils.InitRedis()
	h := server.New(
		server.WithHostPorts(config.Port.Server),
	)
	router.Router(h)
	h.Use(recovery.Recovery()) // 可确保即使在处理请求过程中发生未预期的错误或异常，服务也能维持运行状态
	h.Spin()
}
