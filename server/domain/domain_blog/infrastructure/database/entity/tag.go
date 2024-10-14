package entity

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string `gorm:"default:'';column:name;comment:标签名称"`
	Color string `gorm:"default:'';column:color;comment:标签颜色"`
}
