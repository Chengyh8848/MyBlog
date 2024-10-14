package persistence

import (
	"context"
	"domain_blog/infrastructure/database"
	"domain_blog/infrastructure/database/entity"

	"gorm.io/gorm"
)

type UserPersistenceI interface {
	// 创建用户
	Create(ctx context.Context, model *entity.User) (err error)
	// 更新用户
	Update(ctx context.Context, model *entity.User) (err error)
	// 删除用户
	Delete(ctx context.Context, name string) (err error)
	// 查询用户
	GetUserByName(ctx context.Context, name string) (user *entity.User, err error)
}

type UserPersistence struct {
	model *entity.User
}

func NewUserPersistence() UserPersistenceI {
	return &UserPersistence{model: &entity.User{}}
}

func (u *UserPersistence) GetUserByName(ctx context.Context, name string) (user *entity.User, err error) {
	err = database.Conn.Model(&entity.User{}).Where("username = ?", name).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (u *UserPersistence) Update(ctx context.Context, model *entity.User) (err error) {
	err = database.Conn.Model(&entity.User{}).Where("username=?", model.Username).Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserPersistence) Create(ctx context.Context, model *entity.User) (err error) {
	err = database.Conn.Model(&entity.User{}).Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserPersistence) Delete(ctx context.Context, name string) (err error) {
	err = database.Conn.Model(&entity.User{}).Delete(&entity.User{}, "username=?", name).Error
	if err != nil {
		return err
	}
	return nil
}
