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
		h.GET("/", service.GetIndex)             // 主页
		h.GET("/index", service.GetIndex)        // 主页与"/"显示的界面相同
		h.GET("/toRegister", service.ToRegister) // 注册界面
		h.GET("/toChat", service.ToChat)         // 聊天主页
		h.GET("/chat", service.Chat)
		// 查找所有好友
		h.POST("/searchFriends", service.SearchFriends)
	}
	users := h.Group("/user")
	{
		users.POST("/getUserList", service.GetUserList)                   // 获取所有用户
		users.POST("/createUser", service.CreateUser)                     //创建新用户
		users.POST("/deleteUser", service.DeleteUser)                     // 删除用户
		users.POST("/findUserByNameAndPwd", service.FindUserByNameAndPwd) //根据用户名查找用户
		users.POST("/updateUser", service.UpdateUser)                     //更新用户数据
		users.POST("/find", service.FindByID)                             // 根据用户id查找用户
		//发送消息 群发
		users.GET("/sendMsg", service.SendMsg)
		//发送消息
		users.GET("/sendUserMsg", service.SendUserMsg)
		// redis中获取消息顺序
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
	h.POST("/attach/upload", service.Upload)
}
