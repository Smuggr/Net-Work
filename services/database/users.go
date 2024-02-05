package database

import (
	"log"

	"overseer/data/errors"
	"overseer/data/models"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

const (
	DefaultAdminLogin string = "admin"
	DefaultAdminPassword string = "admin"
	DefaultAdminUsername string = "admin"
)

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

func UpdateUser(db *gorm.DB, updatedUser *models.User) *errors.ErrorWrapper {
	var existingUser models.User
	if result := db.Where("login = ?", updatedUser.Login).First(&existingUser); result.Error != nil {
		return errors.ErrUserNotFound
	}

	// if updatedUser.ID != existingUser.ID {
	// 	return gorm.ErrRecordNotFound
	// }

	existingUser.Username = updatedUser.Username

	if updatedUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.ErrHashingPassword
		}

		existingUser.Password = string(hashedPassword)
	}

	if result := db.Save(&existingUser); result.Error != nil {
		return errors.ErrUpdatingUserInDB
	}

	log.Printf("User '%s' updated successfully", existingUser.Login)
	return nil
}

func RegisterUser(db *gorm.DB, newUser *models.User) *errors.ErrorWrapper {
	var existingUser models.User
	if result := db.Where("login = ?", newUser.Login).First(&existingUser); result.Error == nil {
		return errors.ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.ErrHashingPassword
	}

	newUser.Password = string(hashedPassword)
	if result := db.Create(&newUser); result.Error != nil {
		return errors.ErrRegisteringUserInDB
	}

	log.Printf("User '%s' registered successfully", newUser.Login)
	return nil
}