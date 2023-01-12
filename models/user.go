package models

type User struct {
	UserId        int64  `gorm:"column:Id;type:int"json:"user_id,omitempty"`
	UserName      string `gorm:"column:username;type:varchar(20)"json:"username"`
	Password      string `gorm:"column:password;type:varchar(50)"json:"password"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// TableName 返回数据库 表名称
func (table *User) TableName() string {
	return "user"
}
