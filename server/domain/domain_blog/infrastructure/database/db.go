package database

import (
	"domain_blog/common"
	"domain_blog/conf"
	"domain_blog/infrastructure/database/entity"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

func Init(cfg conf.Database) error {
	var db *gorm.DB
	var err error
	if cfg.Type == conf.Sqlite {
		dbName := fmt.Sprintf("%s.db", cfg.DbName)
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	} else {
		dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.IP, cfg.Port, cfg.DbName)
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,
			SkipInitializeWithVersion: false,
		}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: gormlogger.Default.LogMode(gormlogger.Silent),
		})
	}
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
