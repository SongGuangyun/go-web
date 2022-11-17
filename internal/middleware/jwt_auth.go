package middleware

import (
	"github.com/Songguangyun/go-web/internal/logger"
	"github.com/Songguangyun/go-web/utils/response"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.Abort()
			response.Result(10001, "未携带token", map[string]interface{}{}, c)
			return
		}
		logger.NewInfo(authorization)
		//for k, v := range c.Request.Header {
		//	fmt.Println(k, v)
		//}
	}
}
