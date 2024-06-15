package routes

import (
	"smuggr.xyz/net-work/api/v1/handlers"
	"smuggr.xyz/net-work/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

func InitializePluginsRoutes(rootGroup *gin.RouterGroup) {
	pluginsGroup := rootGroup.Group("/plugins")
	pluginsGroup.Use(middleware.UserAuthentication())
	{
		pluginsGroup.GET("/:plugin_name", handlers.GetPluginProviderInfoHandler)
		pluginsGroup.GET("/all", handlers.GetAllPluginProvidersInfoHandler)
		pluginsGroup.GET("/limited", handlers.GetLimitedPluginProvidersInfoHandler)
		pluginsGroup.GET("/paginated", handlers.GetPaginatedPluginProvidersInfoHandler)
	}
}
