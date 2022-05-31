package api

import (
	"github.com/my-gin-web/service"
)

var (
	BaseService                  = new(service.BaseService)
	UserService                  = new(service.UserService)
	GroupService                 = new(service.GroupService)
	OperationRecordService       = new(service.OperationRecordService)
	FileUploadAndDownloadService = new(service.FileUploadAndDownloadService)
)
