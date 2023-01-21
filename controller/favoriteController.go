package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tiktok/models"
)

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	// 1. 参数获取
	token := c.Query("token")
	// 2. 校验token
	err := models.RedisUserInfo.Get(c, token).Err()
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "当前用户未登录, 请先登录",
		})
		return
	}
	// 3. 根据token获取用户信息
	objStr := models.RedisUserInfo.Get(c, token).Val()
	b := []byte(objStr)
	user := models.User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Println("UserInfo Unmarshal Error:", err)
		return
	}

	c.JSON(http.StatusOK, models.VideoListResponse{
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "查询成功",
		},
		VideoList: DemoVideos,
	})
}

func Favorite(c *gin.Context) {
	// 1. 参数获取
	token := c.Query("token")
	videoId := c.Query("video_id")
	vId, _ := strconv.ParseInt(videoId, 10, 64)
	// 值为1：点赞，值为2取消点赞.
	actionType := c.Query("action_type")
	//在redis中查询token
	objStr := models.RedisUserInfo.Get(c, token).Val()
	if objStr == "" {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "当前用户尚未登录，请先登录",
		})
		return
	}
	b := []byte(objStr)
	user := models.User{}
	err := json.Unmarshal(b, &user)
	if err != nil {
		log.Println("UserInfo Unmarshal Error:", err)
		return
	}
	// 2 逻辑处理
	// 2.1 取消点赞
	if actionType == "2" {
		// 将status字段修改为0
		models.DB.Table("tb_likes").
			Where("video_id = ? AND user_id = ?", vId, user.UserId).
			Update("status", "0")
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 0,
			StatusMsg:  "取消点赞",
		})
		return
	}
	// 2.2 点赞
	if actionType == "1" {
		// 2.2.1 判断是否该用户之前已经对该视频点过赞。
		var count int64
		// 此处应该建立联合索引
		models.DB.Table("tb_likes").
			Where("video_id = ? AND user_id = ?", vId, user.UserId).
			Count(&count)
		if count == 1 {
			// 说明数据已经存在，直接进行修改
			models.DB.Table("tb_likes").
				Where("video_id = ? AND user_id = ?", vId, user.UserId).
				Update("status", "1")
		} else {
			// 说明数据还未存在，进行创建
			favorite := &models.Favorite{
				VideoId: vId,
				UserId:  user.UserId,
				Status:  1,
			}
			err = models.DB.Table("tb_likes").Create(favorite).Error
			if err != nil {
				log.Println("creating favorite data is Error: ", err)
			}
		}
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		})
		return
	}
}
