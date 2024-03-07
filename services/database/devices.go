package database

import (
	"network/utils/validation"
	"network/utils/errors"
	"network/utils/models"

	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetDeviceByUsername(db *gorm.DB, username string) *models.Device {
	var device models.Device
	if result := db.Where("username = ?", username).First(&device); result.Error != nil {
		return nil
	}

	return &device
}

func GetDevice(db *gorm.DB, clientID string) *models.Device {
	var device models.Device
	if result := db.Where("client_id = ?", clientID).First(&device); result.Error != nil {
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

func AuthenticateDevicePassword(existingDevice *models.Device, devicePassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(existingDevice.Password), []byte(devicePassword)); err != nil {
		return err
	}

	return nil
}

func UpdateDevice(db *gorm.DB, updatedDevice *models.Device) *errors.ErrorWrapper {
	var existingDevice *models.Device = GetDevice(db, updatedDevice.ClientID)
	if existingDevice == nil {
		return errors.ErrDeviceNotFound.Format(updatedDevice.Username)
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
		return errors.ErrUpdatingDeviceInDB.Format(existingDevice.ClientID)
	}

	log.Infof("device %s updated successfully", existingDevice.ClientID)
	return nil
}

func RegisterDevice(db *gorm.DB, newDevice *models.Device) *errors.ErrorWrapper {
	existingDevice := GetDevice(db, newDevice.ClientID)
	if existingDevice != nil {
		log.Debug(existingDevice.ClientID)
		return errors.ErrDeviceAlreadyExists.Format(newDevice.ClientID)
	}

	existingDevice = GetDeviceByUsername(db, newDevice.Username)
	if existingDevice != nil {
		log.Debug(existingDevice.Username)
		return errors.ErrDeviceAlreadyExists.Format(newDevice.Username)
	}

	if err := validation.ValidateClientID(newDevice.ClientID); err != nil {
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
		return errors.ErrRegisteringDeviceInDB.Format(newDevice.ClientID)
	}

	log.Infof("device %s registered successfully", newDevice.ClientID)
	return nil
}

func RemoveDevice(db *gorm.DB, deviceToRemove *models.Device) *errors.ErrorWrapper {
	if result := db.Delete(&deviceToRemove); result.Error != nil {
		return errors.ErrRemovingDeviceFromDB.Format(deviceToRemove.ClientID)
	}

	log.Infof("device %s removed successfully", deviceToRemove.ClientID)
	return nil
}
