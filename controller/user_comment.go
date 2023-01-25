package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"tiktok/models"
	"time"
)

func CommentAction(c *gin.Context) {
	token := c.Query("token")
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
	b := []byte(objStr)
	user := models.User{}
	err := json.Unmarshal(b, &user)
	if err != nil {
		log.Println("UserInfo Unmarshal Error:", err)
		return
	}

	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		log.Println("VideoId Error:", err)
		return
	}

	actionType := c.Query("action_type")
	switch actionType {
	case "1":
		// 发布评论
		commentText := c.Query("comment_text")
		now := time.Now().Format("01-02")
		comment := models.Comment{
			UserId:      user.UserId,
			VideoId:     videoId,
			CommentText: commentText,
			CreateDate:  now,
		}
		err := models.DB.Create(&comment).Error
		if err != nil {
			log.Println("Comment is Error: ", err)
			c.JSON(http.StatusOK, models.Response{
				StatusCode: 1,
				StatusMsg:  "评论失败",
			})
			return
		}
		comment.User = user
		c.JSON(http.StatusOK, models.CommentResponse{
			Response: models.Response{
				StatusCode: 0,
				StatusMsg:  "评论成功",
			},
			Comment: comment,
		})
		if models.RedisCommentsInfo.Get(c, string(videoId)).Val() != "" {
			models.RedisCommentsInfo.Del(c, string(videoId))
		}

	case "2":
		// 删除评论
		commentId := c.Query("comment_id")
		comment := &models.Comment{}
		err := models.DB.Where("id = ?", commentId).First(comment).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusOK, models.Response{
					StatusCode: 1,
					StatusMsg:  "不存在此评论",
				})
				return
			}
		}
		comment.Status = 0
		err = models.DB.Save(comment).Error
		if err != nil {
			c.JSON(http.StatusOK, models.Response{
				StatusCode: 1,
				StatusMsg:  "删除失败",
			})
			return
		}
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 0,
			StatusMsg:  "删除成功",
		})
		if models.RedisCommentsInfo.Get(c, string(videoId)).Val() != "" {
			models.RedisCommentsInfo.Del(c, string(videoId))
		}
	default:
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "参数错误",
		})
	}
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	//在redis中查询token
	objStr := models.RedisUserInfo.Get(c, token).Val()
	if objStr == "" {
		log.Println("redis get token Error: ")
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "当前用户不合法",
		})
		return
	}
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		log.Println("VideoId Error:", err)
		return
	}

	// 将当前视频评论序列化从rides中取出来
	commentsData := models.RedisCommentsInfo.Get(c, string(videoId)).Val()
	if commentsData != "" {
		var comments = make([]models.Comment, 0)
		data := []byte(commentsData)
		err := json.Unmarshal(data, &comments)
		if err != nil {
			log.Println("json Unmarshal Error: ", err)
			return
		}
		c.JSON(http.StatusOK, models.CommentListResponse{
			Response: models.Response{StatusCode: 0},
			Comments: comments,
		})
		return
	}

	// redis中无评论缓存
	b := []byte(objStr)
	user := models.User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Println("UserInfo Unmarshal Error:", err)
		return
	}

	var comments = make([]models.Comment, 0)
	err = models.DB.
		Where("video_id = ? and status = ?", videoId, 1).
		Order("id desc").Find(&comments).Error
	if err != nil {
		log.Println("Video select is Error:", err)
		return
	}

	for i, value := range comments {
		var u = models.User{}
		e := models.DB.Where("user_id = ?", value.UserId).First(&u).Error
		if e != nil {
			log.Println("User select is Error:", err)
			return
		}
		comments[i].User = u
	}

	commentJson, _ := json.Marshal(comments)
	err = models.RedisCommentsInfo.Set(c, string(videoId), string(commentJson), 6*time.Minute).Err()
	if err != nil {
		log.Println("redis set Error:", err)
		return
	}

	c.JSON(http.StatusOK, models.CommentListResponse{
		Response: models.Response{StatusCode: 0},
		Comments: comments,
	})
}
