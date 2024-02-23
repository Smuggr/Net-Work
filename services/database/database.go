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


func getDSN(config *configuration.DatabaseConfig) string {
    return "host=" + config.Host +
        " port=" + strconv.Itoa(int(config.Port)) +
        " user=" + os.Getenv("DB_USER") +
        " password=" + os.Getenv("DB_PASSWORD") +
        " dbname=" + os.Getenv("DB_NAME") +
        " sslmode=disable TimeZone=UTC"
}

func Initialize(config *configuration.DatabaseConfig) error {
	log.Println("initializing database")

    dsn := getDSN(config)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    DB = db
    DB.AutoMigrate(&models.User{})

    if err := RegisterUser(db, &models.User{
		Login:           DefaultAdminLogin,
		Username:        DefaultAdminUsername,
		Password:        DefaultAdminPassword,
		PermissionLevel: 1,
	}); err != nil {
		return err
	}

	return nil
}

func Cleanup(config *configuration.DatabaseConfig) error {
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