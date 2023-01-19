package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"tiktok/define"
	"tiktok/models"
	"tiktok/service"
	"time"
)

var DemoUser = models.User{
	UserId:        1,
	Name:          "Flamingo",
	FansCount:     0,
	FollowerCount: 0,
}

var DemoVideos = []models.AccordVideo{
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

// Publish 实现视频文件上传
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	// 鉴权
	objStr := models.RedisUserInfo.Get(c, token).Val()
	if objStr == "" {
		log.Println("redis get token Error: ")
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "当前用户不合法",
		})
		return
	}

	data, err := c.FormFile("data")
	file, err1 := data.Open()
	if err1 != nil {
		log.Println("data Open Error: ", err)
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if err != nil {
		log.Println("formFile Error: ", err)
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	title := c.PostForm("title")
	//取出的token值，转换为对象
	user := &models.User{}
	b := []byte(objStr)
	err = json.Unmarshal(b, user)
	if err != nil {
		log.Println("json Unmarshal Error: ", err)
		return
	}
	uid := uuid.NewV4().String()
	finalName := fmt.Sprintf("%s_%s", user.Name, uid) //生成视频名称
	err1 = service.Publish(file, finalName)           //进行视频保存操作
	//缺一个进行图片保存操作

	if err1 != nil {
		log.Println("视频保存失败 Error: ", err)
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "上传文件失败",
		})
		return
	}
	defer file.Close()
	//根据当前用户id 添加数据进入数据库
	createTime := time.Now().Format(define.DataTime)
	video := &models.Video{
		AuthorId:      user.UserId,
		Title:         title,
		PlayUrl:       define.PlayUrlAddress + finalName + ".mp4",
		CoverUrl:      define.CoverUrlAddress + finalName + ".jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		CreateTime:    createTime,
	}
	err = models.DB.Save(&video).Error
	if err != nil {
		log.Println("video create Error: ", err)
		return
	}
	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  "上传成功",
	})
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
