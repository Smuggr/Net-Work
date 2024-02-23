package database

import (
	"log"

	"network/data/errors"
	"network/data/models"
	"network/services/validation"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

const (
	DefaultAdminLogin string = "administrator"
	DefaultAdminPassword string = "Password123$"
	DefaultAdminUsername string = "Administrator"
)

func UpdateUser(db *gorm.DB, updatedUser *models.User) *errors.ErrorWrapper {
	var existingUser models.User
	if result := db.Where("login = ?", updatedUser.Login).First(&existingUser); result.Error != nil {
		return errors.ErrUserNotFound
	}

	if updatedUser.Username != "" {
		if err := validation.ValidateUsername(updatedUser.Username); err != nil {
			return err
		}

		existingUser.Username = updatedUser.Username
	}

	if updatedUser.Password != "" {
		if err := validation.ValidatePassword(updatedUser.Password); err != nil {
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.ErrHashingPassword
		}

		existingUser.Password = string(hashedPassword)
	}

	if result := db.Save(&existingUser); result.Error != nil {
		return errors.ErrUpdatingUserInDB
	}

	log.Printf("user '%s' updated successfully", existingUser.Login)
	return nil
}

func RegisterUser(db *gorm.DB, newUser *models.User) *errors.ErrorWrapper {
	var existingUser models.User
	if result := db.Where("login = ?", newUser.Login).First(&existingUser); result.Error == nil {
		return errors.ErrUserAlreadyExists
	}

	if err := validation.ValidateLogin(newUser.Login); err != nil {
		return err
	}

	if err := validation.ValidateUsername(newUser.Username); err != nil {
		return err
	}

	if err := validation.ValidatePassword(newUser.Password); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.ErrHashingPassword
	}

	newUser.Password = string(hashedPassword)
	if result := db.Create(&newUser); result.Error != nil {
		return errors.ErrRegisteringUserInDB
	}

	log.Printf("user '%s' registered successfully", newUser.Login)
	return nil
}