package service

import (
	"github.com/my-gin-web/model/operationlog"
)

type OperationRecordService struct{}

// CreateSysOperationRecord 创建一条操作记录
func (This *OperationRecordService) CreateSysOperationRecord(sysOperationRecord operationlog.SysOperationRecord) {
	operationLogModel.InsertNewRecord(sysOperationRecord)
}

// GetSysOperationRecordInfoList 分页获取操作记录列表
func (This *OperationRecordService) GetSysOperationRecordInfoList(pageInfo operationlog.SysOperationRecordSearch) (list []operationlog.SysOperationRecord, total int64) {

	return operationLogModel.GetPageRecordByKey(pageInfo)
}
