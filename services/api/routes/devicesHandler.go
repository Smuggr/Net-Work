package routes

import (
	"net/http"
	"strconv"

	"network/data/errors"
	"network/data/messages"
	"network/data/models"
	"network/services/database"

	"github.com/gin-gonic/gin"
)


func RegisterDeviceHandler(c *gin.Context) {
	var device models.Device
	if err := c.BindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	if err := database.RegisterDevice(database.DB, &device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	tokenString, err := createToken(device.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgDeviceRegisterSuccess.Format(device.ClientID), "token": tokenString})
}

func UpdateDeviceHandler(c *gin.Context) {
	var device models.Device
	if err := c.BindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	if err := database.UpdateDevice(database.DB, &device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgDeviceUpdateSuccess.Format(device.ClientID)})
}

func RemoveDeviceHandler(c *gin.Context) {
	clientID := c.Query("client_id")

	if clientID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	device := database.GetDevice(database.DB, clientID)
    if device == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrDeviceNotFound.Format(clientID)})
        return
    }

	if err := database.RemoveDevice(database.DB, device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrRemovingDeviceFromDB.Format(clientID)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgDeviceRemoveSuccess.Format(device.ClientID)})
}


func GetDeviceHandler(c *gin.Context) { 
	clientID := c.Param("client_id")

	device := database.GetDevice(database.DB, clientID)
	if device == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrDeviceNotFound.Format(clientID)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": messages.MsgDeviceFetchSuccess.Format(clientID), 
		"device":  device,
	})
}

func GetAllDevicesHandler(c *gin.Context) {
	devices, err := database.GetLimitedDevices(database.DB, -1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFetchingDevicesFromDB})
		return
	}

    c.JSON(http.StatusOK, gin.H{
		"message": messages.MsgDevicesFetchSuccess.Format(len(devices)), 
		"devices": devices,
	})
}

func GetLimitedDevicesHandler(c *gin.Context) {
	limitStr := c.Query("limit")
    limit, err := strconv.Atoi(limitStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
    }

    devices, err := database.GetLimitedDevices(database.DB, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFetchingDevicesFromDB})
		return
	}

    c.JSON(http.StatusOK, gin.H{
		"message": messages.MsgDevicesFetchSuccess.Format(len(devices)),
		"limit":   limit,
		"devices": devices,
	})
}

func GetPaginatedDevicesHandler(c *gin.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	devices, err := database.GetPaginatedDevices(database.DB, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFetchingDevicesFromDB})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  messages.MsgDevicesFetchSuccess.Format(len(devices)),
		"page":     page,
		"pageSize": pageSize,
		"devices":  devices,
	})
}
