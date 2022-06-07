package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/model/common/response"
	"github.com/my-gin-web/model/operationlog"
	"github.com/my-gin-web/utils"
)

type OperationRecordApi struct{}

// GetSysOperationRecordList 分页获取操作记录列表
func (s *OperationRecordApi) GetSysOperationRecordList(c *gin.Context) {
	var pageInfo operationlog.SysOperationRecordSearch
	utils.Verify(&pageInfo, utils.Rules{}, c)

	list, total := OperationRecordService.GetSysOperationRecordInfoList(pageInfo)

	response.OkWithDetailed(response.PageResult{
		List:  list,
		Total: total,
	}, "获取成功", c)
}
