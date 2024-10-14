package entity

import "gorm.io/gorm"

type BlogType struct {
	gorm.Model
	Name string `gorm:"default:'';column:name;comment:文章类型"` // 文章类型名称
}
