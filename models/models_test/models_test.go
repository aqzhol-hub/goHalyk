package models_test

import (
	"fmt"
	"home/config"
	"home/models"
	"home/repository"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db repository.Repository

func TestConnection(t *testing.T) {
	con := &config.Config{
		Database: struct {
			Dialect   string `default:"sqlite3"`
			Host      string `default:"database.db"`
			Port      string `default:""`
			Dbname    string `default:""`
			Username  string `default:""`
			Password  string `default:""`
			Migration bool   `default:"true"`
		}{
			Dialect:   "postgres",
			Host:      "rogue.db.elephantsql.com",
			Port:      "5432",
			Dbname:    "meqgahia",
			Username:  "meqgahia",
			Password:  "1rXp4q5laC2QIKemZlOzvZPiG1fdDASG",
			Migration: true,
		},
	}

	err, database := repository.InitRepository(con)
	assert.NoError(t, err)

	db = database
}

func TestHashPassword(t *testing.T) {
	user := models.User{
		Password: "secret",
	}
	err := user.HashPassword()
	assert.NoError(t, err)

	fmt.Println(db, user.Password)
	os.Setenv("passwordHash", user.Password)
}

func TestCreateUserRecord(t *testing.T) {

	err := db.AutoMigrate(&models.UserState{})
	assert.NoError(t, err)

	state := models.UserState{
		ID:          1,
		Name:        "online",
		Description: "desc",
	}

	err = state.CreateUserStateRecord(db)
	assert.NoError(t, err)

	state = models.UserState{
		ID:          2,
		Name:        "online",
		Description: "desc",
	}
	err = state.CreateUserStateRecord(db)
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.User{})
	assert.NoError(t, err)

	user := models.User{
		ID:       10001,
		Username: "aqzhol",
		Password: "123",
		State:    &state,
	}
	err = user.CreateUserRecord(db)
	assert.NoError(t, err)

	// var searchResult models.User
	res := user.FindByUsername(db, user.Username)
	assert.NotNil(t, res)

	res = user.FindByID(db, 10001)
	fmt.Println(res, " ------- ")
	assert.NotNil(t, res)

	err = user.UpdateState(db, 1)
	assert.NoError(t, err)

	db.Delete(&res)
	db.Delete(&state)
}
