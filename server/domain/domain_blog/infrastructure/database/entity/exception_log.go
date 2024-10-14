package entity

import "gorm.io/gorm"

type exceptionLog struct {
	gorm.Model
	Url         string `gorm:"default:'';column:url;comment:请求"`               // 请求
	Method      string `gorm:"default:'';column:method;comment:请求方法"`          // 请求方法
	Param       string `gorm:"default:'';column:param;comment:请求参数"`           // 请求参数
	Description string `gorm:"default:'';column:description;comment:描述"`       // 描述
	Error       string `gorm:"default:'';column:error;comment:Error"`          // 异常信息
	Ip          string `gorm:"default:'';column:ip;comment:IP"`                // IP
	IpSource    string `gorm:"default:'';column:ip_source;comment:IPSource"`   // IP来源
	Os          string `gorm:"default:'';column:os;comment:OS"`                // 操作系统
	Browser     string `gorm:"default:'';column:browser;comment:浏览器"`          // 浏览器
	UserAgent   string `gorm:"default:'';column:user_agent;comment:UserAgent"` // 用户代理
}
