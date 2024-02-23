package database

import (
	"log"
	"os"
	"strconv"

	"network/data/configuration"
	"network/data/errors"
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

func createDefaultUser(db *gorm.DB) {
	err := RegisterUser(db, &models.User{
		Login:           DefaultAdminLogin,
		Username:        DefaultAdminUsername,
		Password:        DefaultAdminPassword,
		PermissionLevel: 1,
	})

	if err != nil {
		log.Println(err)
	}
}


func Initialize(config *configuration.DatabaseConfig) {
	log.Println("initializing database")

    dsn := getDSN(config)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    DB = db
    DB.AutoMigrate(&models.User{})

    createDefaultUser(DB)
}

func Cleanup(config *configuration.DatabaseConfig) *errors.ErrorWrapper{
	log.Println("closing database connection")
    sqlDB, err := DB.DB()
	
    if err != nil {
        return errors.ErrGettingDBConnection
    }

    if sqlDB != nil {
        if err := sqlDB.Close(); err != nil {
            return errors.ErrClosingDBConnection
        }
    }

	return nil
}