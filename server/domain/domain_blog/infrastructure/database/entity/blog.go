package entity

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title          string `gorm:"default:'';column:title;comment:标题"`            //文章标题
	FirstPicture   string `gorm:"default:'';column:first_picture;comment:首图"`    //文章首图，用于随机文章展示
	Content        string `gorm:"default:'';column:content;comment:内容"`          //文章内容
	Description    string `gorm:"default:'';column:description;comment:描述"`      //文章描述
	Published      int    `gorm:"default:0;column:published;comment:是否公开"`       //是否公开
	Recommend      int    `gorm:"default:0;column:recommend;comment:是否推荐"`       //是否推荐
	Appreciation   int    `gorm:"default:0;column:appreciation;comment:是否赞赏"`    //是否赞赏
	CommentEnabled int    `gorm:"default:0;column:comment_enabled;comment:评论功能"` //评论功能是否开启
	Top            int    `gorm:"default:0;column:top;comment:置顶"`               //是否置顶
	Views          int32  `gorm:"default:0;column:views;comment:浏览量"`            //浏览量
	Words          int32  `gorm:"default:0;column:words;comment:字数"`             //字数
	ReadTime       int32  `gorm:"default:0;column:read_time;comment:阅读时间"`       //阅读时间
	Password       string `gorm:"default:'';column:password;comment:密码"`         //文章密码
	BlogType       int32  `gorm:"default:0;column:blog_type;comment:博客类型"`       //博客类型
}
