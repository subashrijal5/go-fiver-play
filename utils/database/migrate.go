package database

import (
	"github.com/subashrijal5/go-fiber-boilerplate/app/model"
	"github.com/subashrijal5/go-fiber-boilerplate/utils/logger"
)

func Migrate() {
	lgr := logger.GetLogger()
	migrateErr := db.AutoMigrate(&model.User{}, &model.Role{})
	if migrateErr != nil {
		lgr.Panicf("failed database migration. error: %v", migrateErr)
	}
}
