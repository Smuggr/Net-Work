package database

import (
	"os"
	"log"

	"overseer/services/models"

	"gorm.io/gorm"
)


func getDefaultUser(db *gorm.DB) (models.User, error) {
    var defaultUser models.User
    result := db.First(&defaultUser, "id = ?", 1)

    return defaultUser, result.Error
}

func createDefaultUser(db *gorm.DB) {
	defaultUser, err := getDefaultUser(db)
	if err == nil {
		log.Println("Default user already exists.")

		adminLogin := os.Getenv("ADMIN_LOGIN")

		result := db.Where("login = ? AND id != ?", adminLogin, 1).First(&defaultUser)
		if result.Error == nil {
			log.Fatalf("Error updating default user, user of this login already exists: %v", result.Error)
		}

		defaultUser.Login = adminLogin
		defaultUser.Username = os.Getenv("ADMIN_USERNAME")
		defaultUser.Password = os.Getenv("ADMIN_PASSWORD")

		result = db.Save(&defaultUser)
		if result.Error == nil {
			log.Println("Default user updated successfully")
		} else {
			log.Fatalf("Error updating default user: %v", result.Error)
		}

		return
	} else {
		log.Fatalf("Error checking default user existence: %v", err)
	}

	defaultUser = models.User{
		Login:           os.Getenv("ADMIN_LOGIN"),
		Username:        os.Getenv("ADMIN_USERNAME"),
		Password:        os.Getenv("ADMIN_PASSWORD"),
		PermissionLevel: 1,
	}

	result := db.Create(&defaultUser)
	if result.Error == nil {
		log.Println("Default user created successfully.")
	} else {
		log.Fatalf("Error creating default user: %v", result.Error)
	}
}

func GetUserByID(db *gorm.DB, userID uint) (models.User, error) {
	var user models.User

	result := db.First(&user, userID)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func GetUserByLogin(db *gorm.DB, login string) (models.User, error) {
	var user models.User

	result := db.Where("login = ?", login).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func GetUserByUsername(db *gorm.DB, username string) (models.User, error) {
	var user models.User

	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func RegisterUser(db *gorm.DB, newUser models.User) error {
	if _, err := GetUserByLogin(db, newUser.Login); err != nil {
		return err
	}

	if _, err := GetUserByLogin(db, os.Getenv("ADMIN_LOGIN")); err != nil {
		return err
	}

	if result := db.Create(&newUser); result.Error != nil {
		return result.Error
	}

	log.Printf("User '%s' registered successfully with ID: %d\n", newUser.Username, newUser.ID)
	return nil
}