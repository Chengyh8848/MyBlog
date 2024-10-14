package entity

import "gorm.io/gorm"

type VisitLog struct {
	gorm.Model
	Uuid      string `gorm:"default:'';column:uuid;comment:访客标识码"`            // 	访客标识码
	Uri       string `gorm:"default:'';column:uri;comment:请求接口"`              // 请求接口
	Method    string `gorm:"default:'';column:method;comment:请求方式"`           // 请求方式
	Param     string `gorm:"default:'';column:param;comment:请求参数"`            // 请求参数
	Behavior  string `gorm:"default:'';column:behavior;comment:访问行为"`         // 访问行为
	Content   string `gorm:"default:'';column:content;comment:访问内容"`          // 访问内容
	Remark    string `gorm:"default:'';column:remark;comment:备注"`             // 备注
	Ip        string `gorm:"default:'';column:ip;comment:ip"`                 // ip
	IpSource  string `gorm:"default:'';column:ip_source;comment:ip来源"`        // ip来源
	Os        string `gorm:"default:'';column:os;comment:操作系统"`               // 操作系统
	Browser   string `gorm:"default:'';column:browser;comment:浏览器"`           // 浏览器
	Times     int32  `gorm:"default:0;column:times;comment:请求耗时（毫秒）"`         // 请求耗时（毫秒）
	UserAgent string `gorm:"default:'';column:user_agent;comment:user_agent"` // user_agent
}
