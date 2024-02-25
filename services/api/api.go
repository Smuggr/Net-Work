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


func Initialize(config *configuration.APIConfig, ch chan error) { 
    log.Println("initializing api/v1")

	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()
	l := tollbooth.NewLimiter(1, nil)

	apiV1Group := r.Group("/api/v1")
	apiV1Group.Use(tollbooth_gin.LimitHandler(l))
	{
		userGroup := apiV1Group.Group("/user")
		userGroup.Use(UserAuthMiddleware())
		{
			userGroup.POST("/register", routes.RegisterUser)
			userGroup.POST("/update", routes.UpdateUser)
		}

		noAuthUserGroup := apiV1Group.Group("/user")
		{
			noAuthUserGroup.POST("/authenticate", routes.AuthenticateUser)
		}

		// deviceGroup := apiV1Group.Group("/device")
		// {
		// 	deviceGroup.POST("/register", routes.RegisterDevice)
		// }
	}

	http.Handle("/", r)

	if err := http.ListenAndServe(":" + strconv.Itoa(int(config.Port)), r); err != nil {
		ch <- err
	}
}

func Cleanup(config *configuration.APIConfig) error {
	log.Println("cleaning up api/v1")
	return nil
}