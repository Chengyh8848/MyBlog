package entity

import "gorm.io/gorm"

type VisitRecord struct {
	gorm.Model
	Pv   int32  `gorm:"default:0;column:pv;comment:访问页数统计"` // 访问页数统计
	Uv   int32  `gorm:"default:0;column:uv;comment:独立访客统计"`
	Date string `gorm:"column:date;comment:访问日期"` // 访问日期
}
