package persistence

import (
	"context"
	"domain_blog/infrastructure/database"
	"domain_blog/infrastructure/database/entity"
)

type BlogPersistenceI interface {
	//查询公开博客年月List
	GetGroupYearMonthByIsPublished(ctx context.Context) (values []string, err error)
}

type BlogPersistence struct {
	model *entity.Blog
}

func NewBlogPersistence() BlogPersistenceI {
	return &BlogPersistence{model: &entity.Blog{}}
}

func (b *BlogPersistence) GetGroupYearMonthByIsPublished(ctx context.Context) (values []string, err error) {
	var yerMonths []string
	err = database.Conn.Model(&entity.Blog{}).Select("date_format(created_at, '%Y年%m月') as year_month").Where("published =?", 1).Pluck("year_month", &yerMonths).Group("date_format(created_at, '%Y年%m月')").Order("date_format(created_at, '%Y年%m月') desc").Error
	if err != nil {
		return nil, err
	}
	return yerMonths, nil
}
