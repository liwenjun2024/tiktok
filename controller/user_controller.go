package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"tiktok/models"
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
	//中间进行业务操作
	//start
	//...
	//end
	data := new(models.User)
	//会用到token
	token := "user"
	err := models.DB.Where("username =? and password =?", username, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, models.UserLoginResponse{
				Response: models.Response{StatusCode: 1, StatusMsg: "不存在此账户"},
			})
		}
	} else {
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
	var count int64 = 0
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

			UserName: username,
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
