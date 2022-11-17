package api

import (
	"fmt"
	"github.com/Songguangyun/go-web/internal/global"
	"github.com/Songguangyun/go-web/request"
	"github.com/Songguangyun/go-web/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func LoginApi(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	// 检验字段
	if err := global.Validator.Struct(req); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			response.ErrorWithMessage(err.Translate(global.ValidatorTrans), c)
			return
		}
	}
	fmt.Printf("oginRequest: %+v\n", req)
	response.SuccessWithMessage("login success", c)
}
