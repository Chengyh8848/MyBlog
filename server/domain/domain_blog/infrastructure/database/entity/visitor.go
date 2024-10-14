package entity

import "gorm.io/gorm"

type Visitor struct {
	gorm.Model
	Uuid      string `gorm:"default:'';column:uuid;comment:访客标识码"`            // 	访客标识码
	Ip        string `gorm:"default:'';column:ip;comment:ip"`                 // ip
	IpSource  string `gorm:"default:'';column:ip_source;comment:ip来源"`        // ip来源
	Os        string `gorm:"default:'';column:os;comment:操作系统"`               // 操作系统
	Browser   string `gorm:"default:'';column:browser;comment:浏览器"`           // 浏览器
	Pv        int32  `gorm:"default:0;column:pv;comment:访问页数统计"`              // 访问页数统计
	Useragent string `gorm:"default:'';column:user_agent;comment:user_agent"` // user_agent
}
