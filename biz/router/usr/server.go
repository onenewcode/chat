package usr

import (
	"chat/biz/domain"
	"chat/biz/handler/usr"
	"chat/biz/repository"
	"chat/common"
	"chat/common/vo"
	"chat/models"

	"chat/utils"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(ctx context.Context, c *app.RequestContext) {
	data := repository.UserBasicRepo.GetList(ctx, "", "", "", -1, 10)
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
// @Router /user/createUser [post]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	data := vo.UserRegisterVo{}
	err := c.BindAndValidate(&data)
	if err != nil {
		c.JSON(200,
			common.Result{
				Code:    -1,
				Message: common.UserNamePassWordISEmpty,
				Data:    data,
			})
		return
	}
	hlog.Info(data.Name, "  >>>>>>>>>>>  ", data.PassWord, data.Identity)
	// 判断两次密码是否相等
	if data.PassWord != data.Identity {
		c.JSON(200,
			common.Result{
				Code:    -1,
				Message: common.UserPasswordInconsistent,
				Data:    data,
			})
		return
	}
	// 判断是否存在重名用户
	user := repository.UserBasicRepo.FindByName(ctx, data.Name)
	if user.Name != "" {
		c.JSON(200,
			common.Result{
				Code:    -1,
				Message: common.UserNameExist,
				Data:    data,
			})
		return
	}

	salt := fmt.Sprintf("%06d", rand.Int31())
	user = domain.UserBasic{}
	copier.Copy(&user, &data)
	user.PassWord = utils.MakePassword(user.PassWord, salt)
	user.Salt = salt
	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	repository.UserBasicRepo.Create(ctx, user)
	c.JSON(http.StatusOK, common.Result{
		Code:    0,
		Message: "注册成功",
		Data:    user,
	})
}

// GetUserList
// @Summary 用户登录
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/login[post]
func Login(ctx context.Context, c *app.RequestContext) {
	data := vo.UserLoginVo{}
	err := c.BindAndValidate(&data)
	if err != nil {
		c.JSON(200,
			common.Result{
				Code:    -1,
				Message: common.UserNamePassWordISEmpty,
				Data:    data,
			})
		return
	}
	hlog.Info(data.Name, data.PassWord)
	user := repository.UserBasicRepo.FindByName(ctx, data.Name)
	if user.Name == "" {
		c.JSON(http.StatusOK, common.Result{
			Code:    -1,
			Message: common.UserISEmpty,
			Data:    data,
		})
		return
	}

	flag := utils.ValidPassword(data.PassWord, user.Salt, user.PassWord)
	if !flag {
		c.JSON(http.StatusOK, common.Result{
			Code:    -1,
			Message: common.UserPasswordError,
			Data:    data,
		})
		return
	}
	c.JSON(http.StatusOK, common.Result{
		Code:    0,
		Message: common.UserLoginSucceed,
		Data:    user,
	})
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param id query int true "用户 ID"
// @Success 200 {object} common.Result{data=models.UserBasic} "成功"
// @Failure 400 {object} common.Result "无效的请求参数"
// @Failure 500 {object} common.Result "服务器内部错误"
// @Router /user/deleteUser [delete]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	// 删除所有用户缓存
	// utils.DeleteALLFriends()
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, common.Result{
			Code:    -1,
			Message: common.UserParamError,
			Data:    nil,
		})
		return
	}
	hlog.Info("删除用户", id)
	repository.UserBasicRepo.Delete(ctx, uint(id))
	//models.DeleteUser(user)
	c.JSON(http.StatusOK, common.Result{
		Code:    0,
		Message: common.UserDeletedSucceed,
		Data:    nil,
	})

}

// SearchFriends 查找与指定用户相关的所有好友。
// 此操作需要一个有效的用户 ID 作为查询条件，该 ID 通过 POST 请求体中的"userId"字段传递。
// 注意：此 API 要求提供完整的用户名称（根据上下文理解，这可能是文档中的误导；实际实现中是基于用户 ID 查找）。
// @Summary 查找所有好友
// @Description 根据提供的用户 ID 查找该用户的所有好友。注意，这里并非使用用户名进行查找，而是用户 ID。
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param userId formData integer true "用户 ID"
// @Success 200 {object} common.Result{data=common.PaginatedResponse{rows=[]models.UserBasic,total=int}} "成功获取好友列表"
// @Failure 400 {object} common.Result "无效的请求参数"
// @Failure 500 {object} common.Result "服务器内部错误"
// @Router /user/searchFriends [get]
func SearchFriends(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, common.Result{
			Code:    -1,
			Message: common.UserParamError,
			Data:    nil,
		})
		return
	}
	users := repository.ContactRepo.SearchFriend(ctx, uint(id))
	c.JSON(http.StatusOK, common.H{
		Code:  0,
		Rows:  users,
		Total: len(*users),
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
// @Router /user/updateUser [put]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	// 删除用户缓存
	// utils.DeleteALLFriends()
	user_vo := vo.UserUpdate{}
	err := c.BindAndValidate(&user_vo)
	user := models.UserBasic{}
	copier.Copy(&user, &user_vo)
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

// AddFriend 添加好友
// @Summary 添加好友关系
// @Description 根据提供的用户 ID 和目标用户名创建一个新的好友关系。
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param friend body vo.FriendVo true "好友信息"
// @Success 201 {object} common.Result "好友添加成功"
// @Failure 400 {object} common.Result "无效的请求参数或验证失败"
// @Failure 500 {object} common.Result "服务器内部错误"
// @Router /user/addFriend [post]
func AddFriend(ctx context.Context, c *app.RequestContext) {
	user := vo.FriendVo{}
	// 检验数据是否合法
	err := c.BindAndValidate(&user)
	if err != nil {
		c.JSON(http.StatusOK, common.H{
			Code: -1,
			Data: nil,
			Msg:  common.UserParamError,
		})
		return
	}
	err = repository.ContactRepo.AddFriend(ctx, user.UserId, user.TargeId)
	// 查找不到直接返回
	if err != nil {
		c.JSON(http.StatusOK, common.H{
			Code: -1,
			Data: nil,
			Msg:  common.UserParamError,
		})
		return
	}
	c.JSON(http.StatusOK, common.H{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	})
}

// FindByID 根据用户 ID 查找用户信息
// @Summary 根据用户 ID 获取用户详情
// @Description 使用用户 ID 来检索用户的详细信息。
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param userId query int true "用户 ID"
// @Success 200 {object} common.Result{data=models.UserBasic} "成功获取用户信息"
// @Failure 400 {object} common.Result "无效的请求参数"
// @Failure 404 {object} common.Result "找不到用户"
// @Failure 500 {object} common.Result "服务器内部错误"
// @Router /user/{userId} [get]
func FindByID(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, common.Result{
			Code:    -1,
			Message: common.UserParamError,
			Data:    nil,
		})
		return
	}
	data := repository.UserBasicRepo.FindByID(ctx, uint(id))
	c.JSON(http.StatusOK, common.H{
		Code: 0,
		Data: data,
		Msg:  "ok",
	})
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

// LoadCommunity 加载用户所属的群列表。
// @Summary 加载用户所属的群列表
// @Description 根据提供的用户 ID 加载该用户所拥有的或加入的所有群组信息。
// @Tags 群组模块
// @Accept json
// @Produce json
// @Param ownerId query int true "所有者的用户 ID"
// @Success 200 {object} common.Result{data=common.PaginatedResponse{rows=[]models.Community,total=int}} "成功获取群列表"
// @Failure 400 {object} common.Result "无效的请求参数"
// @Failure 500 {object} common.Result "服务器内部错误"
// @Router /community/load/:ownerId [get]
func LoadCommunity(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.ParseUint(c.Param("ownerId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.H{
			Code: -1,
			Msg:  common.UserParamError,
		})
	}
	data, err := repository.CommunityRepo.Load(ctx, uint(id))
	if err != nil {
		c.JSON(http.StatusOK, common.H{
			Code: -1,
			Msg:  common.UserParamError,
		})
	}
	c.JSON(http.StatusOK, common.H{
		Code:  0,
		Rows:  data,
		Total: "ok",
	})
}

// SendMsg 向指定群组发送消息。
// @Summary 向指定群组发送消息
// @Description 根据上下文和请求内容向特定群组发送一条消息。
// @Tags 消息模块
// @Accept json
// @Produce json
// @Param message body vo.MessageVo true "消息内容"
// @Success 200 {object} common.Result "消息发送成功"
// @Failure 400 {object} common.Result "无效的请求参数"
// @Failure 500 {object} common.Result "服务器内部错误"
// @Router /message/send [post]
func SendMsg(ctx context.Context, c *app.RequestContext) {
	usr.SendMsg(ctx, c)
}

// TODO
// JoinGroups 用户加入指定群组。
// @Summary 用户加入指定群组
// @Description 根据提供的用户 ID 和群组 ID 将用户加入到指定的群组中。
// @Tags 群组模块
// @Accept json
// @Produce json
// @Param userId formData integer true "用户 ID"
// @Param comId formData string true "群组 ID"
// @Success 201 {object} common.Result "成功加入群组"
// @Failure 400 {object} common.Result "无效的请求参数"
// @Failure 500 {object} common.Result "服务器内部错误"
// @Router /group/join [post]
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

// TODO
// 发送消息
func RedisMsg(ctx context.Context, c *app.RequestContext) {
	userIdA, _ := strconv.Atoi(c.PostForm("userIdA"))
	userIdB, _ := strconv.Atoi(c.PostForm("userIdB"))
	start, _ := strconv.Atoi(c.PostForm("start"))
	end, _ := strconv.Atoi(c.PostForm("end"))
	isRev, _ := strconv.ParseBool(c.PostForm("isRev"))
	// 构建消息，发送给 redis，用 redis 作为消息沟通的中间件，
	res := models.RedisMsg(int64(userIdA), int64(userIdB), int64(start), int64(end), isRev)
	c.JSON(http.StatusOK, common.H{
		Code:  0,
		Rows:  "0k",
		Total: res,
	})
}

// TODO
// 发送指定消息
func SendUserMsg(ctx context.Context, c *app.RequestContext) {
	models.Chat(c)
}
