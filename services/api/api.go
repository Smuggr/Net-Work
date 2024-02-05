package api

import (
	"log"
	"net/http"
	"strconv"

	"overseer/services/api/routes"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)


func Initialize(port int) {
    log.Println("initializing api/v1")
    
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
	}

	http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))
}