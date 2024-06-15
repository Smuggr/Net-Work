package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"smuggr.xyz/net-work/common/configurator"
	"smuggr.xyz/net-work/common/logger"
	"smuggr.xyz/net-work/core/datastorer"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var APIConfig = &configurator.Config.API

func createDeviceToken(clientID string) (string, *logger.MessageWrapper) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"client_id": clientID,
		"exp":       time.Now().Add(time.Duration(APIConfig.JWTLifespanMinutes) * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	if err != nil {
		return "", logger.ErrSigningToken
	}

	return tokenString, nil
}

func bindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.BindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": logger.ErrInvalidRequestPayload})
		return false
	}
	return true
}

func AuthenticateDeviceHandler(c *gin.Context) {
	var device datastorer.Device
	if !bindJSON(c, &device) {
		return
	}

	existingDevice := datastorer.GetDevice(device.ClientID)
	if existingDevice == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": logger.ErrResourceNotFound.Format(device.ClientID, logger.ResourceDevice)})
		return
	}

	if err := datastorer.AuthenticateDevicePassword(existingDevice, device.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": logger.ErrInvalidCredentials})
		return
	}

	tokenString, err := createDeviceToken(device.ClientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": logger.MsgResourceAuthenticateSuccess.Format(device.ClientID, logger.ResourceDevice),
		"token":   tokenString,
	})
}

func RegisterDeviceHandler(c *gin.Context) {
	var device datastorer.Device
	if !bindJSON(c, &device) {
		return
	}

	if err := datastorer.RegisterDevice(&device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	tokenString, err := createDeviceToken(device.ClientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": logger.MsgResourceRegisterSuccess.Format(device.ClientID, logger.ResourceDevice),
		"token":   tokenString,
	})
}

func UpdateDeviceHandler(c *gin.Context) {
	var device datastorer.Device
	if !bindJSON(c, &device) {
		return
	}

	if err := datastorer.UpdateDevice(&device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": logger.MsgResourceUpdateSuccess.Format(device.ClientID, logger.ResourceDevice),
	})
}

func RemoveDeviceHandler(c *gin.Context) {
	clientID := c.Query("client_id")

	if clientID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": logger.ErrInvalidRequestPayload})
		return
	}

	device := datastorer.GetDevice(clientID)
	if device == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": logger.ErrResourceNotFound.Format(clientID, logger.ResourceDevice),
		})

		return
	}

	if err := datastorer.RemoveDevice(device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": logger.ErrRemovingResourceFromDB.Format(clientID, logger.ResourceDevice),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": logger.MsgResourceRemoveSuccess.Format(device.ClientID, logger.ResourceDevice),
	})
}

func GetDeviceHandler(c *gin.Context) {
	clientID := c.Param("client_id")

	device := datastorer.GetDevice(clientID)
	if device == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": logger.ErrResourceNotFound.Format(clientID, logger.ResourceDevice),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": logger.MsgResourceFetchSuccess.Format(clientID, logger.ResourceDevice),
		"device":  device,
	})
}

func GetAllDevicesHandler(c *gin.Context) {
	limit := -1
	devices := datastorer.GetLimitedDevices(limit)
	if devices == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": logger.ErrFetchingResourceFromDB.Format(limit, logger.ResourceDevice),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": logger.MsgResourceFetchSuccess.Format(len(devices), logger.ResourceDevice),
		"devices": devices,
	})
}

func GetLimitedDevicesHandler(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": logger.ErrInvalidRequestPayload,
		})

		return
	}

	devices := datastorer.GetLimitedDevices(limit)
	if devices == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": logger.ErrFetchingResourceFromDB.Format(limit, logger.ResourceDevice),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": logger.MsgResourceFetchSuccess.Format(len(devices), logger.ResourceDevice),
		"limit":   limit,
		"devices": devices,
	})
}

func GetPaginatedDevicesHandler(c *gin.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": logger.ErrInvalidRequestPayload,
		})

		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": logger.ErrInvalidRequestPayload,
		})

		return
	}

	devices := datastorer.GetPaginatedDevices(page, pageSize)
	if devices == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": logger.ErrFetchingResourceFromDB.Format(fmt.Sprintf("%d / %d", page, pageSize), logger.ResourceDevice),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  logger.MsgResourceFetchSuccess.Format(len(devices), logger.ResourceDevice),
		"page":     page,
		"pageSize": pageSize,
		"devices":  devices,
	})
}
