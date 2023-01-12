package helper

import (
	"crypto/md5"
	"fmt"
)

// Getmd5 md5生成  传入string 返回string
func Getmd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
