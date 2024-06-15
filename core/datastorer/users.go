package datastorer

import (
	"fmt"

	"smuggr.xyz/net-work/common/logger"
	"smuggr.xyz/net-work/common/validator"

	"golang.org/x/crypto/bcrypt"
)

func GetUser(login string) (*User, *logger.MessageWrapper) {
	var user User
	if result := DB.Where("login = ?", login).First(&user); result.Error != nil {
		return nil, logger.ErrFetchingResourceFromDB.Format(login, logger.ResourceUser)
	}

	return &user, nil
}

func GetLimitedUsers(limit int) ([]*User, *logger.MessageWrapper) {
	var users []*User
	if err := DB.Limit(limit).Find(&users).Error; err != nil {
		return nil, logger.ErrFetchingResourceFromDB.Format(fmt.Sprintf("limit %d", limit), logger.ResourceUser)
	}

	return users, nil
}

func GetPaginatedUsers(page int, pageSize int) ([]*User, *logger.MessageWrapper) {
	var users []*User
	if err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, logger.ErrFetchingResourceFromDB.Format(fmt.Sprintf("page %d, pageSize %d", page, pageSize), logger.ResourceUser)
	}

	return users, nil
}

func AuthenticateUserPassword(existingUser *User, userPassword string) *logger.MessageWrapper {
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userPassword)); err != nil {
		return logger.ErrAuthenticatingResource.Format(userPassword, logger.ResourceUser)
	}

	return logger.MsgResourceAuthenticateSuccess.Format(existingUser.Login, logger.ResourceUser)
}

func UpdateUser(updatedUser *User) *logger.MessageWrapper {
	existingUser, _ := GetUser(updatedUser.Login)
	if existingUser == nil {
		return logger.ErrResourceNotFound.Format(updatedUser.Login, logger.ResourceUser)
	}

	if updatedUser.PermissionLevel < 0 {
		return logger.ErrOperationNotPermitted
	}

	if updatedUser.Username != "" {
		if err := validator.ValidateUsername(updatedUser.Username); err != nil {
			return err
		}

		existingUser.Username = updatedUser.Username
	}

	if updatedUser.Password != "" {
		if err := validator.ValidatePassword(updatedUser.Password); err != nil {
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return logger.ErrHashingPassword
		}

		existingUser.Password = string(hashedPassword)
	}

	if result := DB.Save(&existingUser); result.Error != nil {
		return logger.ErrUpdatingResourceInDB.Format(existingUser.Login, logger.ResourceUser)
	}

	return logger.MsgResourceUpdateSuccess.Format(existingUser.Login, logger.ResourceUser)
}

func RegisterUser(newUser *User) *logger.MessageWrapper {
	if existingUser, _ := GetUser(newUser.Login); existingUser != nil {
		return logger.ErrResourceAlreadyExists.Format(newUser.Login, logger.ResourceUser)
	}

	if err := validator.ValidateLogin(newUser.Login); err != nil {
		return err
	}

	if err := validator.ValidateUsername(newUser.Username); err != nil {
		return err
	}

	if err := validator.ValidatePassword(newUser.Password); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return logger.ErrHashingPassword
	}

	newUser.Password = string(hashedPassword)
	if result := DB.Create(&newUser); result.Error != nil {
		return logger.ErrRegisteringResourceInDB.Format(newUser.Login, logger.ResourceUser)
	}

	return logger.MsgResourceRegisterSuccess.Format(newUser.Login, logger.ResourceUser)
}

func RemoveUser(userToRemove *User) *logger.MessageWrapper {
	if userToRemove.PermissionLevel < 0 {
		return logger.ErrOperationNotPermitted
	}

	if result := DB.Unscoped().Delete(&userToRemove); result.Error != nil {
		return logger.ErrRemovingResourceFromDB.Format(userToRemove.Login, logger.ResourceUser)
	}

	return logger.MsgResourceRemoveSuccess.Format(userToRemove.Login, logger.ResourceUser)
}

func RegisterDefaultAdmin() *logger.MessageWrapper {
	userModel := User{
		Login:           DefaultAdminLogin,
		Username:        DefaultAdminUsername,
		Password:        DefaultAdminPassword,
		PermissionLevel: DefaultAdminPermissionLevel,
	}

	if existingUser, _ := GetUser(DefaultAdminLogin); existingUser != nil {
		UpdateUser(&userModel)
		return logger.ErrResourceAlreadyExists.Format(DefaultAdminLogin, logger.ResourceUser)
	}

	if err := RegisterUser(&userModel); err != nil {
		return err
	}

	return logger.MsgResourceRegisterSuccess.Format(DefaultAdminLogin, logger.ResourceUser)
}
