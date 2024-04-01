package service

import (
	"chat/common"
	"chat/models"
	"chat/models/vo"
	"chat/utils"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
	"github.com/jinzhu/copier"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(ctx context.Context, c *app.RequestContext) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(http.StatusOK, common.Result{
		Code:    0,
		Message: common.UserNameExist,
		Data:    data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	data := vo.UserRegisterVo{}
	err := c.BindAndValidate(&data)
	// 修改，现在通过结构体的vd字段进行校验
	// 判断前端传来的数据是否正常
	//if user.Name == "" || user.PassWord == "" || user.Identity == "" {
	//	c.JSON(200,
	//		common.Result{
	//			-1,
	//			common.UserNamePassWordISEmpty,
	//			user,
	//		})
	//	return
	//}

	if err != nil {
		c.JSON(200,
			common.Result{
				-1,
				common.UserNamePassWordISEmpty,
				data,
			})
		return
	}
	hlog.Info(data.Name, "  >>>>>>>>>>>  ", data.PassWord, data.Identity)
	// 生成盐
	salt := fmt.Sprintf("%06d", rand.Int31())
	// 判断两次密码是否相等
	if data.PassWord != data.Identity {
		c.JSON(200,
			common.Result{
				-1,
				common.UserPasswordInconsistent,
				data,
			})
		return
	}
	// 按照用户名查找用户
	user := models.FindUserByName(data.Name)
	// 判断是否查询到用户
	if user.Name != "" {
		c.JSON(200,
			common.Result{
				-1,
				common.UserNameExist,
				data,
			})
		return
	}
	user = models.UserBasic{}
	copier.Copy(&user, &data)
	user.PassWord = utils.MakePassword(user.PassWord, salt)
	user.Salt = salt
	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	models.CreateUser(user)
	c.JSON(http.StatusOK, common.Result{
		0,
		common.UserNamePassWordISEmpty,
		user,
	})
}

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(ctx context.Context, c *app.RequestContext) {
	data := vo.UserLoginVo{}
	err := c.BindAndValidate(&data)
	if err != nil {
		c.JSON(200,
			common.Result{
				-1,
				common.UserNamePassWordISEmpty,
				data,
			})
		return
	}
	hlog.Info(data.Name, data.PassWord)
	user := models.FindUserByName(data.Name)
	if user.Name == "" {
		c.JSON(http.StatusOK, common.Result{
			-1,
			common.UserISEmpty,
			data,
		})
		return
	}

	flag := utils.ValidPassword(data.PassWord, user.Salt, user.PassWord)
	if !flag {
		c.JSON(http.StatusOK, common.Result{
			-1,
			common.UserPasswordError,
			data,
		})
		return
	}
	c.JSON(http.StatusOK, common.Result{
		0,
		common.UserLoginSucceed,
		user,
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	hlog.Info("删除用户", id)
	user.ID = uint(id)
	//models.DeleteUser(user)
	c.JSON(http.StatusOK, common.Result{
		Code:    0,
		Message: common.UserDeletedSucceed,
		Data:    user,
	})

}

// 查找所有好友
func SearchFriends(ctx context.Context, c *app.RequestContext) {
	id, _ := strconv.Atoi(c.PostForm("userId"))
	users := models.SearchFriend(uint(id))
	c.JSON(http.StatusOK, common.H{
		Code:  0,
		Rows:  users,
		Total: len(users),
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	user := models.UserBasic{}
	err := c.BindAndValidate(&user)
	if err != nil {
		c.JSON(http.StatusOK, common.Result{
			Code:    -1,
			Message: common.UserParamError,
			Data:    user,
		})
		return
	}
	hlog.Info("update :", user)
	models.UpdateUser(user)
	c.JSON(http.StatusOK, common.Result{
		Code:    0,
		Message: common.UserUpdateSucceed,
		Data:    user,
	})
}

// 添加好友
func AddFriend(ctx context.Context, c *app.RequestContext) {
	user := vo.FriendVo{}
	// 检验数据是否合法
	err := c.BindAndValidate(&user)
	if err != nil {
		c.JSON(http.StatusOK, common.H{
			Code: -1,
			Data: nil,
			Msg:  common.UserISEmpty,
		})
		return
	}
	code, msg := models.AddFriend(user.UserId, user.TargetName)
	// 查找不到直接返回
	if code == 0 {
		hlog.Info(msg)
		c.JSON(http.StatusOK, common.H{
			Code: 0,
			Data: code,
			Msg:  msg,
		})
	} else {
		c.JSON(http.StatusOK, common.H{
			Code: -1,
			Data: nil,
			Msg:  msg,
		})
	}
}
func FindByID(ctx context.Context, c *app.RequestContext) {
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	data := models.FindByID(uint(userId))
	c.JSON(http.StatusOK, common.H{
		Code: 0,
		Data: data,
		Msg:  "ok",
	})
}

// 防止跨域站点伪造请求
var upGrader = websocket.HertzUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
}

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

// 发送消息
func RedisMsg(ctx context.Context, c *app.RequestContext) {
	userIdA, _ := strconv.Atoi(c.PostForm("userIdA"))
	userIdB, _ := strconv.Atoi(c.PostForm("userIdB"))
	start, _ := strconv.Atoi(c.PostForm("start"))
	end, _ := strconv.Atoi(c.PostForm("end"))
	isRev, _ := strconv.ParseBool(c.PostForm("isRev"))
	res := models.RedisMsg(int64(userIdA), int64(userIdB), int64(start), int64(end), isRev)
	c.JSON(http.StatusOK, common.H{
		Code:  0,
		Rows:  "0k",
		Total: res,
	})
}

// 消息处理器
func MsgHandler(ctx context.Context, ws *websocket.Conn) {
	// 由前端控制websocket的关闭
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

// 发送指定消息
func SendUserMsg(ctx context.Context, c *app.RequestContext) {
	//models.Chat(c.Writer, c.Request)
}

// 新建群
func CreateCommunity(ctx context.Context, c *app.RequestContext) {
	community := models.Community{}
	err := c.BindAndValidate(&community)
	if err != nil {
		c.JSON(http.StatusOK, common.H{
			Code: -1,
			Msg:  common.UserParamError,
		})
	}
	data, msg := models.CreateCommunity(community)
	// 查找不到直接返回
	if data == 0 {
		hlog.Info(msg)
		c.JSON(http.StatusOK, common.H{
			Code: 0,
			Data: data,
			Msg:  msg,
		})
	} else {
		c.JSON(http.StatusOK, common.H{
			Code: -1,
			Data: nil,
			Msg:  msg,
		})
	}
}

// 加载群列表
func LoadCommunity(ctx context.Context, c *app.RequestContext) {
	ownerId, _ := strconv.Atoi(c.PostForm("ownerId"))
	data, msg := models.LoadCommunity(uint(ownerId))
	// 查找不到直接返回
	if len(data) != 0 {
		hlog.Info(msg)
		c.JSON(http.StatusOK, common.H{
			Code:  0,
			Rows:  data,
			Total: msg,
		})
	} else {
		c.JSON(http.StatusOK, common.H{
			Code: -1,
			Data: nil,
			Msg:  msg,
		})
	}
}

// 加入群 userId uint, comId uint
func JoinGroups(ctx context.Context, c *app.RequestContext) {
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	comId := c.PostForm("comId")
	data, msg := models.JoinGroup(uint(userId), comId)
	if data == 0 {
		hlog.Info(msg)
		c.JSON(http.StatusOK, common.H{
			Code: 0,
			Data: data,
			Msg:  msg,
		})
	} else {
		c.JSON(http.StatusOK, common.H{
			Code: -1,
			Data: nil,
			Msg:  msg,
		})
	}
}
