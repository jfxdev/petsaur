package config

import (
	"github.com/spf13/viper"
	"gopkg.in/validator.v2"
)

const (
	DatabaseUserEnv     = "DB_USERNAME"
	DatabasePassEnv     = "DB_PASSWORD"
	DatabaseURLEnv      = "DB_URL"
	DatabasePortEnv     = "DB_PORT"
	DatabaseDatabaseEnv = "DB_DATABASE"
)

type Database struct {
	Username string `mapstructure:"DB_USERNAME" validate:"-"`
	Password string `mapstructure:"DB_PASSWORD" validate:"-"`
	URL      string `mapstructure:"DB_URL"      validate:"-"`
	Port     string `mapstructure:"DB_PORT"     validate:"-"`
	Database string `mapstructure:"DB_DATABASE" validate:"-"`
}

func init() {
	viper.AutomaticEnv()
}

func DatabaseConnection() (result Database, err error) {
	viper.BindEnv(DatabaseUserEnv)
	viper.BindEnv(DatabasePassEnv)
	viper.BindEnv(DatabaseURLEnv)
	viper.BindEnv(DatabasePortEnv)
	viper.BindEnv(DatabaseDatabaseEnv)

	/*
		if viper.GetString("ENV") != "prd" {
			viper.SetConfigFile(".env")
			viper.SetConfigType("env")
			err = viper.ReadInConfig()
			if err != nil {
				return
			}
		}
	*/
	err = viper.Unmarshal(&result)
	if err != nil {
		return
	}

	if err = validator.Validate(result); err != nil {
		return
	}
	return
}
