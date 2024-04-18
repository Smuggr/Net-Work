package routes

import (
	"net/http"
	"strconv"

	"network/common/provider"
	"network/utils/errors"
	"network/utils/messages"

	"github.com/gin-gonic/gin"
)

// TO-DO: Change plugin provider info to metadata

func GetPluginProviderInfoHandler(c *gin.Context) {
	pluginName := c.Param("plugin_name")

	pluginProvider := provider.LoadedPluginProviders[pluginName]
	if pluginProvider == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrPluginProviderNotFound.Format(pluginName)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  messages.MsgPluginProviderInfoFetchSuccess.Format(pluginName),
		"metadata": *pluginProvider.Metadata,
	})
}

func GetAllPluginProvidersInfoHandler(c *gin.Context) {
	providersMetadata := make(map[string]interface{})
	for pluginName, pluginProvider := range provider.LoadedPluginProviders {
		metadata := *pluginProvider.Metadata
		providersMetadata[pluginName] = metadata
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   messages.MsgPluginProvidersInfoFetchSuccess.Format(len(providersMetadata)),
		"metadatas": providersMetadata,
	})
}

func GetLimitedPluginProvidersInfoHandler(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	providersMetadata := make(map[string]interface{})
	count := 0
	for pluginName, pluginProvider := range provider.LoadedPluginProviders {
		if count >= limit {
			break
		}

		metadata := *pluginProvider.Metadata
		providersMetadata[pluginName] = metadata
		count++
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   messages.MsgPluginProvidersInfoFetchSuccess.Format(len(providersMetadata)),
		"limit":     limit,
		"metadatas": providersMetadata,
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

	providersMetadata := make(map[string]interface{})
	count := 0
	for pluginName, pluginProvider := range provider.LoadedPluginProviders {
		if count >= (page-1)*pageSize && count < page*pageSize {
			metadata := *pluginProvider.Metadata
			providersMetadata[pluginName] = metadata
		}
		count++
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   messages.MsgPluginProvidersInfoFetchSuccess.Format(len(providersMetadata)),
		"page":      page,
		"pageSize":  pageSize,
		"metadatas": providersMetadata,
	})
}
