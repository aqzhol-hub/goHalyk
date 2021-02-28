package migration

import (
	"home/custom"
	"home/models"
)

func MigrateDatabase(cs custom.Custom) {
	conf := cs.GetConfig()
	if conf.Database.Migration {
		db := cs.GetRepository()

		db.AutoMigrate(models.User{})
		db.AutoMigrate(models.UserState{})
	}
}
