package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	Success = int64(iota)
	Error
	TokenError
)

func Result(code int64, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})

	panic(nil)
}

func Ok(c *gin.Context) {
	Result(Success, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(Success, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(Success, data, message, c)
}

func Fail(c *gin.Context) {
	Result(Error, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(Error, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(Error, data, message, c)
}

func TokenFailMessage(message string, c *gin.Context) {
	Result(TokenError, map[string]interface{}{}, message, c)
}
