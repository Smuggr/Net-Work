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

		client, err := GetClient(clientID)
		if err != nil {
			log.Warn("client not found", "client", clientID)

			c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrClientNotFound.Format(clientID)})
			c.Abort()
			return
		}

		plugin, _err := provider.GetDevicePlugin(clientID)
		if _err != nil {
			log.Warn("plugin not found", "client", clientID)

			c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrDevicePluginNotFound.Format(clientID)})
			c.Abort()
			return
		}

		group := plugin.Router.GetGroup(directory)
		if group == nil {
			log.Warn("group not found", "client", clientID, "directory", directory)

			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.Set("plugin", plugin)
		c.Set("group", group)
		c.Set("client", client)

		c.Next()
	}
}
