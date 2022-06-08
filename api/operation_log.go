package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/model/common"
	"github.com/my-gin-web/model/operationlog"
	"github.com/my-gin-web/utils"
	"github.com/my-gin-web/utils/answer"
)

type OperationLogApi struct{}

// GetOperationLogList 分页获取操作记录列表
func (Api *OperationLogApi) GetOperationLogList(c *gin.Context) {
	
	var pageInfo operationlog.SearchOperationLog
	utils.Verify(&pageInfo, utils.Rules{}, c)

	list, total := OperationLogService.GetOperationLogList(pageInfo)

	answer.OkWithDetailed(common.PageResult{
		List:  list,
		Total: total,
	}, "获取成功", c)
}
