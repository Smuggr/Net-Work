package routes

import (
	"smuggr.xyz/net-work/api/v1/handlers"
	// "smuggr.xyz/net-work/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUsersRoutes(rootGroup *gin.RouterGroup) {
	userGroup := rootGroup.Group("/user")
	// userGroup.Use(middleware.UserAuthentication())
	{
		userGroup.GET("/:login", handlers.GetUserHandler)
		userGroup.POST("/register", handlers.RegisterUserHandler)
		userGroup.PUT("/update", handlers.UpdateUserHandler)
		userGroup.DELETE("/remove", handlers.RemoveUserHandler)
	}

	usersGroup := rootGroup.Group("/users")
	// usersGroup.Use(middleware.UserAuthentication())
	{
		usersGroup.GET("/all", handlers.GetAllUsersHandler)
		usersGroup.GET("/limited", handlers.GetLimitedUsersHandler)
		usersGroup.GET("/paginated", handlers.GetPaginatedUsersHandler)
	}
}
