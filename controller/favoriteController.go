package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tiktok/models"
)

// FavoriteList 用户喜欢列表功能实现
func FavoriteList(c *gin.Context) {
	// 1. 参数获取
	token := c.Query("token")
	//在redis中查询token
	user := models.GetUserByToken(c, token)
	// 2. 获取用户点赞的全部视频信息
	// 2.1 根据用户id 获取点赞的视频id集合    ---> 优化1：利用redis的set结构存储用户点赞视频id
	favoriteList, err := models.FindFavoriteByUserID(user.UserId)
	if err != nil {
		log.Printf("查询出错,err：%#v\n", err)

	}
	ret := new(models.VideoListResponse)
	ret.StatusCode = 0
	ret.StatusMsg = "查询成功"
	ret.VideoList = make([]models.AccordVideo, 0, len(favoriteList))
	for _, favoite := range favoriteList {
		temp, _ := models.GetVideoListWithVideoId(favoite.VideoId)
		res := models.AccordVideo{}
		res.Id = temp.Id
		res.Author, _ = models.GetUserWithId(temp.AuthorId)
		res.PlayUrl = "https://www.w3schools.com/html/movie.mp4"
		res.CoverUrl = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
		// res.PlayUrl = "address" + strconv.FormatInt(temp.Id, 10) + ".mp4"
		// res.CoverUrl = "address" + strconv.FormatInt(temp.Id, 10) + ".png"
		res.FavoriteCount = temp.FavoriteCount
		res.CommentCount = temp.CommentCount
		//待修改
		res.IsFavorite = true
		res.Title = temp.Title
		ret.VideoList = append(ret.VideoList, res)
	}
	c.JSON(http.StatusOK, ret)
}

// Favorite 用户的点赞和取消功能实现
func Favorite(c *gin.Context) {
	// 1. 参数获取
	token := c.Query("token")
	videoId := c.Query("video_id")
	vId, _ := strconv.ParseInt(videoId, 10, 64)
	// 值为1：点赞，值为2取消点赞.
	actionType := c.Query("action_type")
	//在redis中查询token
	user := models.GetUserByToken(c, token)
	favorite := models.Favorite{UserId: user.UserId, VideoId: vId}
	// 2 逻辑处理
	// 2.1 取消点赞
	var err error
	if actionType == "2" {
		// 将status字段修改为0
		err = models.UpdateStatus(favorite, actionType)
		if err != nil {
			c.JSON(http.StatusOK, models.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 0,
			StatusMsg:  "取消点赞",
		})
		return
	} else if actionType == "1" {
		// 2.2 点赞
		// 2.2.1 判断是否该用户之前已经对该视频点过赞。
		exist := models.IsFavorite(favorite)
		if exist {
			// 说明数据已经存在，直接进行修改
			models.UpdateStatus(favorite, actionType)
		} else {
			// 说明数据还未存在，进行创建
			err = models.InsertFavorite(favorite)
			if err != nil {
				log.Println("creating favorite data is Error: ", err)
				c.JSON(http.StatusOK, models.Response{
					StatusCode: 0,
					StatusMsg:  err.Error(),
				})
				return
			}
		}
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		})
		return
	}
}
