package router

import "github.com/my-gin-web/api"

// 基本系统API
var (
	BaseApi            = new(api.BaseApi)
	UserApi            = new(api.UserApi)
	GroupApi           = new(api.GroupApi)
	OperationRecordApi = new(api.OperationLogApi)
)

// 业务系统API
var (
	ExcelApi                 = new(api.ExcelApi)
	FileUploadAndDownloadApi = new(api.FileUploadAndDownloadApi)
)

type Group struct {
	BaseRouter
	UserRouter
	GroupRouter
	OperationLogRouter

	ExcelRouter
	FileUploadAndDownloadRouter
}

var GroupApp = new(Group)
