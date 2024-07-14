package config

import (
	"os"
	"strconv"
)

type DbConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	SslMode      string
	DatabaseName string
	TimeZone     string
}

var dbConfig = &DbConfig{}

func GetDbConfig() *DbConfig {
	return dbConfig
}

func GetDbConfigString() string {
	return "host=" + dbConfig.Host + " port=" + strconv.Itoa(dbConfig.Port) + " user=" + dbConfig.Username + " password=" + dbConfig.Password + " dbname=" + dbConfig.DatabaseName
}

// ConfigureDb set the db config
func ConfigureDb() {
	dbConfig.Host = os.Getenv("DB_HOST")
	dbConfig.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	dbConfig.Username = os.Getenv("DB_USERNAME")
	dbConfig.Password = os.Getenv("DB_PASSWORD")
	dbConfig.SslMode = os.Getenv("DB_SSLMODE")
	dbConfig.DatabaseName = os.Getenv("DB_DATABASE")
	dbConfig.TimeZone = os.Getenv("TIMEZONE")
}
