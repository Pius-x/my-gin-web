package api

import (
	"github.com/my-gin-web/service"
)

// 基本系统服务
var (
	BaseService         = new(service.BaseService)
	UserService         = new(service.UserService)
	GroupService        = new(service.GroupService)
	OperationLogService = new(service.OperationLogService)
)

// 业务系统服务
var (
	FileUploadAndDownloadService = new(service.FileUploadAndDownloadService)
)
