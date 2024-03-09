package database

import (
	"network/utils/errors"
	"network/utils/models"
	"network/utils/validation"

	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultAdminLogin           string = "administrator"
	DefaultAdminPassword        string = "Password123$"
	DefaultAdminUsername        string = "Administrator"
	DefaultAdminPermissionLevel int    = -1
)

func GetUser(login string) *models.User {
	var user models.User
	if result := DB.Where("login = ?", login).First(&user); result.Error != nil {
		return nil
	}

	return &user
}

func GetLimitedUsers(limit int) ([]models.User, error) {
	var users []models.User
	if err := DB.Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetPaginatedUsers(page int, pageSize int) ([]models.User, error) {
	var users []models.User
	if err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func AuthenticateUserPassword(existingUser *models.User, userPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userPassword)); err != nil {
		return err
	}

	return nil
}

func UpdateUser(updatedUser *models.User) *errors.ErrorWrapper {
	var existingUser *models.User = GetUser(updatedUser.Login)
	if existingUser == nil {
		return errors.ErrUserNotFound.Format(updatedUser.Login)
	}

	if updatedUser.PermissionLevel < 0 {
		return errors.ErrOperationNotPermitted
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

	if result := DB.Save(&existingUser); result.Error != nil {
		return errors.ErrUpdatingUserInDB.Format(existingUser.Login)
	}

	log.Infof("user %s updated successfully", existingUser.Login)
	return nil
}

func RegisterUser(newUser *models.User) *errors.ErrorWrapper {
	if existingUser := GetUser(newUser.Login); existingUser != nil {
		return errors.ErrUserAlreadyExists.Format(newUser.Login)
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
	if result := DB.Create(&newUser); result.Error != nil {
		return errors.ErrRegisteringUserInDB.Format(newUser.Login)
	}

	log.Infof("user %s registered successfully", newUser.Login)
	return nil
}

func RemoveUser(userToRemove *models.User) *errors.ErrorWrapper {
	if userToRemove.PermissionLevel < 0 {
		return errors.ErrOperationNotPermitted
	}

	if result := DB.Delete(&userToRemove); result.Error != nil {
		return errors.ErrRemovingUserFromDB.Format(userToRemove.Login)
	}

	log.Infof("user %s removed successfully", userToRemove.Login)
	return nil
}

func RegisterDefaultAdmin() error {
	userModel := models.User{
		Login:           DefaultAdminLogin,
		Username:        DefaultAdminUsername,
		Password:        DefaultAdminPassword,
		PermissionLevel: DefaultAdminPermissionLevel,
	}

	if existingUser := GetUser(DefaultAdminLogin); existingUser != nil {
		UpdateUser(&userModel)
		return nil
	}

	if err := RegisterUser(&userModel); err != nil {
		return err
	}

	return nil
}
