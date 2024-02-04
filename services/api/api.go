package api

import (
	"log"
	"net/http"
	"strconv"

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
			userGroup.POST("/register", RegisterUser)
			userGroup.POST("/update", UpdateUser)
		}

		noAuthUserGroup := apiV1Group.Group("/user")
		{
			noAuthUserGroup.POST("/authenticate", AuthenticateUser)
		}
	}

	http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))
}