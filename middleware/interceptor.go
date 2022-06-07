package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/model/common/response"
	"github.com/my-gin-web/utils"
)

// InterceptHandler 拦截器
func InterceptHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := utils.GetClaims(c); err != nil {
			response.TokenFailMessage(err.Error(), c)
			c.Abort()
			return
		}

		// Panic 处理
		defer func() {
			if r := recover(); r != nil {
				var msg = fmt.Sprintf("Panic: %+v", r)
				utils.ZapErrorLog(msg)

				switch r.(type) {
				case error:
					response.FailWithMessage(r.(error).Error(), c)
				default:
					response.FailWithMessage(msg, c)
				}
			}
		}()

		c.Next()
	}
}
