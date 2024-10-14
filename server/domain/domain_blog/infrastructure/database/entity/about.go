package entity

import "gorm.io/gorm"

type About struct {
	gorm.Model
	NameEn string `gorm:"default:'';column:name_en;comment:英文名"`
	NameZh string `gorm:"default:'';column:name_zh;comment:中文名"`
	Value  string `gorm:"default:'';column:value;comment:内容"`
}
