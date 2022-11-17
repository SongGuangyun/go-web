package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT")
		ctx.Header("Access-Control-Allow-Headers", "*")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
		ctx.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
		ctx.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
		ctx.Header("Content-Type", "application/json")                                                                                                                                                           // 设置返回格式是json
		if ctx.Request.Method == http.MethodOptions {
			ctx.JSON(http.StatusOK, "Options Request!")
		}
		ctx.Next()
	}
}
