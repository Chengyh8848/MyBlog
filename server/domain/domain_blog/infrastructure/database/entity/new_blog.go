package entity

import "gorm.io/gorm"

type NewBlog struct {
	gorm.Model
	Title    string `gorm:"default:'';column:title;comment:标题"`    //标题
	Password string `gorm:"default:'';column:password;comment:密码"` //密码
	Privacy  int    `gorm:"default:0;column:privacy;comment:隐私"`   //隐私
}
