package service

import (
	"io"
	"log"
	"tiktok/models"
)

// Publish
// 保存进ftp
func Publish(file io.Reader, videoName string) error {
	err := models.MyFTP.Cwd("/video")
	if err != nil {
		log.Println("FTP Cwd Error: ", err)
		return err
	}
	models.MyFTP.Stor(videoName+".mp4", file)
	if err != nil {
		log.Println("FTP Stor Error: ", err)
		return err
	}
	log.Println("视频上传成功")

	return nil
}

//图片上传方法
