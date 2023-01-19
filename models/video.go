package models

// Video 添加视频操作结构体
type Video struct {
	Id            int64  `gorm:"column:id;"json:"id,omitempty"`
	AuthorId      int64  `gorm:"column:author_id;"json:"author_id"`
	Title         string `gorm:"column:title;"json:"title"`
	PlayUrl       string `gorm:"column:play_url;"json:"play_url,omitempty"`
	CoverUrl      string `gorm:"column:cover_url;"json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"column:favorite_count;"json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"column:comment_count;"json:"comment_count,omitempty"`
	CreateTime    string `gorm:"column:create_time;"json:"create_time"`
}

// AccordVideo Feed流显示视频结构体
type AccordVideo struct {
	Id            int64  `gorm:"column:id;"json:"id,omitempty"`
	Author        User   `json:"author"`
	AuthorId      int64  `gorm:"column:author_id;"json:"author_id"`
	Title         string `gorm:"column:title;"json:"title"`
	PlayUrl       string `gorm:"column:play_url;"json:"play_url,omitempty"`
	CoverUrl      string `gorm:"column:cover_url;"json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"column:favorite_count;"json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"column:comment_count;"json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	CreateTime    string `gorm:"column:create_time;"json:"create_time"`
}

type VideoListResponse struct {
	Response
	VideoList []AccordVideo `json:"video_list"`
}

type VideoFeedResponse struct {
	Response
	VideoList []AccordVideo `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

func (video *Video) TableName() string {
	return "tb_video"
}
