package internal

import (
	"fmt"
	"github.com/Songguangyun/go-web/internal/constant"
	"github.com/Songguangyun/go-web/internal/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func Viper() *viper.Viper {
	var c string
	switch gin.Mode() {
	case gin.DebugMode:
		c = constant.ConfigDefaultFile
		fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, constant.ConfigDefaultFile)
	case gin.ReleaseMode:
		c = constant.ConfigReleaseFile
		fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, constant.ConfigReleaseFile)
	case gin.TestMode:
		c = constant.ConfigTestFile
		fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, constant.ConfigTestFile)
	}
	v := viper.New()
	v.SetConfigFile(c)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.ConfigServer); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.ConfigServer); err != nil {
		fmt.Println(err)
	}
	return v
}
