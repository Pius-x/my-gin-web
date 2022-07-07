package router

import (
	"github.com/gin-gonic/gin"
)

type OperationLogRouter struct{}

func (s *OperationLogRouter) InitOperationLogRouter(Router *gin.RouterGroup) {
	operationLogRouter := Router.Group("operationLog")
	{
		operationLogRouter.GET("getOperationLogList", OperationRecordApi.GetOperationLogList) // 获取SysOperationRecord列表
	}
}
