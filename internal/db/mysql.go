package db

import (
	"fmt"
	"github.com/Songguangyun/go-web/internal/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		global.ConfigServer.Mysql.DBUsername,
		global.ConfigServer.Mysql.DBPassword,
		global.ConfigServer.Mysql.DBHost,
		global.ConfigServer.Mysql.DBPort,
		"mysql")
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		fmt.Println("0", "Open failed ", err.Error(), dsn)
	}
	return db
}
