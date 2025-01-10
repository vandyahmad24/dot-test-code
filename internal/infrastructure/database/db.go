package infrastructure

import (
	"dot-test-vandy/config"
	"fmt"
	"log"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql() *gorm.DB {
	var err error

	username := config.EnvConfig.MYSQL_USERNAME
	password := config.EnvConfig.MYSQL_PASSWORD
	host := config.EnvConfig.MYSQL_HOST
	port := config.EnvConfig.MYSQL_PORT
	dbname := config.EnvConfig.MYSQL_DATABASE

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db

}
