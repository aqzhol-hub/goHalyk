package config

import (
	"errors"

	"github.com/jinzhu/configor"
)

type Config struct {
	Database struct {
		Dialect   string `default:"sqlite3"`
		Host      string `default:"database.db"`
		Port      string `default:""`
		Dbname    string `default:""`
		Username  string `default:""`
		Password  string `default:""`
		Migration bool   `default:"true"`
	}
	Redis struct {
		Network  string `default:"tcp"`
		Addr     string `default:"localhost:6379"`
		Password string `default:""`
		DB       int    `default:"0"`
	}
}

func Load() (error, *Config) {
	config := &Config{}
	if err := configor.Load(config, "../config.yml"); err != nil {
		return errors.New("Failed to read config.yml"), nil
	}
	return nil, config
}
