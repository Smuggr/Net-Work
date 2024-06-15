package routes

import (
	"smuggr.xyz/net-work/api/v1/handlers"
	// "smuggr.xyz/net-work/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

func SetupDevicesRoutes(rootGroup *gin.RouterGroup) {
	deviceGroup := rootGroup.Group("/device")
	// deviceGroup.Use(middleware.UserAuthentication())
	{
		deviceGroup.POST("/authenticate", handlers.AuthenticateDeviceHandler)
		deviceGroup.GET("/:client_id", handlers.GetDeviceHandler)
		deviceGroup.POST("/register", handlers.RegisterDeviceHandler)
		deviceGroup.PUT("/update", handlers.UpdateDeviceHandler)
		deviceGroup.DELETE("/remove", handlers.RemoveDeviceHandler)
	}

	devicesGroup := rootGroup.Group("/devices")
	// devicesGroup.Use(middleware.UserAuthentication())
	{
		devicesGroup.GET("/all", handlers.GetAllDevicesHandler)
		devicesGroup.GET("/limited", handlers.GetLimitedDevicesHandler)
		devicesGroup.GET("/paginated", handlers.GetPaginatedDevicesHandler)
	}

	// devicesInteractionsGroup := rootGroup.Group("/devices/interactions/:client_id")
	// devicesInteractionsGroup.Use(DeviceAuthenticationMiddleware())
	// devicesInteractionsGroup.Use(bridger.RouteEnabledMiddleware())
	// {
	// 	devicesInteractionsGroup.GET("/*directory", bridger.InteractionsHandler)
	// 	devicesInteractionsGroup.POST("/*directory", bridger.InteractionsHandler)
	// }
}
