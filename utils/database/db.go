package database

import (
	"fmt"
	"sync"

	"github.com/subashrijal5/go-fiber-boilerplate/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Connect to database
func Connect() error {
	var err error

	once.Do(func() {
		fmt.Println("Opening db connection")
		dsn := config.GetDbConfigString()
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	})
	return err
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	fmt.Println("Closing db connection")
	if db == nil {
		return nil
	}
	dbSQL, err := db.DB()
	if err != nil {
		return err
	}
	return dbSQL.Close()
}
