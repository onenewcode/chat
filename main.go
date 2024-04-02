package main

import (
	"chat/config"
	"chat/models"
	"chat/router"
	"chat/utils"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/viper"
	"time"
)

func init() {
	config.InitConfig()
	// 初始化定时器，用于定期清除过期的node连接
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
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
