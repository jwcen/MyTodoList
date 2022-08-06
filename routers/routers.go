package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"MyTodoList/api"
	"MyTodoList/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))  // 存储session

	// 编写基础路由
	v1 := r.Group("api/v1")
	{
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 身份验证，jwt鉴权
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task", api.CreateTask)
		}
	}

	return r
}