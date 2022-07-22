package db

import (
	"fmt"

	"github.com/jfxdev/petsaur/src/config"
	"github.com/jfxdev/petsaur/src/models/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func Setup() (err error) {
	cfg, err := config.DatabaseConnection()
	if err != nil {
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.URL, cfg.Port, cfg.Database)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = migrations.New(db).Run()
	if err != nil {
		return
	}

	command, err := db.DB()
	if err != nil {
		return err
	}

	return command.Ping()
}
