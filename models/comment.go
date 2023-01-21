package models

type Comment struct {
	Id          int64  `gorm:"column:id;type:int"json:"id,omitempty"`
	UserId      int64  `gorm:"column:user_id;type:int"json:"-"`
	User        User   `gorm:"-"json:"user,omitempty"`
	VideoId     int64  `gorm:"column:video_id;type:int"json:"-"`
	CommentText string `gorm:"column:comment_text;type:varchar(256)"json:"content,omitempty"`
	CreateDate  string `gorm:"column:create_date;type:varchar(20)"json:"create_date,omitempty"`
	Status      int    `gorm:"column:status;type:tinyint(1);default:1"json:"-"`
}

type CommentResponse struct {
	Response
	Comment Comment `json:"comment"`
}

type CommentListResponse struct {
	Response
	Comments []Comment `json:"comment_list"`
}

// TableName 返回数据库 表名称
func (table *Comment) TableName() string {
	return "tb_comments"
}
