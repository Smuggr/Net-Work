package routes

import (
	"smuggr.xyz/net-work/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func SetupNoAuthRoutes(rootGroup *gin.RouterGroup) {
	noAuthUserGroup := rootGroup.Group("/user")
	{
		noAuthUserGroup.POST("/authenticate", handlers.AuthenticateUserHandler)
		noAuthUserGroup.POST("/validateToken", handlers.ValidateUserTokenHandler)
	}
}
