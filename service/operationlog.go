package service

import (
	"github.com/my-gin-web/model/operationlog"
)

type OperationRecordService struct{}

// CreateSysOperationRecord 创建一条操作记录
func (This *OperationRecordService) CreateSysOperationRecord(sysOperationRecord operationlog.SysOperationRecord) (err error) {
	err = operationLogModel.InsertNewRecord(sysOperationRecord)
	return err
}

// GetSysOperationRecordInfoList 分页获取操作记录列表
func (This *OperationRecordService) GetSysOperationRecordInfoList(pageInfo operationlog.SysOperationRecordSearch) (list []operationlog.SysOperationRecord, total int64, err error) {

	return operationLogModel.GetPageRecordByKey(pageInfo)
}
