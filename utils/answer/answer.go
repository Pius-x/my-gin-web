package answer

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

func OkWithMessage(message string, c *gin.Context) {
	Result(Success, map[string]any{}, message, c)
}

func OkWithDetailed(data any, message string, c *gin.Context) {
	Result(Success, data, message, c)
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
