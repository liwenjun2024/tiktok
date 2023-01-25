package models

type Favorite struct {
	Id      int64 `gorm:"column:id;type:int"json:"id,omitempty"`
	UserId  int64 `gorm:"column:user_id;type:int"json:"user_id,omitempty"`
	VideoId int64 `gorm:"column:video_id;type:int"json:"video_id,omitempty"`
	Status  int8  `gorm:"column:status;type:int"json:"status,omitempty"`
}

// 取消点赞
func UpdateStatus(favorite Favorite, actionType string) error {
	var err error
	if actionType == "2" {
		err = DB.Table("tb_likes").
			Where("user_id=? and video_id=?", favorite.UserId, favorite.VideoId).
			Update("status", "0").Error
	} else {
		err = DB.Table("tb_likes").
			Where("user_id=? and video_id=?", favorite.UserId, favorite.VideoId).
			Update("status", "1").Error
	}

	return err
}

// 用户是否点赞过
func IsFavorite(favorite Favorite) bool {
	var count int64
	DB.Table("tb_likes").
		Where("user_id=? and video_id=?", favorite.UserId, favorite.VideoId).
		Count(&count)
	return count != 0
}

func InsertFavorite(favorite Favorite) error {
	err := DB.Table("tb_likes").Create(favorite).Error
	return err
}

// 查询用户点赞视频
func FindFavoriteByUserID(UserId int64) ([]Favorite, error) {
	var ret []Favorite
	err := DB.Table("tb_likes").Where("user_id=?", UserId).Find(&ret).Error
	return ret, err
}
