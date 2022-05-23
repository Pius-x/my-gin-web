package middleware

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

// InterceptHandler 拦截器
func InterceptHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)
		c.Next()
		fmt.Printf("%+v \n", waitUse)
		//// 获取请求的PATH
		//obj := c.Request.URL.Path
		//// 获取请求方法
		//act := c.Request.Method
		//// 获取用户的角色
		//sub := waitUse.AuthorityId
		//e := casbinService.Casbin()
		//// 判断策略中是否存在
		//success, _ := e.Enforce(sub, obj, act)
		//if global.GVA_CONFIG.System.Env == "develop" || success {
		//	c.Next()
		//} else {
		//	response.FailWithDetailed(gin.H{}, "权限不足", c)
		//	c.Abort()
		//	return
		//}
	}
}
