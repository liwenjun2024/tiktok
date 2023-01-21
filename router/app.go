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

	appRouter := r.Group("/douyin")
	//用户
	appRouter.POST("/user/login/", controller.Login)       //用户登录
	appRouter.POST("/user/register/", controller.Register) //用户注册
	appRouter.POST("/publish/action/", controller.Publish) //用户发布视频
	appRouter.GET("/user/", controller.UserInfo)           //用户个人信息

	//作品
	appRouter.GET("/publish/list/", controller.PublishList)   //公开作品
	appRouter.GET("/favorite/list/", controller.FavoriteList) //个人喜欢作品
	appRouter.GET("/feed/", controller.Feed)                  //视频流

	appRouter.POST("/favorite/action/", controller.Favorite)

	return r
}
