package persistence

import (
	"context"
	"domain_blog/infrastructure/database"
	"domain_blog/infrastructure/database/entity"
	"gorm.io/gorm"
)

type AboutPersistenceI interface {
	GetList(ctx context.Context) (abouts []entity.About, err error)
	UpdateAbout(ctx context.Context, about *entity.About) (err error)
	GetAboutCommentEnable(ctx context.Context) (enable bool, err error)
}

type AboutPersistence struct {
	model *entity.About
}

func NewAboutPersistence() AboutPersistenceI {
	return &AboutPersistence{model: &entity.About{}}
}

func (a *AboutPersistence) GetList(ctx context.Context) (abouts []entity.About, err error) {
	err = database.Conn.Model(&entity.About{}).Find(&abouts).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return abouts, nil
		}
		return nil, err
	}
	return abouts, nil
}

func (a *AboutPersistence) UpdateAbout(ctx context.Context, about *entity.About) (err error) {
	err = database.Conn.Model(&entity.About{}).Where("id = ?", about.ID).Updates(about).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *AboutPersistence) GetAboutCommentEnable(ctx context.Context) (enable bool, err error) {
	var about entity.About
	err = database.Conn.Where("name_en = commentEnabled").Find(&about).Error
	if err != nil {
		return false, err
	}
	if about.Value == "true" {
		return true, nil
	} else {
		return false, nil
	}
}
