package routes

import (
	//"encoding/json"
	"net/http"

	// "network/common/pluginer"
	"network/services/provider"
	//"network/utils/errors"
	"network/utils/messages"

	//"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func GetAllPluginProvidersInfoHandler(c *gin.Context) {
	providers := make(map[string]interface{})
	for pluginName, pluginProvider := range provider.LoadedPluginProviders {
		info := *pluginProvider.Info
		providers[pluginName] = info
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   messages.MsgPluginProvidersInfoFetchSuccess,
		"providers": providers,
	})
}
