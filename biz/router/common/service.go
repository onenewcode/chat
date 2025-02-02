package common

import (
	"chat/common"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"path"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func Upload(ctx context.Context, c *app.RequestContext) {
	uploadLocal(c)
}

// 上传文件到本地
func uploadLocal(c *app.RequestContext) {
	head, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, common.H{Code: -1, Data: nil, Msg: err.Error()})
	}
	// 获取文件拓展名
	suffix := path.Ext(head.Filename)
	// 生成随机文件名称
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	err = c.SaveUploadedFile(head, "./asset/upload/"+fileName)
	if err != nil {
		c.JSON(http.StatusOK, common.H{Code: -1, Data: nil, Msg: err.Error()})
	}
	url := "./asset/upload/" + fileName
	c.JSON(http.StatusOK, common.H{Code: 0,
		Data: url,
		Msg:  "发送图片成功"})
}
