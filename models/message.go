package models

import (
	"chat/config"
	"chat/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/deckarep/golang-set/v2"
	"github.com/hertz-contrib/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net"
	"strconv"
	"sync"
	"time"
)

// 消息
type Message struct {
	gorm.Model
	UserId     int64  `json:"userId,omitempty"` //发送者
	TargetId   int64  //接受者
	Type       int    //发送类型  1私聊  2群聊  3心跳
	Media      int    //消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string //消息内容
	CreateTime uint64 //创建时间
	ReadTime   uint64 //读取时间
	Pic        string
	Url        string
	Desc       string
	Amount     int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

// 需要重写此方法才能完整的msg转byte[]
func (msg Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(msg)
}

// 存储消息
func Save(msg Message) {
	utils.DB.Create(&msg)
}

// 通过user_id获取消息列表
func ListUserId(id int64) *[]Message {
	var nums []Message
	utils.DB.Where("user_id=?", id).Or("target_id=?", id).Find(&nums)
	return &nums
}

// 连接节点
type Node struct {
	Conn          *websocket.Conn //连接
	Addr          string          //客户端地址
	FirstTime     uint64          //首次连接时间
	HeartbeatTime uint64          //心跳时间
	LoginTime     uint64          //登录时间
	DataQueue     chan []byte     //消息,从
	GroupSets     mapset.Set[int] //好友 / 群
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁,保障clientMap的并发安全
var rwLocker sync.RWMutex

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
	// 判断是否已有连接，已有的话直接删除，防止重建连接导致程序推出。
	//{
	//	rwLocker.Lock()
	//	clientMap[userId] = node
	//	rwLocker.Unlock()
	//}

	// 升级协议，
	err := (&websocket.HertzUpgrader{
		//token 校验
		CheckOrigin: func(ctx *app.RequestContext) bool {
			return isvalida
		},
	}).Upgrade(c, func(conn *websocket.Conn) {
		//2.获取conn
		currentTime := uint64(time.Now().Unix())
		// 构建我们的连接节点
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
		//if err != nil {
		//	hlog.Info(err)
		//	return
		//}
		//5.完成发送逻辑
		go sendProc(node)
		//6.完成接受逻辑
		go recvProc(node)

		//7.发送历史消息
		//{
		//	nums := ListUserId(userId)
		//	for _, v := range *nums {
		//		b, _ := json.Marshal(&v)
		//		node.DataQueue <- b
		//	}
		//}
		//8.加入在线用户到缓存
		SetUserOnlineInfo("online_"+Id, []byte(node.Addr), time.Duration(config.Timeout.RedisOnlineTime)*time.Hour)
		// 监听，应为一旦升级结束便会关闭websocket连接
		for {
			// 监听集合中是否有，没有就直接推出
			_, flag := clientMap[userId]
			if flag == false {
				return
			}
			time.Sleep(1 * time.Second)
		}
	})
	if err != nil {
		hlog.Info(err)
		return
	}
}

// 发送信息到客户端
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

// 从客户端接收信息
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
			broadMsg(data) // 广播信息
			hlog.Info("[ws] recvProc <<<<< ", string(data))
		}

	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
	hlog.Info("init goroutine ")
}

// 完成udp数据发送协程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 255),
		Port: viper.GetInt("port.udp"),
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case data := <-udpsendChan:
			fmt.Println("udpSendProc  data :", string(data))
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}

// 完成udp数据接收协程
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: viper.GetInt("port.udp"),
	})
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("udpRecvProc  data :", string(buf[0:n]))
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑处理，同时把消息存入数据库
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

// 群发信息
func sendGroupMsg(targetId int64, msg []byte) {
	hlog.Info("开始群发消息")
	userIds := SearchUserByGroupId(uint(targetId))
	for i := 0; i < len(userIds); i++ {
		//排除给自己的
		if targetId != int64(userIds[i]) {
			sendMsg(int64(userIds[i]), msg)
		}
	}
}

func JoinGroup(userId uint, comId string) (int, string) {
	contact := Contact{}
	contact.OwnerId = userId
	//contact.TargetId = comId
	contact.Type = 2
	community := Community{}

	utils.DB.Where("id=? or name=?", comId, comId).Find(&community)
	if community.Name == "" {
		return -1, "没有找到群"
	}
	utils.DB.Where("owner_id=? and target_id=? and type =2 ", userId, comId).Find(&contact)
	if !contact.CreatedAt.IsZero() {
		return -1, "已加过此群"
	} else {
		contact.TargetId = community.ID
		utils.DB.Create(&contact)
		return 0, "加群成功"
	}
}

// 从node的消息队列中接受消息，把接收到的消息存入redis，数据库不在存储消息数据
func sendMsg(userId int64, msg []byte) {
	// 从维护的map中获取目标目标用户的node
	node, ok := clientMap[userId]
	jsonMsg := Message{}
	json.Unmarshal(msg, &jsonMsg)
	ctx := context.Background()
	targetIdStr := strconv.Itoa(int(userId))
	userIdStr := strconv.Itoa(int(jsonMsg.UserId))
	jsonMsg.CreateTime = uint64(time.Now().Unix())
	// 获取指定用户是否上线
	r, err := utils.Red.Get(ctx, "online_"+userIdStr).Result()
	if err != nil {
		hlog.Info(err)
	}
	if r != "" {
		if ok {
			hlog.Info("sendMsg >>> userID: ", userId, "  msg:", string(msg))
			// 发送给目标用户的chan
			node.DataQueue <- msg
		}
	}
	// 发送数据给redis
	var key string
	if userId > jsonMsg.UserId {
		key = "msg_" + userIdStr + "_" + targetIdStr
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
	}
	// 获取全部信息
	res, err := utils.Red.ZRevRange(ctx, key, 0, -1).Result()
	if err != nil {
		hlog.Info(err)
	}
	// 信息排名加一，目的是让信息有序
	score := float64(cap(res)) + 1
	// 发送信息到redis
	ress, e := utils.Red.ZAdd(ctx, key, redis.Z{score, msg}).Result() //jsonMsg
	if e != nil {
		hlog.Info(e)
	}
	hlog.Info(ress)
}

// 获取缓存里面的消息
func RedisMsg(userIdA int64, userIdB int64, start int64, end int64, isRev bool) []string {
	ctx := context.Background()
	userIdStr := strconv.Itoa(int(userIdA))
	targetIdStr := strconv.Itoa(int(userIdB))
	var key string
	if userIdA > userIdB {
		key = "msg_" + targetIdStr + "_" + userIdStr
	} else {
		key = "msg_" + userIdStr + "_" + targetIdStr
	}

	var rels []string
	var err error
	if isRev {
		rels, err = utils.Red.ZRange(ctx, key, start, end).Result()
	} else {
		rels, err = utils.Red.ZRevRange(ctx, key, start, end).Result()
	}
	if err != nil {
		hlog.Error(err) //没有找到
	}
	return rels
}

// 更新用户心跳
func (node *Node) Heartbeat(currentTime uint64) {
	node.HeartbeatTime = currentTime
}

// 清理超时连接
func CleanConnection(param interface{}) (result bool) {
	result = true
	defer func() {
		if r := recover(); r != nil {
			hlog.Error("cleanConnection err", r)
		}
	}()
	currentTime := uint64(time.Now().Unix())
	for i := range clientMap {
		node := clientMap[i]
		if node.IsHeartbeatTimeOut(currentTime) {
			hlog.Info("心跳超时..... 关闭连接：", node)
			node.Conn.Close()
			rwLocker.Lock()
			delete(clientMap, i)
			rwLocker.Unlock()
		}

	}
	return result
}

// 用户心跳是否超时
func (node *Node) IsHeartbeatTimeOut(currentTime uint64) (timeout bool) {
	if node.HeartbeatTime+config.Timeout.HeartbeatHz <= currentTime {
		hlog.Info("心跳超时。。。自动下线", node)
		timeout = true
	}
	return
}
