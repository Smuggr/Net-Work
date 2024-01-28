package database

import (
	"os"
	"log"

    "overseer/services/models"

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

func getDefaultUser(db *gorm.DB) (models.User, error) {
    var defaultUser models.User
    result := db.First(&defaultUser, "username = ?", os.Getenv("ADMIN_USERNAME"))

    return defaultUser, result.Error
}

func createDefaultUser(db *gorm.DB) {    
    defaultUser, err := getDefaultUser(db)
    if err == nil {
        log.Println("Default user already exists.")
        return
    }

    defaultUser = models.User{
        Username:        os.Getenv("ADMIN_USERNAME"),
        Password:        os.Getenv("ADMIN_PASSWORD"),
        Identifier:      1,
        PermissionLevel: 1,
    }

    result := db.Create(&defaultUser)
    if result.Error == nil {
        log.Println("Default user created successfully.")
    } else {
        log.Fatalf("Error creating default user: %v", result.Error)
    }
}

func Initialize() {
	log.Println("Initializing database")

    dsn := getDSN()
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    DB = db
    DB.AutoMigrate(&models.User{})

    createDefaultUser(DB)
}