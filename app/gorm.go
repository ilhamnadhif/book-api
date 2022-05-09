package app

import (
	"book-api/config"
	"book-api/helper"
	"book-api/model/schema"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm(config config.DatabaseConfig) *gorm.DB {
	//dsn := "root:root@tcp(127.0.0.1:3306)/book-api?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.HostPort, config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	db.AutoMigrate(&schema.User{}, &schema.Book{})

	return db
}
