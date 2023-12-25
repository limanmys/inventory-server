package database

import (
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/gorm"
)

var once sync.Once
var connection *gorm.DB

func Connection() *gorm.DB {
	once.Do(func() {
		connection = initialize()

		sqlDB, err := connection.DB()
		if err != nil {
			log.Fatalf("error when getting connection db, ERROR: %s", err.Error())
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})

	return connection
}

func initialize() *gorm.DB {
	switch os.Getenv("DB_DRIVER") {
	case "postgres":
		return initializePostgres()
	case "mysql":
		return initializeMysql()
	case "sqlite":
		return initializeSQLite()
	default:
		log.Fatalln("You must specify a database driver. Choices are 'postgres' / 'mysql' / 'sqlite'")
		return nil
	}
}
