package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

// Init 数据库连接
func Init() *gorm.DB {
	dsn := "账户:密码@tcp(localhost:3306)/数据库名称?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error: ", err)
	}
	return db
}

//下面可以放其他的比如 Redis
