package router

import (
	"github.com/gin-gonic/gin"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload")
	{
		fileUploadAndDownloadRouter.POST("upload", FileUploadAndDownloadApi.UploadFile)                                 // 上传文件
		fileUploadAndDownloadRouter.POST("getFileList", FileUploadAndDownloadApi.GetFileList)                           // 获取上传文件列表
		fileUploadAndDownloadRouter.POST("deleteFile", FileUploadAndDownloadApi.DeleteFile)                             // 删除指定文件
		fileUploadAndDownloadRouter.POST("editFileName", FileUploadAndDownloadApi.EditFileName)                         // 编辑文件名或者备注
		fileUploadAndDownloadRouter.POST("breakpointContinue", FileUploadAndDownloadApi.BreakpointContinue)             // 断点续传
		fileUploadAndDownloadRouter.GET("findFile", FileUploadAndDownloadApi.FindFile)                                  // 查询当前文件成功的切片
		fileUploadAndDownloadRouter.POST("breakpointContinueFinish", FileUploadAndDownloadApi.BreakpointContinueFinish) // 切片传输完成
		fileUploadAndDownloadRouter.POST("removeChunk", FileUploadAndDownloadApi.RemoveChunk)                           // 删除切片
	}
}
