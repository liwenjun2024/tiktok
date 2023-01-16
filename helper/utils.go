package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt"
)

// Getmd5 md5生成  传入string 返回string
func Getmd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenerateToken 生成Token值
func GenerateToken(mapClaims jwt.MapClaims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(key))
}
