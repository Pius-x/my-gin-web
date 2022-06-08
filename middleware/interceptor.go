package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/utils"
	"github.com/my-gin-web/utils/answer"
)

// InterceptPublic 拦截器
func InterceptPublic() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Panic 处理
		defer deferHandler(c)()

		c.Next()
	}
}

// InterceptPrivate 拦截器
func InterceptPrivate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := utils.GetClaims(c); err != nil {
			answer.TokenFailMessage(err.Error(), c)
			c.Abort()
			return
		}

		// Panic 处理
		defer deferHandler(c)()

		c.Next()
	}
}

func deferHandler(c *gin.Context) func() {
	return func() {
		if r := recover(); r != nil {
			var msg = fmt.Sprintf("Panic: %+v", r)
			utils.ZapErrorLog(msg)

			switch r.(type) {
			case error:
				answer.FailWithMessage(r.(error).Error(), c)
			default:
				answer.FailWithMessage(msg, c)
			}
		}
	}
}
