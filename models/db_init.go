package models

import (
	"github.com/dutchcoders/goftp"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"tiktok/define"
	"time"
)

var DB *gorm.DB

var (
	RedisUserInfo *redis.Client //用户信息redis  存储库0
	RedisFeedInfo *redis.Client //视频信息redis  存储库1
)

var MyFTP *goftp.FTP

// DBInit 数据库连接
func DBInit() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := "root:123456@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Println("gorm Init Error: ", err)
		return
	}
}

// RedisInit  Redis
func RedisInit() {
	RedisUserInfo = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	RedisFeedInfo = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
}

// FTPInit FTP服务
func FTPInit() {
	var err error
	MyFTP, err = goftp.Connect(define.FTPAddress)
	if err != nil {
		log.Println("goftp Connection Error: ", err)
		return
	}

	err = MyFTP.Login(define.FTPUserName, define.FTPPassword) //登录操作
	if err != nil {
		log.Println("ftp Login Error: ", err)
		return
	}
	go func() {
		time.Sleep(time.Duration(120) * time.Second)
		MyFTP.Noop()
	}()
}
