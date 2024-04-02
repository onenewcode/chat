# 工作流程
## 模板设置
所访问的界面均由hertz框架的加载生成他们的工作流程由以下代码实现
```go
	//静态资源
h.Static("/asset", ".")
// 为单个文件提供映射
h.StaticFile("/favicon.ico", "asset/images/favicon.ico")
//	r.StaticFS()
h.LoadHTMLGlob("views/**/*")
```
## 用户基础
### 静态界面
```go
		h.GET("/", service.GetIndex) // 主页
		h.GET("/index", service.GetIndex) // 主页
		h.GET("/toRegister", service.ToRegister) //注册界面
```
### 基础功能
用户的基础功能为用户信息的crud
**代码如下**
```go
		users.POST("/getUserList", service.GetUserList)                   // 获取所有用户
users.POST("/createUser", service.CreateUser)                     //创建新用户
users.POST("/deleteUser", service.DeleteUser)                     // 删除用户
users.POST("/findUserByNameAndPwd", service.FindUserByNameAndPwd) //根据用户名查找用户
users.POST("/updateUser", service.UpdateUser)                     //更新用户数据
users.POST("/find", service.FindByID)                             // 根据用户id查找用户
```
因为是简单的crud就不再对代码进行赘述。
### 核心功能
作为我们的MI系统，聊天便是我们的核心功能，接下来我们便介绍我们的核心代码。
#### SendMsg
函数详细结构，函数主要用于聊天协议的升级，把http升级为websocket，升级后的连接交给MsgHandler函数来处理收发的信息。                
```go
func SendMsg(ctx context.Context, c *app.RequestContext) {
	err := upGrader.Upgrade(c, func(conn *websocket.Conn) {
		// 关闭连接
		defer func(ws *websocket.Conn) {
			err := ws.Close()
			if err != nil {
				hlog.Info(err)
			}
			hlog.Info("websocket关闭")
		}(conn)
		MsgHandler(ctx, conn)
	})
	if err != nil {
		hlog.Info(err)
		return
	}

}
```
#### MsgHandler
函数详细描写如下，他们首先会用utils.Subscribe使用redis的发布订阅模型，利用redis进行用户之间的通信。该模型属于拉去模型，读扩散，之后再通过websocket把数据，推送给前端。
```go
func MsgHandler(ctx context.Context, ws *websocket.Conn) {
	for {
		msg, err := utils.Subscribe(ctx, utils.PublishKey)
		if err != nil {
			hlog.Info(" MsgHandler 接受失败", err)
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			hlog.Error(err)
		}
	}
}
```
## 聊天核心
### chat
**详细解释**
我们在接受请求之后，构建一个node节点，clientMap维护了一个map包含了所有上线交流的用户，接着便会启动两个协程，分别接受和发送信息，同时把这个用户上线的信息，存入redis。
```go
// 需要 ：发送者ID ，接受者ID ，消息类型，发送的内容，发送类型
func Chat(c *app.RequestContext) {
	//1.  获取参数 并 检验 token 等合法性
	//token := query.Get("token")
	Id := c.Query("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	//msgType := query.Get("type")
	//targetId := query.Get("targetId")
	//context := query.Get("context")
	isvalida := true //checkToke()  待补充进行校验
	err := (&websocket.HertzUpgrader{
		//token 校验
		CheckOrigin: func(ctx *app.RequestContext) bool {
			return isvalida
		},
	}).Upgrade(c, func(conn *websocket.Conn) {
		// 关闭连接
		defer func(ws *websocket.Conn) {
			err := ws.Close()
			if err != nil {
				hlog.Info(err)
			}
			hlog.Info("websocket关闭")
		}(conn)
		//2.获取conn
		currentTime := uint64(time.Now().Unix())
		node := &Node{
			Conn:          conn,
			Addr:          conn.RemoteAddr().String(), //客户端地址
			HeartbeatTime: currentTime,                //心跳时间
			LoginTime:     currentTime,                //登录时间
			DataQueue:     make(chan []byte, 50),
			GroupSets:     mapset.NewSet(9),
		}
		//3. 用户关系 todo
		//4. userid 跟 node绑定 并加锁
		rwLocker.Lock()
		clientMap[userId] = node
		rwLocker.Unlock()
		//5.完成发送逻辑
		go sendProc(node)
		//6.完成接受逻辑
		go recvProc(node)
		//7.加入在线用户到缓存
		SetUserOnlineInfo("online_"+Id, []byte(node.Addr), time.Duration(viper.GetInt("timeout.RedisOnlineTime"))*time.Hour)
        sendMsg(userId, []byte("欢迎进入聊天系统"))
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
```
### sendProc
函数详解，通过连接的map集合，获取要发送人，的websocket连接，把信息直接发送给指定前端。
```go
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			hlog.Info("[ws]sendProc >>>> msg :", string(data))
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				hlog.Info(err)
				return
			}
		}
	}
}
```
### recvProc
首先我们一直监听连接，不断地从websocket连接中读取数据。
```go
func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			hlog.Info(err)
			return
		}
		msg := Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			hlog.Info(err)
		}
		// 类型为心跳
		if msg.Type == 3 {
			// 跟新心跳时间
			currentTime := uint64(time.Now().Unix())
			node.Heartbeat(currentTime)
		} else {
			dispatch(data)
			broadMsg(data) //todo 将消息广播到局域网
			hlog.Info("[ws] recvProc <<<<< ", string(data))
		}

	}
}
```
### 后端调度策略
```go

func broadMsg(data []byte) {
	udpsendChan <- data
}

// 后端调度逻辑处理
func dispatch(data []byte) {
msg := Message{}
msg.CreateTime = uint64(time.Now().Unix())
err := json.Unmarshal(data, &msg)
if err != nil {
hlog.Info(err)
return
}
switch msg.Type {
case 1: //私信
hlog.Info("dispatch  data :", string(data))
sendMsg(msg.TargetId, data)
case 2: //群发
sendGroupMsg(msg.TargetId, data) //发送的群ID ，消息内容
}
}
```
### 群聊功能

