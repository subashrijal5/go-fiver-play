package database

import (
	"github.com/subashrijal5/go-fiber-boilerplate/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db = new(gorm.DB)

// Connect to database
func Connect() (db *gorm.DB, err error) {
	dsn := config.GetDbConfigString()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	dbSQL, err := db.DB()
	if err != nil {
		return err
	}
	return dbSQL.Close()
}
