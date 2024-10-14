package entity

import "gorm.io/gorm"

type BlogTag struct {
	gorm.Model
	BlogId uint `gorm:"default:0;column:blog_id;comment:博客ID"`
	TagId  uint `gorm:"default:0;column:tag_id;comment:标签ID"`
}
