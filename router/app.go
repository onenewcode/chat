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
		h.GET("/toChat", service.ToChat)
		h.GET("/chat", service.Chat)
		// 查找所有好友
		h.POST("/searchFriends", service.SearchFriends)
	}
	users := h.Group("/user")
	{
		users.POST("/getUserList", service.GetUserList)
		users.POST("/createUser", service.CreateUser)
		users.POST("/deleteUser", service.DeleteUser)
		users.POST("/findUserByNameAndPwd", service.FindUserByNameAndPwd)
		users.POST("/updateUser", service.UpdateUser)
		users.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
		users.POST("/find", service.FindByID)
		//发送消息 群发
		users.GET("/sendMsg", service.SendMsg)
		//发送消息
		users.GET("/sendUserMsg", service.SendUserMsg)
		users.POST("/redisMsg", service.RedisMsg)
	}
	// 群聊信息
	contact := h.Group("/contact")
	{
		//添加好友
		contact.POST("/addfriend", service.AddFriend)
		//群列表
		contact.POST("/loadcommunity", service.LoadCommunity)
		// 创建群
		contact.POST("/createCommunity", service.CreateCommunity)
		// 添加如群
		contact.POST("/joinGroup", service.JoinGroups)
	}
	//上传文件
	//h.POST("/attach/upload", service.Upload)
}
