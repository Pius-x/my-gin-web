package example

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ExcelApi
	FileUploadAndDownloadApi
}

var (
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
