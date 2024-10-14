package entity

import "gorm.io/gorm"

type Moment struct {
	gorm.Model
	Content string `gorm:"default:'';column:content;comment:内容"` //内容
	Likes   int32  `gorm:"default:0;column:likes;comment:点赞数;"`  //点赞数
	Hidden  int    `gorm:"default:0;column:hidden;comment:是否隐藏"` //是否隐藏
}
