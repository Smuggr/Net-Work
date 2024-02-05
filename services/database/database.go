package database

import (
	"log"
	"os"

	"overseer/data/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func getDSN() string {
    return "host=" + os.Getenv("DB_HOST") +
        " port=" + os.Getenv("DB_PORT") +
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


func Initialize() {
	log.Println("initializing database")

    dsn := getDSN()
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    DB = db
    DB.AutoMigrate(&models.User{})

    createDefaultUser(DB)
}