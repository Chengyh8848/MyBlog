package database

import (
	"domain_blog/common"
	"domain_blog/conf"
	"domain_blog/infrastructure/database/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Init(cfg conf.Database) error {
	db, err := gorm.Open(sqlite.Open(cfg.DbName), &gorm.Config{})
	if err != nil {
		common.Log.ErrorMsg("failed to connect database %s", err.Error())
		return err
	}
	if cfg.AutoMigrate {
		db.AutoMigrate(
			&entity.User{},
		)
	}
	Conn = db
	return nil
}
