package usr

import (
	"chat/common"
	"chat/utils"
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
)

// 防止跨域站点伪造请求
var upGrader = websocket.HertzUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
}

// 消息处理器
func msgHandler(ctx context.Context, ws *websocket.Conn) {
	// 由前端控制 websocket 的关闭
	for {
		msg, err := utils.Subscribe(ctx, utils.PublishKey)
		if err != nil {
			hlog.Info(" MsgHandler 接受失败", err)
		}
		tm := time.Now().Format(common.DataFormatChinese)
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		// 写入消息
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			hlog.Error(err)
		}
	}
}
func SendMsg(ctx context.Context, c *app.RequestContext) {
	err := upGrader.Upgrade(c, func(conn *websocket.Conn) {
		// 关闭连接
		defer func(ws *websocket.Conn) {
			err := ws.Close()
			if err != nil {
				hlog.Info(err)
			}
			hlog.Info("websocket 关闭")
		}(conn)
		msgHandler(ctx, conn)
	})
	if err != nil {
		hlog.Info(err)
		return
	}
}
