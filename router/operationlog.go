package router

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/middleware"
)

type OperationLogRouter struct{}

func (s *OperationLogRouter) InitOperationLogRouter(Router *gin.RouterGroup) {
	operationLogRouter := Router.Group("operationLog").Use(middleware.OperationRecord())
	{
		operationLogRouter.GET("getOperationLogList", OperationRecordApi.GetOperationLogList) // 获取SysOperationRecord列表
	}
}
