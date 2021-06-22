package db

import (
	"database/sql"
	"fmt"
	"gin-sample/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type Mysql struct {
	db    *gorm.DB
	sqlDb *sql.DB
}

var instance *Mysql

func InitMysql(settings config.Database) {
	cfg := gorm.Config{
		QueryFields: true,
	}
	if settings.Debug {
		cfg.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(settings.URI), &cfg)
	if err != nil {
		log.Fatal("Failed to connect sql")
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal("Failed to connect sql")
	}
	sqlDb.SetMaxIdleConns(settings.MaxIdleConns)
	sqlDb.SetConnMaxLifetime(time.Duration(settings.ConnMaxLifeTime) * time.Millisecond)
	sqlDb.SetMaxOpenConns(settings.MaxOpenConns)
	instance = &Mysql{
		db:    db,
		sqlDb: sqlDb,
	}
}

func GetMysql() (db *gorm.DB) {
	if instance == nil {
		return
	}
	return instance.db
}

func CloseMysql() {
	err := instance.sqlDb.Close()
	if err != nil {
		fmt.Printf("Error close mysql %s", err.Error())
	}
}
