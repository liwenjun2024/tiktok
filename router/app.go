package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tiktok/controller"
	_ "tiktok/docs"
	"tiktok/service"
)

func InitRouter() *gin.Engine {
	go service.RunMessageServer()
	r := gin.Default()
	r.Static("/static", "./public")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter := r.Group("/douyin")
	apiRouter.POST("/user/login/", controller.Login)
	return r
}
