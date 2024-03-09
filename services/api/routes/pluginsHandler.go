package routes

import (
	"net/http"
	"strconv"

	"network/services/provider"
	"network/utils/errors"
	"network/utils/messages"

	"github.com/gin-gonic/gin"
)

func GetPluginProviderInfoHandler(c *gin.Context) {
	pluginName := c.Param("plugin_name")

	pluginProvider := provider.LoadedPluginProviders[pluginName]
	if pluginProvider == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrPluginProviderNotFound.Format(pluginName)})
		return
	}

	info := *pluginProvider.Info

	c.JSON(http.StatusOK, gin.H{
		"message":  messages.MsgPluginProviderInfoFetchSuccess.Format(pluginName),
		"provider": info,
	})
}

func GetAllPluginProvidersInfoHandler(c *gin.Context) {
	providers := make(map[string]interface{})
	for pluginName, pluginProvider := range provider.LoadedPluginProviders {
		info := *pluginProvider.Info
		providers[pluginName] = info
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   messages.MsgPluginProvidersInfoFetchSuccess.Format(len(providers)),
		"providers": providers,
	})
}

func GetLimitedPluginProvidersInfoHandler(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	providers := make(map[string]interface{})
	count := 0
	for pluginName, pluginProvider := range provider.LoadedPluginProviders {
		if count >= limit {
			break
		}
		info := *pluginProvider.Info
		providers[pluginName] = info
		count++
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   messages.MsgPluginProvidersInfoFetchSuccess.Format(len(providers)),
		"limit":     limit,
		"providers": providers,
	})
}

func GetPaginatedPluginProvidersInfoHandler(c *gin.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	providers := make(map[string]interface{})
	count := 0
	for pluginName, pluginProvider := range provider.LoadedPluginProviders {
		if count >= (page-1)*pageSize && count < page*pageSize {
			info := *pluginProvider.Info
			providers[pluginName] = info
		}
		count++
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   messages.MsgPluginProvidersInfoFetchSuccess.Format(len(providers)),
		"page":      page,
		"pageSize":  pageSize,
		"providers": providers,
	})
}
