package main

import (
	"chat/config"
	"fmt"
)

func init() {
	config.InitConfig()
}
func main() {
	//h := server.New(
	//	server.WithHostPorts(),
	//	)
	//h.Use(recovery.Recovery()) // 可确保即使在处理请求过程中发生未预期的错误或异常，服务也能维持运行状态
	//h.Spin()
	fmt.Println(config.Mysql.DNS)
}
