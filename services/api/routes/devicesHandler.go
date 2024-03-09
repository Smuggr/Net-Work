package routes

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"network/services/database"
	"network/utils/configuration"
	"network/utils/errors"
	"network/utils/messages"
	"network/utils/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func createDeviceToken(client_id string) (string, *errors.ErrorWrapper) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"client_id": client_id,
		"exp":       time.Now().Add(time.Duration(configuration.Config.API.JWTLifespanMinutes) * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	if err != nil {
		return "", errors.ErrSigningToken
	}

	return tokenString, nil
}

func AuthenticateDeviceHandler(c *gin.Context) {
	var device models.Device
	if err := c.BindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	existingDevice := database.GetDevice(device.ClientID)
	if existingDevice == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrDeviceNotFound.Format(device.ClientID)})
		return
	}

	if err := database.AuthenticateDevicePassword(existingDevice, device.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrInvalidCredentials})
		return
	}

	tokenString, err := createDeviceToken(device.ClientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": messages.MsgDeviceAuthenticateSuccess.Format(device.ClientID),
		"token":   tokenString,
	})
}

func RegisterDeviceHandler(c *gin.Context) {
	var device models.Device
	if err := c.BindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	if err := database.RegisterDevice(&device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	tokenString, err := createDeviceToken(device.Username)
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

	if err := database.UpdateDevice(&device); err != nil {
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

	device := database.GetDevice(clientID)
	if device == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrDeviceNotFound.Format(clientID)})
		return
	}

	if err := database.RemoveDevice(device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrRemovingDeviceFromDB.Format(clientID)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgDeviceRemoveSuccess.Format(device.ClientID)})
}

func GetDeviceHandler(c *gin.Context) {
	clientID := c.Param("client_id")

	device := database.GetDevice(clientID)
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
	devices, err := database.GetLimitedDevices(-1)
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

	devices, err := database.GetLimitedDevices(limit)
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

	devices, err := database.GetPaginatedDevices(page, pageSize)
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
