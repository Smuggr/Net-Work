package api

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"network/data/configuration"
	"network/services/api/routes"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)


var Config = &configuration.Config.API

func Initialize() chan error {
	log.Println("initializing api/v1")

	Config = &configuration.Config.API

	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()
	l := tollbooth.NewLimiter(1, nil)

	apiV1Group := r.Group("/api/v1")
	apiV1Group.Use(tollbooth_gin.LimitHandler(l))
	{
		userGroup := apiV1Group.Group("/user")
		userGroup.Use(UserAuthenticationMiddleware())
		{
			userGroup.POST("/register", routes.RegisterUserHandler)
			userGroup.PUT("/update", routes.UpdateUserHandler)
		}

		noAuthUserGroup := apiV1Group.Group("/user")
		{
			noAuthUserGroup.POST("/authenticate", routes.AuthenticateUserHandler)
		}

		usersGroup := apiV1Group.Group("/users")
		usersGroup.Use(UserAuthenticationMiddleware())
		{
			usersGroup.GET("/all", routes.GetAllUsersHandler)
			usersGroup.GET("/limited", routes.GetLimitedUsersHandler)
			usersGroup.GET("/paginated", routes.GetPaginatedUsersHandler)
		}


		deviceGroup := apiV1Group.Group("/device")
		{
			deviceGroup.POST("/register", routes.RegisterDeviceHandler)
			deviceGroup.PUT("/update", routes.UpdateDeviceHandler)
		}

		devicesGroup := apiV1Group.Group("/devices")
		devicesGroup.Use(UserAuthenticationMiddleware())
		{
			devicesGroup.GET("/all", routes.GetAllDevicesHandler)
			devicesGroup.GET("/limited", routes.GetLimitedDevicesHandler)
			devicesGroup.GET("/paginated", routes.GetPaginatedDevicesHandler)
		}
	}

	http.Handle("/", r)

	errCh := make(chan error)
	go func() {
		if err := http.ListenAndServe(":" + strconv.Itoa(int(Config.Port)), r); err != nil {
			errCh <- err
		}
	}()

	return errCh
}

func Cleanup() error {
	log.Println("cleaning up api/v1")
	return nil
}