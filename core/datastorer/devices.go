package datastorer

import (
	"smuggr.xyz/net-work/common/logger"
	"smuggr.xyz/net-work/common/validator"

	"golang.org/x/crypto/bcrypt"
)

func InitializeDevice(device *Device) *logger.MessageWrapper {
	// if _, err := bridger.GetClient(device.ClientID); err != nil {
	// 	log.Warn("client not found, probably offline", "client_id", device.ClientID, "error", err)
	// 	return err
	// }

	// if _, err := provider.CreateDevicePlugin(device.Plugin, device.ClientID); err != nil {
	// 	return err
	// }

	return nil
}

func InitializeDevices() *logger.MessageWrapper {
	devices := GetLimitedDevices(-1)
	if devices == nil {
		return logger.ErrInitializingResource.Format(logger.ResourceDevice)
	}

	for _, device := range devices {
		if err := InitializeDevice(&device); err != nil {
			continue
		}
	}

	return nil
}

func AuthenticateDevicePassword(existingDevice *Device, devicePassword string) *logger.MessageWrapper {
	if err := bcrypt.CompareHashAndPassword([]byte(existingDevice.Password), []byte(devicePassword)); err != nil {
		return logger.ErrAuthenticatingResource.Format(devicePassword, logger.ResourceDevice)
	}

	return nil
}

func GetDevice(clientID string) *Device {
	var device Device
	if result := DB.Where("client_id = ?", clientID).First(&device); result.Error != nil {
		return nil
	}

	return &device
}

func GetLimitedDevices(limit int) []Device {
	var devices []Device
	if err := DB.Limit(limit).Find(&devices).Error; err != nil {
		return nil
	}

	return devices
}

func GetPaginatedDevices(page int, pageSize int) []Device {
	var devices []Device
	if err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&devices).Error; err != nil {
		return nil
	}

	return devices
}

func UpdateDevice(updatedDevice *Device) *logger.MessageWrapper {
	var existingDevice *Device = GetDevice(updatedDevice.ClientID)
	if existingDevice == nil {
		return logger.ErrResourceNotFound.Format(updatedDevice.ClientID, logger.ResourceDevice)
	}

	if updatedDevice.DisplayName != "" {
		if err := validator.ValidateDisplayName(updatedDevice.DisplayName); err != nil {
			return err
		}

		existingDevice.DisplayName = updatedDevice.DisplayName
	}

	if updatedDevice.Password != "" {
		if err := validator.ValidatePassword(updatedDevice.Password); err != nil {
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedDevice.Password), bcrypt.DefaultCost)
		if err != nil {
			return logger.ErrHashingPassword
		}

		existingDevice.Password = string(hashedPassword)
	}

	if result := DB.Save(&existingDevice); result.Error != nil {
		return logger.ErrUpdatingResourceInDB.Format(existingDevice.ClientID, logger.ResourceDevice)
	}

	return nil
}

func RegisterDevice(newDevice *Device) *logger.MessageWrapper {
	existingDevice := GetDevice(newDevice.ClientID)
	if existingDevice != nil {
		return logger.ErrResourceAlreadyExists.Format(newDevice.ClientID, logger.ResourceDevice)
	}

	existingDevice = GetDevice(newDevice.ClientID)
	if existingDevice != nil {
		return logger.ErrResourceAlreadyExists.Format(newDevice.ClientID, logger.ResourceDevice)
	}

	if err := validator.ValidateClientID(newDevice.ClientID); err != nil {
		return err
	}

	if err := validator.ValidateDisplayName(newDevice.DisplayName); err != nil {
		return err
	}

	if err := validator.ValidatePassword(newDevice.Password); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newDevice.Password), bcrypt.DefaultCost)
	if err != nil {
		return logger.ErrHashingPassword
	}

	newDevice.Password = string(hashedPassword)
	if result := DB.Create(&newDevice); result.Error != nil {
		return logger.ErrRegisteringResourceInDB.Format(newDevice.ClientID, logger.ResourceDevice)
	}

	// if err = InitializeDevice(newDevice); err != nil {
	// 	log.Warn("failed to initialize device", "client", newDevice.ClientID, "error", err)
	// }

	return nil
}

// Delete config file from disk and remove the plugin instance in loader device plugins, disconnect client from broker
func RemoveDevice(deviceToRemove *Device) *logger.MessageWrapper {
	// if err := bridger.DisconnectClient(deviceToRemove.ClientID); err != nil {
	// 	log.Warn("failed to disconnect client", "client_id", deviceToRemove.ClientID, "error", err)
	// }

	// if err := provider.RemoveDevicePlugin(deviceToRemove.ClientID); err != nil {
	// 	log.Error("failed to remove device plugin", "client_id", deviceToRemove.ClientID, "error", err)
	// }

	// if result := DB.Unscoped().Delete(&deviceToRemove); result.Error != nil {
	// 	log.Error("failed to remove device from db", "client_id", deviceToRemove.ClientID, "error", result.Error)
	// 	return errors.ErrRemovingDeviceFromDB.Format(deviceToRemove.ClientID)
	// }

	// log.Infof("device %s removed successfully", deviceToRemove.ClientID)
	return nil
}
