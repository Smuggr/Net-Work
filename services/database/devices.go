package database

import (
	"network/common/bridger"
	"network/common/provider"
	"network/utils/errors"
	"network/utils/models"
	"network/utils/validation"

	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

func GetDeviceByUsername(username string) *models.Device {
	var device models.Device
	if result := DB.Where("username = ?", username).First(&device); result.Error != nil {
		log.Debug("failed to get device", "username", username, "error", result.Error)
		return nil
	}

	return &device
}

func GetDevice(clientID string) *models.Device {
	var device models.Device
	if result := DB.Where("client_id = ?", clientID).First(&device); result.Error != nil {
		log.Debug("failed to get device", "client_id", clientID, "error", result.Error)
		return nil
	}

	return &device
}

func GetLimitedDevices(limit int) ([]models.Device, error) {
	var devices []models.Device
	if err := DB.Limit(limit).Find(&devices).Error; err != nil {
		log.Debug("failed to get limited devices", "error", err)
		return nil, err
	}

	return devices, nil
}

func GetPaginatedDevices(page int, pageSize int) ([]models.Device, error) {
	var devices []models.Device
	if err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&devices).Error; err != nil {
		log.Debug("failed to get paginated devices", "error", err)
		return nil, err
	}

	return devices, nil
}

func InitializeDevice(device *models.Device) error {
	log.Debug("creating device plugin", "client_id", device.ClientID, "plugin", device.Plugin)

	if _, err := bridger.GetClient(device.ClientID); err != nil {
		log.Warn("client not found, probably offline", "client_id", device.ClientID, "error", err)
		return err
	}

	if _, err := provider.CreateDevicePlugin(device.Plugin, device.ClientID); err != nil {
		return err
	}

	return nil
}

func InitializeDevices() error {
	devices, err := GetLimitedDevices(-1)
	if err != nil {
		return err
	}

	for _, device := range devices {
		if err := InitializeDevice(&device); err != nil {
			continue
		}
	}

	return nil
}

func AuthenticateDevicePassword(existingDevice *models.Device, devicePassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(existingDevice.Password), []byte(devicePassword)); err != nil {
		log.Debug("failed to authenticate device password", "client_id", existingDevice.ClientID, "error", err)
		return err
	}

	return nil
}

func UpdateDevice(updatedDevice *models.Device) *errors.ErrorWrapper {
	var existingDevice *models.Device = GetDevice(updatedDevice.ClientID)
	if existingDevice == nil {
		log.Debug("device not found", "client_id", updatedDevice.ClientID)
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
			log.Debug("failed to hash password", "error", err)
			return errors.ErrHashingPassword
		}

		existingDevice.Password = string(hashedPassword)
	}

	if result := DB.Save(&existingDevice); result.Error != nil {
		log.Debug("failed to update device in db", "client_id", existingDevice.ClientID, "error", result.Error)
		return errors.ErrUpdatingDeviceInDB.Format(existingDevice.ClientID)
	}

	log.Infof("device %s updated successfully", existingDevice.ClientID)
	return nil
}

func RegisterDevice(newDevice *models.Device) *errors.ErrorWrapper {
	existingDevice := GetDevice(newDevice.ClientID)
	if existingDevice != nil {
		log.Debug("device already exists", "client_id", newDevice.ClientID, "username", existingDevice.Username)
		return errors.ErrDeviceAlreadyExists.Format(newDevice.ClientID)
	}

	existingDevice = GetDeviceByUsername(newDevice.Username)
	if existingDevice != nil {
		log.Debug("device already exists", "username", newDevice.Username, "client_id", existingDevice.ClientID)
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
		log.Debug("failed to hash password", "error", err)
		return errors.ErrHashingPassword
	}

	newDevice.Password = string(hashedPassword)
	if result := DB.Create(&newDevice); result.Error != nil {
		log.Debug("failed to register device in db", "client_id", newDevice.ClientID, "error", result.Error)
		return errors.ErrRegisteringDeviceInDB.Format(newDevice.ClientID)
	}

	if err = InitializeDevice(newDevice); err != nil {
		log.Debug("failed to initialize device", "client", newDevice.ClientID, "error", err)
		return errors.ErrCreatingDevicePlugin.Format(newDevice.ClientID, newDevice.Plugin)
	}

	log.Infof("device %s registered successfully", newDevice.ClientID)
	return nil
}

// Delete config file from disk and remove the plugin instance in loader device plugins, disconnect client from broker
func RemoveDevice(deviceToRemove *models.Device) *errors.ErrorWrapper {
	if err := bridger.DisconnectClient(deviceToRemove.ClientID); err != nil {
		log.Debug("failed to disconnect client", "client_id", deviceToRemove.ClientID, "error", err)
		return errors.ErrClientNotFound.Format(deviceToRemove.ClientID)
	}

	if err := provider.RemoveDevicePlugin(deviceToRemove.ClientID); err != nil {
		log.Debug("failed to remove device plugin", "client_id", deviceToRemove.ClientID, "error", err)
		return errors.ErrRemovingDevicePlugin.Format(deviceToRemove.ClientID, deviceToRemove.Plugin)
	}

	if result := DB.Unscoped().Delete(&deviceToRemove); result.Error != nil {
		log.Debug("failed to remove device from db", "client_id", deviceToRemove.ClientID, "error", result.Error)
		return errors.ErrRemovingDeviceFromDB.Format(deviceToRemove.ClientID)
	}

	log.Infof("device %s removed successfully", deviceToRemove.ClientID)
	return nil
}
