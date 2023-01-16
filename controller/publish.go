package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/models"
)

var DemoUser = models.User{
	UserId:        1,
	Name:          "Flamingo",
	FansCount:     0,
	FollowerCount: 0,
}

var DemoVideos = []models.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, models.VideoListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
