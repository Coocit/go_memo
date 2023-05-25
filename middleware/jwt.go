package middleware

import (
	"github.com/gin-gonic/gin"
	"go_memo/pkg/utils"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 402
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 // 无权限，token是无权限的，是假的
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 // Token 无效
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
