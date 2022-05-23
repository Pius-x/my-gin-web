package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord").Use(middleware.OperationRecord())
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.GET("getSysOperationRecordList", authorityMenuApi.GetSysOperationRecordList)            // 获取SysOperationRecord列表
		operationRecordRouter.DELETE("deleteSysOperationRecordByIds", authorityMenuApi.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
	}
}
