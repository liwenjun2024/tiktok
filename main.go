package main

import (
	"tiktok/models"
	"tiktok/router"
)

func main() {
	r := router.InitRouter()
	initDB()
	r.Run(":8080")
}

// 初始化服务
func initDB() {
	models.DBInit()
	models.RedisInit()
	models.FTPInit()
}
