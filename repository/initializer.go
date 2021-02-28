package repository

import (
	"errors"
	"fmt"
	"home/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitRepository(conf *config.Config) (error, Repository) {
	db, err := getConnection(conf)

	if err != nil {
		return errors.New("can not connect to db"), nil
	}

	return nil, &repository{db: db}
}

func getConnection(conf *config.Config) (*gorm.DB, error) {
	fmt.Println(conf.Database.Dialect)
	if conf.Database.Dialect == POSTGRES {
		connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.Username, conf.Database.Dbname, conf.Database.Password)
		driver := postgres.Open(connectionString)
		return gorm.Open(driver, &gorm.Config{})

	} else if conf.Database.Dialect == MYSQL {
		connectionString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.Database.Username, conf.Database.Password, conf.Database.Host, conf.Database.Dbname)
		driver := mysql.Open(connectionString)
		return gorm.Open(driver, &gorm.Config{})

	} else {
		connectionString := fmt.Sprintf("%s", conf.Database.Dialect)
		driver := sqlite.Open(connectionString)
		return gorm.Open(driver, &gorm.Config{})
	}

}
