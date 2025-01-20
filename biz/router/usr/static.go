package usr

import (
	"chat/models"
	"context"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

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
