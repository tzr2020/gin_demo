package dao

import (
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitDB() (err error) {
	dsn := "root:123456@(127.0.0.1:3306)/db_todo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Panicf("open database failure, err: %v", err)
	}
	return DB.DB().Ping()
}

func CloseDB() {
	DB.Close()
}
