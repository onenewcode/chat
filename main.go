package main

import (
	"chat/biz/router"
	_ "chat/docs" // swagger docs
	"context"

	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	"gorm.io/plugin/opentelemetry/provider"
)

func initALL() {
	// Initialize()
	// 初始化链路追踪
	// initOpentelemetry("chat")
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

// 访问路径 http://localhost:8888/swagger/index.html#/
func initSwagger(h *server.Hertz) {
	// 教程地址：https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/middleware/swagger/#%E4%BD%BF%E7%94%A8%E7%94%A8%E6%B3%95
	url := swagger.URL("http://localhost:8888/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
}

// @title chat
// @version 1.0
// @description 这是一个用于 hertz 的聊天服务

// @contact.name onenewcode
// @contact.url  https://github.com/onenewcode/chat.git

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8888
// @BasePath /
// @schemes http
func main() {
	initALL()
	tracer, cfg := hertztracing.NewServerTracer()

	h := server.New(
		// server.WithHostPorts(config.Port.Server),
		tracer,
	)
	h.Use(hertztracing.ServerMiddleware(cfg))

	// 注册 swagger
	initSwagger(h)
	//注册路由
	router.RegisterRouter(h)
	h.Use(recovery.Recovery()) // 可确保即使在处理请求过程中发生未预期的错误或异常，服务也能维持运行状态
	h.Spin()
}
