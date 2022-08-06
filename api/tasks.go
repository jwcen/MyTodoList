package api

import (
	"MyTodoList/pkg/utils"
	"MyTodoList/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateTask 创建一条备忘录
func CreateTask(c *gin.Context) {
	// 1. 创建服务对象
	var cts service.CreateTaskService
	// 2. 先验证身份
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	// 3. 绑定服务对象
	if err := c.ShouldBind(&cts); err == nil {
		res := cts.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(400, err)	
	}
	
}