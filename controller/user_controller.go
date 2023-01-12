package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
