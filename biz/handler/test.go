package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// PingHandler 测试 handler
// @Summary 测试 Summary
// @Description 测试 Description
// @Accept application/json
// @Produce application/json
// @Router /ping [get]
func PingHandler(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, map[string]string{
		"ping": "pong",
	})
}
