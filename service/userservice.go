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

	user := models.UserBasic{}
	err := c.BindAndValidate(&user)
	if err != nil {
		c.JSON(200,
			common.Result{
				-1,
				common.UserNamePassWordISEmpty,
				user,
			})
		return
	}
	hlog.Info(user.Name, "  >>>>>>>>>>>  ", user.PassWord, user.Identity)
	// 生成盐
	salt := fmt.Sprintf("%06d", rand.Int31())

	data := models.FindUserByName(user.Name)
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

	// 判断是否查询到用户
	if data.Name != "" {
		c.JSON(200,
			common.Result{
				-1,
				common.UserNameExist,
				user,
			})
		return
	}
	if user.PassWord != user.Identity {
		c.JSON(200,
			common.Result{
				-1,
				common.UserPasswordInconsistent,
				user,
			})
		return
	}
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
//func UpdateUser(ctx context.Context, c *app.RequestContext) {
//	user := models.UserBasic{}
//	id, _ := strconv.Atoi(c.PostForm("id"))
//	user.ID = uint(id)
//	user.Name = c.PostForm("name")
//	user.PassWord = c.PostForm("password")
//	user.Phone = c.PostForm("phone")
//	user.Avatar = c.PostForm("icon")
//	user.Email = c.PostForm("email")
//	fmt.Println("update :", user)
//
//	_, err := govalidator.ValidateStruct(user)
//	if err != nil {
//		fmt.Println(err)
//		c.JSON(200, gin.H{
//			"code":    -1, //  0成功   -1失败
//			"message": "修改参数不匹配！",
//			"data":    user,
//		})
//	} else {
//		models.UpdateUser(user)
//		c.JSON(200, gin.H{
//			"code":    0, //  0成功   -1失败
//			"message": "修改用户成功！",
//			"data":    user,
//		})
//	}
//
//}
