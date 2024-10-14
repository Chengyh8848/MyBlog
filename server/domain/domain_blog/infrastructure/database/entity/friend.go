package entity

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	Nickname    string `gorm:"default:'';column:nick_name;comment:昵称"`   // 昵称
	Description string `gorm:"default:'';column:description;comment:描述"` // 描述
	Website     string `gorm:"default:'';column:website;comment:站点"`     //站点
	Avatar      string `gorm:"default:'';column:avatar;comment:头像"`      //头像
	Hidden      int    `gorm:"default:0;column:hidden;comment:是否隐藏"`     //是否隐藏
	Views       int32  `gorm:"default:0;column:views;comment:浏览量"`       //浏览量
}
