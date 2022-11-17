package main

import (
	"github.com/Songguangyun/go-web/internal"
	"github.com/Songguangyun/go-web/internal/db"
	"github.com/Songguangyun/go-web/internal/global"
	"github.com/Songguangyun/go-web/internal/logger"
	"github.com/Songguangyun/go-web/utils"
)

func main() {
	internal.Viper()
	global.Logger = logger.InitLogger("")
	utils.InitValidator()
	global.DB = db.InitMysql()
	internal.RunServer()

}
