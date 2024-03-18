package bridger

import (
	"net/http"

	"network/common/provider"
	"network/utils/errors"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func RouteEnabledMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.Param("client_id")
		directory := c.Param("directory")
		
		_, err := GetClient(clientID)
		if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrClientNotFound.Format(clientID)})
            c.Abort()
            return
        }

		log.Debug("route enabled middleware", "clientID", clientID, "directory", directory)

		devicePlugin, _err := provider.GetDevicePlugin(clientID)
		if _err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrDevicePluginNotFound.Format(clientID)})
            c.Abort()
            return
        }

		log.Debug(devicePlugin.Routes)

		c.Set("client_id", clientID)
		c.Set("directory", directory)

		c.Next()
	}
}
