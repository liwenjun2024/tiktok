package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func GetUserByToken(c *gin.Context, token string) *User {
	objStr := RedisUserInfo.Get(c, token).Val()
	if objStr == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "当前用户尚未登录，请先登录",
		})
		return nil
	}
	b := []byte(objStr)
	user := User{}
	err := json.Unmarshal(b, &user)
	if err != nil {
		log.Println("UserInfo Unmarshal Error:", err)
	}
	return &user
}
