package main

import (
	"chat/biz/router"
	"chat/config"
	_ "chat/docs" // swagger docs
	"chat/models"
	"chat/utils"
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-redis/redis/v8"
	"github.com/hertz-contrib/cache/persist"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/swagger"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	"gorm.io/plugin/opentelemetry/provider"
)

func initALL() {
	Initialize()
	// 初始化定时器，用于定期清除过期的node连接
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
	// 初始化缓存中间件
	utils.RedisStore = persist.NewRedisStore(redis.NewClient(&redis.Options{
		Addr: config.GlobalConfig.Redis.Addr,
	}))
	err := utils.RedisStore.RedisClient.Ping(context.Background()).Err()
	if err != nil {
		fmt.Printf("Redis connection failed: %v\n", err)
		return
	}

}
func initOpentelemetry(serviceName string) {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		// Support setting ExportEndpoint via environment variables: OTEL_EXPORTER_OTLP_ENDPOINT
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
}
func initSwagger(h *server.Hertz) {
	// 教程地址：https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/middleware/swagger/#%E4%BD%BF%E7%94%A8%E7%94%A8%E6%B3%95
	url := swagger.URL("http://localhost:8888/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
}

// @title chat
// @version 1.0
// @description 这是一个用于hertz的聊天服务

// @contact.name onenewcode
// @contact.url  https://github.com/onenewcode/chat.git

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8888
// @BasePath /
// @schemes http
func main() {
	// initOpentelemetry("chat")
	// 初始化链路追踪
	tracer, cfg := hertztracing.NewServerTracer()

	h := server.New(
		server.WithHostPorts(config.Port.Server),
		tracer,
	)
	h.Use(hertztracing.ServerMiddleware(cfg))
	//

	// 注册路由
	router.RegisterRouter(h)
	h.Use(recovery.Recovery()) // 可确保即使在处理请求过程中发生未预期的错误或异常，服务也能维持运行状态
	h.Spin()
}
