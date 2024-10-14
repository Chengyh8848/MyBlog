package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"default:'';column:username;comment:用户名"` //用户名
	Password string `gorm:"default:'';column:password;comment:密码"`  //密码
	Nickname string `gorm:"default:'';column:nickname;comment:昵称"`  //昵称
	Avatar   string `gorm:"default:'';column:avatar;comment:头像"`    //头像
	Email    string `gorm:"default:'';column:email;comment:邮箱"`     //邮箱
	Role     string `gorm:"default:'';column:role;comment:角色"`      //角色
	Salt     string `gorm:"default:'';column:salt"`
}

func (*User) TableName() string {
	return "user"
}
