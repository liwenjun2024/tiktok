package models

type Favorite struct {
	Id      int64 `gorm:"column:id;type:int"json:"id,omitempty"`
	UserId  int64 `gorm:"column:user_id;type:int"json:"user_id,omitempty"`
	VideoId int64 `gorm:"column:video_id;type:int"json:"video_id,omitempty"`
	Status  int8  `gorm:"column:status;type:int"json:"status,omitempty"`
}
