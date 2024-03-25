package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(ctx context.Context, c *app.RequestContext) {
	c.HTML(http.StatusOK, "index_index.html", nil)
}
func ToRegister(ctx context.Context, c *app.RequestContext) {
	c.HTML(http.StatusOK, "register.html", nil)
}
