package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/model"
	"mall/result"
	"mall/service"
	"strconv"
)

var UserService = service.NewUService()

func UserRegister(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	err := UserService.ManagerRegister(&user)
	if err != nil {
		result.Fail(c, result.Response{
			Code:    result.ErrorFailEncryption,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "注册成功")
	}
	//c.Request.Header.Del("Authorization")
}
func UserLogin(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	username := user.Username
	password := user.Password
	token, err := UserService.UserLogin(username, password)

	if err != nil || token == "" {
		result.Fail(c, result.Response{
			Code:    result.ErrorAuthToken,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		//将token设置到请求头中
		c.Header("token", token)
		result.OkWithData(c, gin.H{
			"token": token,
		})
	}

}

func GetUserInfo(c *gin.Context) {
	token := c.MustGet("token")
	user, err := UserService.GetUserInfo(token.(string))
	if err != nil {
		result.Fail(c, result.Response{
			Code:    result.ErrorAuthCheckTokenFail,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, user)
	}

}

func UserUpdate(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	id, _ := strconv.Atoi(c.Param("id"))
	user.ID = uint(id)
	err := UserService.UpdateUser(&user)
	if err != nil {
		result.Fail(c, result.Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithMsg(c, "修改成功")
	}
}
func DisableUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	UserService.DisableUser(id)
	result.OkWithMsg(c, "修改成功")
}
func UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("img")
	if err != nil {
		return
	}
	name := file.Filename
	fmt.Println(name)
}

func Logout(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "" {
		c.Request.Header.Del("token")
		result.OkWithMsg(c, "退出成功")
	}
}

func UserList(c *gin.Context) {
	userList, err := UserService.UserList()
	if err != nil {
		result.Fail(c, result.Response{
			Code:    509,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		result.OkWithData(c, userList)
	}
}

func SendEmail(c *gin.Context) {

}

func ValidEmail(c *gin.Context) {

}
