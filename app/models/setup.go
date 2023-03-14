package models

import (
	"AbitService/app/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbAbit *gorm.DB

func ConnectDatabase() {
	conf := config.New()
	//dsn := "sqlserver://" + conf.Abit.UserName + ":" + conf.Abit.Password + "@" + conf.Abit.Host + "?database=" + conf.Abit.Name
	dsn := conf.Abit.UserName + ":" + conf.Abit.Password + "@tcp(" + conf.Abit.Host + ")/" + conf.Abit.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	DbAbit = db
}
