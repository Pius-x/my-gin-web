package service

import (
	"github.com/my-gin-web/model/operationlog"
)

type OperationLogService struct{}

// CreateOperationLog 创建一条操作记录
func (This *OperationLogService) CreateOperationLog(operationLog operationlog.TOperationLog) {
	
	operationLogModel.InsertNewRecord(operationLog)
}

// GetOperationLogList 分页获取操作记录列表
func (This *OperationLogService) GetOperationLogList(pageInfo operationlog.SearchOperationLog) (list []operationlog.TOperationLog, total int64) {

	return operationLogModel.GetPageRecordByKey(pageInfo)
}
