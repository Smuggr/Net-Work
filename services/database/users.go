package database

import (
	"os"
	"log"

	"overseer/services/models"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)


func getDefaultUser(db *gorm.DB) (models.User, error) {
    var defaultUser models.User
    result := db.First(&defaultUser, "id = ?", 1)

	if result.Error != nil {
		return models.User{}, result.Error
	}

    return defaultUser, result.Error
}

func createDefaultUser(db *gorm.DB) {
	defaultUser, err := getDefaultUser(db)

	if err == nil {
		log.Println("Default user already exists.")

		adminLogin := os.Getenv("ADMIN_LOGIN")

		var existingUser models.User
		result := db.Where("login = ? AND id != ?", adminLogin, 1).First(&existingUser)
		if result.Error == nil {
			log.Fatalf("Error updating default user, user of this login already exists: %v", result.Error)
		}

		newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Error hashing password: %v", err)
		}

		defaultUser.Login = adminLogin
		defaultUser.Username = os.Getenv("ADMIN_USERNAME")
		defaultUser.Password = string(newHashedPassword)

		result = db.Updates(&defaultUser)
		if result.Error == nil {
			log.Println("Default user updated successfully")
		} else {
			log.Fatalf("Error updating default user: %v", result.Error)
		}

		return
	} else {
		log.Fatalf("Error checking default user existence: %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	defaultUser = models.User{
		Login:           os.Getenv("ADMIN_LOGIN"),
		Username:        os.Getenv("ADMIN_USERNAME"),
		Password:        string(hashedPassword),
		PermissionLevel: 1,
	}

	result := db.Create(&defaultUser)
	if result.Error == nil {
		log.Println("Default user created successfully.")
	} else {
		log.Fatalf("Error creating default user: %v", result.Error)
	}
}

func RegisterUser(db *gorm.DB, newUser *models.User) error {
	var existingUser models.User
	if result := db.Where("login = ", newUser.Login).First(&existingUser); result.Error != nil {
		log.Println("User of this login already exists")
		return result.Error
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	newUser.Password = string(hashedPassword)

	if result := db.Create(&newUser); result.Error != nil {
		return result.Error
	}

	log.Printf("User '%s' registered successfully with ID: %d\n", newUser.Username, newUser.ID)
	return nil
}