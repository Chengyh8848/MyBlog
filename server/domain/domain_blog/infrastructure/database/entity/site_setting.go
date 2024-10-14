package entity

import "gorm.io/gorm"

type SiteSetting struct {
	gorm.Model
	NameEn string `gorm:"default:'';column:name_en;comment:英文名"`
	NameZh string `gorm:"default:'';column:name_zh;comment:中文名"`
	Value  string `gorm:"default:'';column:value;comment:内容"`
	Type   string `gorm:"default:'';column:type;comment:类型"`
}
