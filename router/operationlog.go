package router

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/middleware"
)

type OperationLogRouter struct{}

func (s *OperationLogRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord").Use(middleware.OperationRecord())
	{
		operationRecordRouter.GET("getSysOperationRecordList", OperationRecordApi.GetSysOperationRecordList) // 获取SysOperationRecord列表
	}
}
