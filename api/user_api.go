package api

import (
	"github.com/Songguangyun/go-web/utils/response"
	"github.com/gin-gonic/gin"
)

func RemoveUserApi(c *gin.Context) {
	d := [...]int{1, 2, 3, 4, 5, 6}
	response.Data(d, c)
}
