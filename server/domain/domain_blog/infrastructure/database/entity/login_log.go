package entity

import "gorm.io/gorm"

type LoginLog struct {
	gorm.Model
	Username  string `gorm:"default:'';column:username;comment:用户名"`         //用户名
	Ip        string `gorm:"default:'';column:ip;comment:登录IP"`              //登录IP
	Status    int    `gorm:"default:0;column:status;comment:登录状态"`           //登录状态
	IpSource  string `gorm:"default:'';column:ip_source;comment:IPSource"`   // IP来源
	Os        string `gorm:"default:'';column:os;comment:OS"`                // 操作系统
	Browser   string `gorm:"default:'';column:browser;comment:浏览器"`          // 浏览器
	Msg       string `gorm:"default:'';column:msg;comment:登录信息"`             // 登录信息
	UserAgent string `gorm:"default:'';column:user_agent;comment:UserAgent"` // 用户代理
}
