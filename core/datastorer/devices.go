package datastorer

import (
	"fmt"
	"smuggr/net-work/common/logger"
)

func GetDevice(clientID string) (*Device, *logger.MessageWrapper) {
	var device Device
	if result := DB.Where("client_id = ?", clientID).First(&device); result.Error != nil {
		return nil, logger.ErrFetchingResourceFromDB.Format(clientID, logger.ResourceDevice)
	}

	return &device, nil
}

func GetLimitedDevices(limit int) ([]*Device, *logger.MessageWrapper) {
	var devices []*Device
	if err := DB.Limit(limit).Find(&devices).Error; err != nil {
		Logger.Debug("failed to get limited devices", "error", err)
		return nil, logger.ErrFetchingResourceFromDB.Format(fmt.Sprintf("limit %d", limit), logger.ResourceDevice)
	}

	return devices, nil
}

func GetPaginatedDevices(page int, pageSize int) ([]*Device, *logger.MessageWrapper) {
	var devices []*Device
	if err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&devices).Error; err != nil {
		return nil, logger.ErrFetchingResourceFromDB.Format(fmt.Sprintf("page %d, pageSize %d", page, pageSize), logger.ResourceDevice)
	}

	return devices, nil
}
