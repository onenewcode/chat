package service

import (
	"chat/models"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
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

// 进入聊天主页
func ToChat(ctx context.Context, c *app.RequestContext) {
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token
	c.HTML(http.StatusOK, "index.html", user)
}

func Chat(ctx context.Context, c *app.RequestContext) {
	models.Chat(c)
}
