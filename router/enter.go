package router

import "github.com/my-gin-web/api"

var (
	BaseApi            = new(api.BaseApi)
	UserApi            = new(api.UserApi)
	GroupApi           = new(api.GroupApi)
	OperationRecordApi = new(api.OperationRecordApi)
)

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
