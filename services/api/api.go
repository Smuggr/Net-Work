package api

import (
	"os"
	"strconv"

	"network/common/bridger"
	"network/services/api/routes"
	"network/utils/configuration"

	"github.com/charmbracelet/log"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Config = &configuration.Config.API
var DevicesInteractionsGroup *gin.RouterGroup

func Initialize() chan error {
	log.Info("initializing api/v1")

	Config = &configuration.Config.API

	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()
	l := tollbooth.NewLimiter(1, nil)

	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			log.Debug(origin)
			return true
		},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	}))

	r.SetTrustedProxies([]string{})

	InitializeSwagger(r)

	apiV1Group := r.Group("/api/v1")
	apiV1Group.Use(tollbooth_gin.LimitHandler(l))
	{
		userGroup := apiV1Group.Group("/user")
		userGroup.Use(UserAuthenticationMiddleware())
		{
			userGroup.GET("/:login", routes.GetUserHandler)
			userGroup.POST("/register", routes.RegisterUserHandler)
			userGroup.PUT("/update", routes.UpdateUserHandler)
			userGroup.DELETE("/remove", routes.RemoveUserHandler)
		}

		noAuthUserGroup := apiV1Group.Group("/user")
		{
			noAuthUserGroup.POST("/authenticate", routes.AuthenticateUserHandler)
			noAuthUserGroup.POST("/validateToken", routes.ValidateUserTokenHandler)
		}

		usersGroup := apiV1Group.Group("/users")
		usersGroup.Use(UserAuthenticationMiddleware())
		{
			usersGroup.GET("/all", routes.GetAllUsersHandler)
			usersGroup.GET("/limited", routes.GetLimitedUsersHandler)
			usersGroup.GET("/paginated", routes.GetPaginatedUsersHandler)
		}

		deviceGroup := apiV1Group.Group("/device")
		deviceGroup.Use(UserAuthenticationMiddleware())
		{
			deviceGroup.POST("/authenticate", routes.AuthenticateDeviceHandler)
			deviceGroup.GET("/:client_id", routes.GetDeviceHandler)
			deviceGroup.POST("/register", routes.RegisterDeviceHandler)
			deviceGroup.PUT("/update", routes.UpdateDeviceHandler)
			deviceGroup.DELETE("/remove", routes.RemoveDeviceHandler)
		}

		pluginsGroup := apiV1Group.Group("/plugins")
		pluginsGroup.Use(UserAuthenticationMiddleware())
		{
			pluginsGroup.GET("/:plugin_name", routes.GetPluginProviderInfoHandler)
			pluginsGroup.GET("/all", routes.GetAllPluginProvidersInfoHandler)
			pluginsGroup.GET("/limited", routes.GetLimitedPluginProvidersInfoHandler)
			pluginsGroup.GET("/paginated", routes.GetPaginatedPluginProvidersInfoHandler)
		}

		devicesGroup := apiV1Group.Group("/devices")
		devicesGroup.Use(UserAuthenticationMiddleware())
		{
			devicesGroup.GET("/all", routes.GetAllDevicesHandler)
			devicesGroup.GET("/limited", routes.GetLimitedDevicesHandler)
			devicesGroup.GET("/paginated", routes.GetPaginatedDevicesHandler)
		}
		
		devicesInteractionsGroup := apiV1Group.Group("/devices/interactions/:client_id")
		// devicesInteractionsGroup.Use(DeviceAuthenticationMiddleware())
		devicesInteractionsGroup.Use(bridger.RouteEnabledMiddleware())
		{
			devicesInteractionsGroup.GET("/*directory", bridger.InteractionsHandler)
			devicesInteractionsGroup.POST("/*directory", bridger.InteractionsHandler)
		}

		DevicesInteractionsGroup = devicesInteractionsGroup
	}

	errCh := make(chan error)
	go func() {
		errCh <- r.Run(":" + strconv.Itoa(int(Config.Port)))
	}()

	return errCh
}

func Cleanup() error {
	log.Info("cleaning up api/v1")
	return nil
}
