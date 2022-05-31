package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/common/response"
	"github.com/my-gin-web/model/operationlog"
	"go.uber.org/zap"
)

type OperationRecordApi struct{}

// GetSysOperationRecordList 分页获取操作记录列表
func (s *OperationRecordApi) GetSysOperationRecordList(c *gin.Context) {
	var pageInfo operationlog.SysOperationRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := OperationRecordService.GetSysOperationRecordInfoList(pageInfo); err != nil {
		global.ZapLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}
