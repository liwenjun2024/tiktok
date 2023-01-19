package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"tiktok/define"
	"tiktok/helper"
	"tiktok/models"
	"time"
)

// Login
// @Tags 公共方法
// @Summary 用户登录
// @Param username query string false "username"
// @Param password query string false "password"
// @Success 200 {string} json "{"Response":{},"UserId","","Token":""}"
// @Router /douyin/user/login/ [post]
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//进行业务操作
	data := models.User{}
	var nowDate = time.Now().Format(define.DataTime)
	var secret = fmt.Sprintf("%v%v", nowDate, "xxxx")
	user := make(map[string]interface{})
	user["name"] = username
	user["password"] = password
	token, err := helper.GenerateToken(user, secret)
	if err != nil {
		log.Println("token create Error: ", err)
		return
	}
	//查询是否存在此用户
	err = models.DB.Where("username =? and password =?", username, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, models.UserLoginResponse{
				Response: models.Response{StatusCode: 1, StatusMsg: "不存在此账户"},
			})
		}
	} else {
		//将当前登录用户存进redis
		userData, err := json.Marshal(data)
		if err != nil {
			log.Println("json.Marshal Error: ", err)
			return
		}
		// redis userInfo 过期时间
		err = models.RedisUserInfo.Set(c, token, string(userData), 6*time.Hour).Err()
		if err != nil {
			log.Println("redis set Error:", err)
			return
		}

		c.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 0},
			UserId:   data.UserId,
			Token:    token,
		})
	}

}

// Register
// @Tags 公共方法
// @Summary 用户注册
// @Param username query string false "username"
// @Param password query string false "password"
// @Success 200 {string} json "{"StatusCode","","StatusMsg":""}"
// @Router /douyin/user/register/ [post]
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(password) <= 5 {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "用户密码长度必须大于5位",
		})
	}
	var count int64
	err := models.DB.Table("tb_user").Where("username = ?", username).Count(&count).Error
	if err != nil {
		log.Println("Register select is Error: ", err)
		return
	}
	if count >= 1 {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "用户名已被注册",
		})
	} else {
		user := &models.User{
			Name:     username,
			Password: password,
		}
		err := models.DB.Create(user).Error
		if err != nil {
			log.Println("Register create is Error: ", err)
			return
		}
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "用户注册成功",
		})
	}
}

// UserInfo
// @Tags 公共方法
// @Summary 当前用户信息
// @Param token query string false "token"
// @Success 200 {string} json "{"Response":{},"User",{}}"
// @Router /douyin/user/ [get]
func UserInfo(c *gin.Context) {
	token := c.Query("token")
	//在redis中查询token
	objStr := models.RedisUserInfo.Get(c, token).Val()
	if objStr == "" {
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
	c.JSON(http.StatusOK, models.UserResponse{
		Response: models.Response{StatusCode: 0},
		User:     user,
	})

}
