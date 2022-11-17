package internal

import (
	"fmt"
	"github.com/Songguangyun/go-web/internal/global"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func RunServer() {
	address := fmt.Sprintf(":%d", global.ConfigServer.App.AppPort)
	// 初始化路由
	router := InitRouter()
	s := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)

	global.Logger.WithFields(logrus.Fields{
		"PID": os.Getpid(),
	}).Infoln("web服务运行端口", address)

	global.Logger.WithFields(logrus.Fields{
		"PID": os.Getpid(),
	}).Infoln(s.ListenAndServe().Error())
}
