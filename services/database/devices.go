package database

import (
	"network/services/bridge"
	"network/services/provider"
	"network/utils/errors"
	"network/utils/models"
	"network/utils/validation"

	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

func GetDeviceByUsername(username string) *models.Device {
	var device models.Device
	if result := DB.Where("username = ?", username).First(&device); result.Error != nil {
		return nil
	}

	return &device
}

func GetDevice(clientID string) *models.Device {
	var device models.Device
	if result := DB.Where("client_id = ?", clientID).First(&device); result.Error != nil {
		return nil
	}

	return &device
}

func GetLimitedDevices(limit int) ([]models.Device, error) {
	var devices []models.Device
	if err := DB.Limit(limit).Find(&devices).Error; err != nil {
		return nil, err
	}

	return devices, nil
}

func GetPaginatedDevices(page int, pageSize int) ([]models.Device, error) {
	var devices []models.Device
	if err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&devices).Error; err != nil {
		return nil, err
	}

	return devices, nil
}

func InitializeDevices() error {
	devices, err := GetLimitedDevices(-1)
	if err != nil {
		return err
	}

	for _, device := range devices {
		log.Debug("creating device plugin", "client_id", device.ClientID, "plugin", device.Plugin)
		if _, err := provider.CreateDevicePlugin(device.Plugin, device.ClientID); err != nil {
			return err
		}
	}

	return nil
}

func AuthenticateDevicePassword(existingDevice *models.Device, devicePassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(existingDevice.Password), []byte(devicePassword)); err != nil {
		return err
	}

	return nil
}

func UpdateDevice(updatedDevice *models.Device) *errors.ErrorWrapper {
	var existingDevice *models.Device = GetDevice(updatedDevice.ClientID)
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

	if result := DB.Save(&existingDevice); result.Error != nil {
		return errors.ErrUpdatingDeviceInDB.Format(existingDevice.ClientID)
	}

	log.Infof("device %s updated successfully", existingDevice.ClientID)
	return nil
}

func RegisterDevice(newDevice *models.Device) *errors.ErrorWrapper {
	existingDevice := GetDevice(newDevice.ClientID)
	if existingDevice != nil {
		log.Debug(existingDevice.ClientID)
		return errors.ErrDeviceAlreadyExists.Format(newDevice.ClientID)
	}

	existingDevice = GetDeviceByUsername(newDevice.Username)
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
	if result := DB.Create(&newDevice); result.Error != nil {
		return errors.ErrRegisteringDeviceInDB.Format(newDevice.ClientID)
	}

	if _, err = provider.CreateDevicePlugin(newDevice.Plugin, newDevice.ClientID); err != nil {
		return errors.ErrCreatingDevicePlugin.Format(newDevice.ClientID, newDevice.Plugin)
	}

	log.Infof("device %s registered successfully", newDevice.ClientID)
	return nil
}

// Delete config file from disk and remove the plugin instance in loader device plugins, disconnect client from broker
func RemoveDevice(deviceToRemove *models.Device) *errors.ErrorWrapper {
	if result := DB.Delete(&deviceToRemove); result.Error != nil {
		return errors.ErrRemovingDeviceFromDB.Format(deviceToRemove.ClientID)
	}

	if err := provider.RemoveDevicePlugin(deviceToRemove.ClientID); err != nil {
		return errors.ErrRemovingDevicePlugin.Format(deviceToRemove.ClientID, deviceToRemove.Plugin)
	}

	

	log.Infof("device %s removed successfully", deviceToRemove.ClientID)
	return nil
}
