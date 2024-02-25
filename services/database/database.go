package database

import (
	"log"
	"os"
	"strconv"

	"network/data/configuration"
	"network/data/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Config *configuration.DatabaseConfig

func getDSN() string {
    return "host=" + Config.Host +
        " port=" + strconv.Itoa(int(Config.Port)) +
        " user=" + os.Getenv("DB_USER") +
        " password=" + os.Getenv("DB_PASSWORD") +
        " dbname=" + os.Getenv("DB_NAME") +
        " sslmode=disable TimeZone=UTC"
}

func Initialize() error {
	log.Println("initializing database")

    Config = &configuration.Config.Database

    dsn := getDSN()
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    DB = db
    DB.AutoMigrate(&models.User{})

    if err := RegisterDefaultAdmin(db); err != nil {
        return err
    }

	return nil
}

func Cleanup() error {
	log.Println("closing database connection")
    sqlDB, err := DB.DB()
	
    if err != nil {
        return err
    }

    if sqlDB != nil {
        if err := sqlDB.Close(); err != nil {
            return err
        }
    }

	return nil
}