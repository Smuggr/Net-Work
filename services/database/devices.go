package database

import (
	"log"
	"network/data/errors"
	"network/data/models"
	"network/services/validation"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetDevice(db *gorm.DB, login string) (*models.Device) {
	var device models.Device
	if result := db.Where("login = ?", login).First(&device); result.Error != nil {
		return nil
	}

	return &device
}

func GetLimitedDevices(db *gorm.DB, limit int) ([]models.Device, error) {
	var devices []models.Device
	if err := db.Limit(limit).Find(&devices).Error; err != nil {
		return nil, err
	}

	return devices, nil
}

func GetPaginatedDevices(db *gorm.DB, page int, pageSize int) ([]models.Device, error) {
	var devices []models.Device
	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&devices).Error; err != nil {
		return nil, err
	}

	return devices, nil
}

func UpdateDevice(db *gorm.DB, updatedDevice *models.Device) *errors.ErrorWrapper {
	var existingDevice *models.Device = GetDevice(db, updatedDevice.Login)
	if existingDevice == nil {
		return errors.ErrDeviceNotFound
	}

	if updatedDevice.Username != "" {
		if err := validation.ValidateUsername(updatedDevice.Username); err != nil {
			return err
		}

		existingDevice.Username = updatedDevice.Username
	}

	if updatedDevice.Password != "" {
		if err := validation.ValidatePassword(updatedDevice.Password); err != nil {
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedDevice.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.ErrHashingPassword
		}

		existingDevice.Password = string(hashedPassword)
	}

	if result := db.Save(&existingDevice); result.Error != nil {
		return errors.ErrUpdatingUserInDB
	}

	log.Printf("device '%s' updated successfully", existingDevice.Login)
	return nil
}

func RegisterDevice(db *gorm.DB, newDevice *models.Device) *errors.ErrorWrapper {
	if existingDevice := GetUser(db, newDevice.Login); existingDevice != nil {
		return errors.ErrDeviceAlreadyExists
	}

	if err := validation.ValidateLogin(newDevice.Login); err != nil {
		return err
	}

	if err := validation.ValidateUsername(newDevice.Username); err != nil {
		return err
	}

	if err := validation.ValidatePassword(newDevice.Password); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newDevice.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.ErrHashingPassword
	}

	newDevice.Password = string(hashedPassword)
	if result := db.Create(&newDevice); result.Error != nil {
		return errors.ErrRegisteringDeviceInDB
	}

	log.Printf("device '%s' registered successfully", newDevice.Login)
	return nil
}