package api

import (
	"MyTodoList/service"

	"github.com/gin-gonic/gin"
)


// UserRegister 用户注册路由函数
func UserRegister(c *gin.Context) {
	// 1. 声明服务对象
	var userRegister service.UserService
	// 2. 绑定服务对象
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(400, err)
	} else {
		res := userRegister.Register()
		c.JSON(200, res)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(400, err)
	} else {
		res := userLogin.Login()
		c.JSON(200, res)
	}
}