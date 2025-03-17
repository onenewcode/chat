package router

import (
	"chat/biz/handler"
	"chat/biz/router/common"
	"chat/biz/router/home"
	"chat/biz/router/usr"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRouter(h *server.Hertz) {
	// 注册测试 swagger 路由
	h.GET("/ping", handler.PingHandler)

	// 设置全局的缓存过期时间（会被更细粒度的设置覆盖）
	// my_cache := cache.NewCacheByRequestURI(utils.RedisStore, 2*time.Hour)

	//静态资源
	h.Static("/asset", ".")
	// 为单个文件提供映射
	h.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	//	r.StaticFS()
	h.LoadHTMLGlob("views/**/*")
	//首页
	{
		h.GET("/", home.GetIndex)             // 主页
		h.GET("/index", home.GetIndex)        // 主页与"/"显示的界面相同
		h.GET("/toRegister", home.ToRegister) // 注册界面
	}
	// 聊天主页
	{
		h.GET("/toChat", usr.ToChat) // 聊天主页
		h.GET("/chat", usr.Chat)
		// 查找所有好友
		h.POST("/searchFriends", usr.SearchFriends)
	}
	users := h.Group("/user")
	{
		users.POST("/login", usr.Login)            //根据用户名查找用户
		users.POST("/getUserList", usr.CreateUser) //创建新用户
		users.POST("/deleteUser", usr.DeleteUser)  // 删除用户

		users.POST("/updateUser", usr.UpdateUser) //更新用户数据
		users.POST("/find", usr.FindByID)         // 根据用户 id 查找用户
		//发送消息 群发
		users.GET("/sendMsg", usr.SendMsg)
		//发送消息
		users.GET("/sendUserMsg", usr.SendUserMsg)
		// // 消息存在 redis 中，现在
		users.POST("/redisMsg", usr.RedisMsg)
	}
	// 群聊信息
	contact := h.Group("/contact")
	{
		//添加好友
		contact.POST("/addfriend", usr.AddFriend)
		//群列表
		contact.POST("/loadcommunity", usr.LoadCommunity)
		// 创建群
		contact.POST("/createCommunity", usr.CreateCommunity)
		// 添加如群
		contact.POST("/joinGroup", usr.JoinGroups)
	}
	// //上传文件
	h.POST("/attach/upload", common.Upload)
}
