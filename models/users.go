package models

type User struct {
	UserId   uint   `gorm:"column:user_id;AUTO_INCREMENT;primary_key"`
	UserName string `gorm:"column:user_name;not null;unique;type:varchar(10)"`
	Password string `gorm:"column:password;not null;varchar(10)"`
}
