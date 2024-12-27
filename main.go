package main

import (
	"chat/config"
	"chat/models"
	"chat/router"
	"chat/utils"
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-redis/redis/v8"
	"github.com/hertz-contrib/cache/persist"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/spf13/viper"
)

func init() {
	config := Initialize()
	// 初始化定时器，用于定期清除过期的node连接
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
	// 初始化缓存中间件
	utils.RedisStore = persist.NewRedisStore(redis.NewClient(&redis.Options{
		Addr: config.Redis.Addr,
	}))
	err := utils.RedisStore.RedisClient.Ping(context.Background()).Err()
	if err != nil {
		fmt.Printf("Redis connection failed: %v\n", err)
		return
	}

}
func main() {
	// 导出
	//serviceName := "echo"
	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(serviceName),
	//	provider.WithExportEndpoint("localhost:4317"),
	//	provider.WithInsecure(),
	//)
	//defer p.Shutdown(context.Background())

	tracer, cfg := hertztracing.NewServerTracer()

	h := server.New(
		server.WithHostPorts(config.Port.Server),
		tracer,
	)
	h.Use(hertztracing.ServerMiddleware(cfg))

	router.Router(h)
	h.Use(recovery.Recovery()) // 可确保即使在处理请求过程中发生未预期的错误或异常，服务也能维持运行状态
	h.Spin()
}
