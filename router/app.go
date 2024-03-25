package router

import (
	"chat/service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Router(h *server.Hertz) {
	//swagger
	//docs.SwaggerInfo.BasePath = ""
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//静态资源
	h.Static("/asset", ".")
	// 为单个文件提供映射
	h.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	//	r.StaticFS()
	h.LoadHTMLGlob("views/**/*")
	//首页
	{
		h.GET("/", service.GetIndex)
		h.GET("/index", service.GetIndex)
		h.GET("/toRegister", service.ToRegister)
	}
	users := h.Group("/user")
	{
		users.POST("/getUserList", service.GetUserList)
		users.POST("/createUser", service.CreateUser)
		users.POST("/deleteUser", service.DeleteUser)
	}
}
