package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 0
	FAILED  = -1
)

func Result(code int, message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		message,
	})
}

func Success(c *gin.Context) {
	Result(SUCCESS, "操作成功", map[string]interface{}{}, c)
}

func SuccessWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, msg, map[string]interface{}{}, c)
}

func ErrorWithMessage(msg string, c *gin.Context) {
	Result(FAILED, msg, map[string]interface{}{}, c)
}

func Data(data interface{}, c *gin.Context) {
	Result(SUCCESS, "操作成功", data, c)
}
