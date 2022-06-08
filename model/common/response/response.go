package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int64  `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Success = int64(iota)
	Error
	TokenError
)

func Result(code int64, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})

	panic(nil)
}

func Ok(c *gin.Context) {
	Result(Success, map[string]any{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(Success, map[string]any{}, message, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "操作成功", c)
}

func OkWithDetailed(data any, message string, c *gin.Context) {
	Result(Success, data, message, c)
}

func Fail(c *gin.Context) {
	Result(Error, map[string]any{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(Error, map[string]any{}, message, c)
}

func FailWithDetailed(data any, message string, c *gin.Context) {
	Result(Error, data, message, c)
}

func TokenFailMessage(message string, c *gin.Context) {
	Result(TokenError, map[string]any{}, message, c)
}
