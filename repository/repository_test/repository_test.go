package repository_test

import (
	"home/config"
	"home/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRepository(t *testing.T) {
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
			Dbname:    "zhbvfwki",
			Username:  "zhbvfwki",
			Password:  "9CjDtthVFj3-qBgzxfneKffMAZ-mcqCG",
			Migration: true,
		},
	}

	err, _ := repository.InitRepository(con)
	assert.NoError(t, err)
}
