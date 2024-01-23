package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"junjun-box-api/config"
	"junjun-box-api/model"
	"log/slog"
)

var DB *gorm.DB

func Initialize() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.Conf.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
	}
	if err != nil {
		slog.Error(err.Error())
		return
	}
	err = DB.AutoMigrate(&model.Member{})
	if err != nil {
		return
	}
}
