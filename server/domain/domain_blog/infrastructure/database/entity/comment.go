package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Nickname      string `gorm:"default:'';column:nick_name;comment:昵称"` // 昵称
	Email         string `gorm:"default:'';column:email;comment:邮箱"`     // 邮箱
	Content       string `gorm:"default:'';column:content;comment:内容"`   // 内容
	Avatar        string `gorm:"default:'';column:avatar;comment:头像"`    //头像
	Website       string `gorm:"default:'';column:website;comment:个人网站"` //个人网站
	Ip            string `gorm:"default:'';column:ip;comment:IP"`
	Hidden        int    `gorm:"default:0;column:hidden;comment:是否隐藏"`           //是否隐藏
	Type          int    `gorm:"default:0;column:type;comment:评论类型"`             //评论类型 1-文章评论 2-回复评论
	ReplyNickName string `gorm:"default:'';column:reply_nick_name;comment:回复昵称"` //回复昵称
	ReplyId       int64  `gorm:"default:0;column:reply_id;comment:回复ID"`         //回复ID
	BlogId        int64  `gorm:"default:0;column:blog_id;comment:博客ID"`          //博客ID
}
