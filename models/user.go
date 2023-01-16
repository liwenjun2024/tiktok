package models

type User struct {
	UserId        int64  `gorm:"column:user_id;type:int"json:"user_id,omitempty"`
	Name          string `gorm:"column:username;type:varchar(32)"json:"name"`
	Password      string `gorm:"column:password;type:varchar(32)"json:"password"`
	FansCount     int64  `gorm:"column:fans;type:int"json:"fans,omitempty"`
	FollowerCount int64  `gorm:"column:followee;type:int"json:"followee,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// TableName 返回数据库 表名称
func (table *User) TableName() string {
	return "tb_user"
}
