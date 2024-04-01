package service

import (
	"chat/common"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"math/rand"
	"net/http"
	"path"
	"time"
)

func Upload(ctx context.Context, c *app.RequestContext) {
	UploadLocal(c)
}

// 上传文件到本地
func UploadLocal(c *app.RequestContext) {
	head, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, common.H{Code: -1, Data: nil, Msg: err.Error()})
	}

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

// 上传文件到Miono
func UploadOOS(ctx context.Context, c *app.RequestContext) {
	//w := c.Writer
	//req := c.Request
	//srcFile, head, err := req.FormFile("file")
	//if err != nil {
	//	utils.RespFail(w, err.Error())
	//}
	//suffix := ".png"
	//ofilName := head.Filename
	//tem := strings.Split(ofilName, ".")
	//if len(tem) > 1 {
	//	suffix = "." + tem[len(tem)-1]
	//}
	//fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	//// 创建OSSClient实例。
	//// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	//// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	//client, err := oss.New(viper.GetString("oss.Endpoint"), viper.GetString("oss.AccessKeyId"), viper.GetString("oss.AccessKeySecret"))
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}
	//
	//// 填写存储空间名称，例如examplebucket。
	//bucket, err := client.Bucket(viper.GetString("oss.Bucket"))
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}
	//
	//// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	//err = bucket.PutObject(fileName, srcFile)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}
	//url := "http://" + viper.GetString("oos.Bucket") + "." + viper.GetString("oos.EndPoint") + "/" + fileName
	//utils.RespOK(w, url, "发送图片成功")
}
