package middleware

import (
	"MyTodoList/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := 200 
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401
			}
		}
		if code != 200 {
			ctx.JSON(400, gin.H{
				"status": code,
				"msg": "Token解析失败",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}