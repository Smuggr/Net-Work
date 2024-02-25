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

func GetUser(db *gorm.DB, login string) (*models.User) {
	var user models.User
	if result := db.Where("login = ?", login).First(&user); result.Error != nil {
		return nil
	}

	return &user
}

func GetLimitedUsers(db *gorm.DB, limit int) ([]models.User, error) {
	var users []models.User
	if err := db.Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetPaginatedUsers(db *gorm.DB, page int, pageSize int) ([]models.User, error) {
	var users []models.User
	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(db *gorm.DB, updatedUser *models.User) *errors.ErrorWrapper {
	var existingUser *models.User = GetUser(db, updatedUser.Login)
	if existingUser == nil {
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
	if existingUser := GetUser(db, newUser.Login); existingUser != nil {
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

func RegisterDefaultAdmin(db *gorm.DB) error {
	userModel := models.User{
		Login:           DefaultAdminLogin,
		Username:        DefaultAdminUsername,
		Password:        DefaultAdminPassword,
		PermissionLevel: 1,
	}

	if existingUser := GetUser(db, DefaultAdminLogin); existingUser != nil {
		UpdateUser(db, &userModel)
		return nil
	}

	if err := RegisterUser(db, &userModel); err != nil {
		return err
	}

	return nil
}