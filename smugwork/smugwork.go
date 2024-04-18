package smugwork

import (
	"network/common/bridger"
	"network/common/pluginer"

	"github.com/charmbracelet/log"
)

// TO-DO: Add custom loggers for plugin providers

func InitializePluginProvider(factory pluginer.PluginFactory, metadata *pluginer.PluginMetadata, callbacks pluginer.PluginCallbacks) (*pluginer.PluginProvider, error) {
	log.Info("initializing plugin provider")

	var PluginProvider *pluginer.PluginProvider = &pluginer.PluginProvider{}
	PluginProvider.Metadata = metadata
	PluginProvider.Callbacks = callbacks
	PluginProvider.Factory = factory

	return PluginProvider, nil
}

func InitializePlugin(clientID string, methods *pluginer.PluginMethods) (*pluginer.Plugin, error) {
	log.Info("initializing plugin")

	// Tautological condition? WTF
	client, err := bridger.GetClient(clientID)
	if err != nil {
		return nil, err
	}

	var Plugin *pluginer.Plugin = &pluginer.Plugin{}
	Plugin.Methods = methods
	Plugin.Client = client
	Plugin.Router = pluginer.NewRouter()

	return Plugin, nil
}