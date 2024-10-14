package entity

import "gorm.io/gorm"

type CityVisitor struct {
	gorm.Model
	City  string `gorm:"default:'';column:city;comment:城市"`  // 城市
	Value int32  `gorm:"default:0;column:value;comment:访客数"` // 访客数
}
