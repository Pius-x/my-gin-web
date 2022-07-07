package router

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", BaseApi.Login)
		baseRouter.GET("fsLogin", BaseApi.FsLogin) // 飞书登录(扫码\web)
		baseRouter.GET("fsBind", BaseApi.FsBind)   // 飞书关联绑定
	}
	return baseRouter
}
