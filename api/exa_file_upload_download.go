package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/common"
	"github.com/my-gin-web/model/example"
	exampleRes "github.com/my-gin-web/model/example/response"
	"github.com/my-gin-web/utils/answer"
	"go.uber.org/zap"
)

type FileUploadAndDownloadApi struct{}

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {object} response.Response{data=exampleRes.ExaFileResponse,msg=string} "上传文件示例,返回包括文件详情"
// @Router /fileUploadAndDownload/upload [post]
func (u *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.ZapLog.Error("接收文件失败!", zap.Error(err))
		answer.FailWithMessage("接收文件失败", c)
		return
	}
	err, file = FileUploadAndDownloadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.ZapLog.Error("修改数据库链接失败!", zap.Error(err))
		answer.FailWithMessage("修改数据库链接失败", c)
		return
	}
	answer.OkWithDetailed(exampleRes.ExaFileResponse{File: file}, "上传成功", c)
}

// EditFileName 编辑文件名或者备注
func (u *FileUploadAndDownloadApi) EditFileName(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	if err := FileUploadAndDownloadService.EditFileName(file); err != nil {
		global.ZapLog.Error("编辑失败!", zap.Error(err))
		answer.FailWithMessage("编辑失败", c)
		return
	}
	answer.OkWithMessage("编辑成功", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body example.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {object} response.Response{msg=string} "删除文件"
// @Router /fileUploadAndDownload/deleteFile [post]
func (u *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	if err := FileUploadAndDownloadService.DeleteFile(file); err != nil {
		global.ZapLog.Error("删除失败!", zap.Error(err))
		answer.FailWithMessage("删除失败", c)
		return
	}
	answer.OkWithMessage("删除成功", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=common.PageResult,msg=string} "分页文件列表,返回包括列表,总数,页码,每页数量"
// @Router /fileUploadAndDownload/getFileList [post]
func (u *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo common.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := FileUploadAndDownloadService.GetFileRecordInfoList(pageInfo)
	if err != nil {
		global.ZapLog.Error("获取失败!", zap.Error(err))
		answer.FailWithMessage("获取失败", c)
	} else {
		answer.OkWithDetailed(common.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}
