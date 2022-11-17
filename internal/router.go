package internal

import (
	api "github.com/Songguangyun/go-web/api"
	"github.com/Songguangyun/go-web/internal/global"
	"github.com/Songguangyun/go-web/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func InitRouter() *gin.Engine {
	app := gin.Default()
	// 跨域
	app.Use(middleware.CorsHandler())

	PublicGroup := app.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "200")
		})
	}
	{
		PublicGroup.POST("/login", api.LoginApi)
	}

	PrivateGroup := app.Group("/user")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		PrivateGroup.GET("/remove", api.RemoveUserApi)
	}
	global.Logger.WithFields(logrus.Fields{
		"PID": os.Getpid(),
	}).Infoln("路由注册完成")
	return app
}
