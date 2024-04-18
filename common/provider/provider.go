package provider

import (
	"os"
	"path/filepath"
	"plugin"

	"network/common/pluginer"
	"network/utils/configuration"
	"network/utils/constants"
	"network/utils/errors"

	"github.com/charmbracelet/log"
)

var Config *configuration.ProviderConfig

// Indexed by PluginName
var LoadedPluginProviders map[string]*pluginer.PluginProvider

// Indexed by ClientID
var DevicesPlugins map[string]*pluginer.Plugin

func findPluginProviderConflicts(pluginProvider *pluginer.PluginProvider) error {
	metadata := pluginProvider.Metadata

	if metadata.APIVersion != constants.APIVersion {
		return errors.ErrAPIVersionMismatch.Format(metadata.APIVersion, constants.APIVersion)
	}

	for _, otherPluginProvider := range LoadedPluginProviders {
		otherMetadata := otherPluginProvider.Metadata

		if metadata.Name == otherMetadata.Name {
			if metadata.Version == otherMetadata.Version && metadata.Author == otherMetadata.Author {
				return errors.ErrPluginProviderConflict.Format(metadata.Name)
			}
		}
	}

	return nil
}

func lookupProviders(pluginSymbol *plugin.Plugin, file string, pluginProvider *pluginer.PluginProvider) error {
	log.Debug("looking up provider", "file", file)

	newPluginProviderSymbol, err := pluginSymbol.Lookup("NewPluginProvider")
	log.Debugf("new plugin provider symbol type: %T", newPluginProviderSymbol)
	if err != nil {
		return err
	}

	NewPluginProvider, ok := newPluginProviderSymbol.(func() (*pluginer.PluginProvider, error))
	if !ok {
		return errors.ErrLookingUpPluginSymbol.Format(file)
	}

	newPluginProvider, err := NewPluginProvider()
	if err != nil {
		return err
	}

	*pluginProvider = *newPluginProvider

	// filepath.Dir(file)

	if err := findPluginProviderConflicts(pluginProvider); err != nil {
		return err
	}

	return nil
}

func loadSOFile(file string, pluginProvider *pluginer.PluginProvider) error {
	log.Debug("loading plugin", "file", file)

	pluginSymbol, err := plugin.Open(file)
	if err != nil {
		return err
	}

	log.Debug("opened plugin", "file", file)
	if err := lookupProviders(pluginSymbol, file, pluginProvider); err != nil {
		return err
	}

	log.Debug(pluginProvider, pluginProvider.Callbacks)

	if pluginProvider.Callbacks.OnLoaded != nil {
		pluginProvider.Callbacks.OnLoaded(pluginProvider)
	} else {
		log.Warn("no OnLoaded callback found", "provider", pluginProvider.Metadata.Name)
	}

	log.Debug("loaded plugin", "file", file)

	return nil
}

func loadPluginProviders() (map[string]error, error) {
	subdirs, err := os.ReadDir(Config.PluginsDirectory)
	if err != nil {
		return nil, err
	}

	LoadedPluginProviders = make(map[string]*pluginer.PluginProvider)
	DevicesPlugins = make(map[string]*pluginer.Plugin)

	failedPlugins := make(map[string]error)

	for _, subdir := range subdirs {
		if subdir.IsDir() {
			subdirName := subdir.Name()

			log.Debug("", "subdir", subdirName)

			files, err := filepath.Glob(filepath.Join(Config.PluginsDirectory, subdirName, constants.PluginSOFileName))
			if err != nil {
				log.Fatal(err)
				return nil, err
			}

			// Allow max 1 file per plugin?
			for _, file := range files {
				var NewPluginProvider pluginer.PluginProvider

				if err := LoadPluginProvider(subdirName, &NewPluginProvider); err != nil {
					failedPlugins[file] = err
					log.Error("failed to load plugin", "file", file, "error", err)
				} else {
					LoadedPluginProviders[subdirName] = &NewPluginProvider
					log.Debug("successfully loaded plugin provider", "file", file)
				}
			}
		}
	}

	log.Info("plugin providers:", "loaded", len(LoadedPluginProviders), "failed", len(failedPlugins), "out of", len(subdirs))
	for key, value := range LoadedPluginProviders {
		log.Debug("plugins provider", "key", key, "value", value)
	}

	return failedPlugins, nil
}

func GetPluginProvider(pluginName string) (*pluginer.PluginProvider, error) {
	provider, ok := LoadedPluginProviders[pluginName]
	if !ok {
		return nil, errors.ErrPluginProviderNotFound.Format(pluginName)
	}

	return provider, nil
}

func LoadPluginProvider(pluginName string, pluginProvider *pluginer.PluginProvider) error {
	log.Debug("loading plugin", "plugin", pluginName)

	existingPluginProvider, _ := GetPluginProvider(pluginName)
	if existingPluginProvider != nil {
		return errors.ErrPluginProviderAlreadyLoaded.Format(pluginName)
	}

	if err := loadSOFile(filepath.Join(Config.PluginsDirectory, pluginName, constants.PluginSOFileName), pluginProvider); err != nil {
		return err
	}

	return nil
}

func GetDevicePlugin(clientID string) (*pluginer.Plugin, error) {
	plugin, ok := DevicesPlugins[clientID]
	if !ok {
		return nil, errors.ErrDevicePluginNotFound.Format(clientID)
	}

	return plugin, nil
}

func CreateDevicePlugin(pluginName string, clientID string) (*pluginer.Plugin, error) {
	log.Debug("creating plugin", "plugin", pluginName, "client", clientID)

	pluginProvider, err := GetPluginProvider(pluginName)
	if err != nil {
		return nil, err
	}

	if _, err := GetDevicePlugin(clientID); err == nil {
		return nil, err
	}

	plugin, err := pluginProvider.Factory(clientID)
	if err != nil {
		return nil, err
	}

	DevicesPlugins[clientID] = plugin

	if plugin.Methods.Execute != nil {
		if err := plugin.Methods.Execute(plugin); err != nil {
			return nil, err
		}
	} else {
		log.Warn("no Execute method found", "client", clientID)
	}

	return plugin, nil
}

func RemoveDevicePlugin(clientID string) error {
	log.Debug("removing plugin", "client", clientID)

	plugin, err := GetDevicePlugin(clientID)
	if err != nil {
		log.Error("plugin not found", "client", clientID)
		return errors.ErrRemovingDevicePlugin.Format(clientID)
	}

	length := len(DevicesPlugins)
	log.Debug("before", "length", length)

	if plugin.Methods.Cleanup != nil {
		if err := plugin.Methods.Cleanup(plugin); err != nil {
			return err
		}
	} else {
		log.Warn("no Cleanup method found", "client", clientID)
	}

	delete(DevicesPlugins, clientID)

	log.Debug("after", "length", len(DevicesPlugins))

	if length == len(DevicesPlugins) {
		return errors.ErrRemovingDevicePlugin.Format(clientID)
	}

	return nil
}

func Initialize() (map[string]error, error) {
	Config = &configuration.Config.Provider

	log.Info("initializing provider/v1")

	failedPlugins, err := loadPluginProviders()
	if err != nil {
		return nil, err
	}

	return failedPlugins, nil
}

func Cleanup() error {
	for _, devicePlugin := range DevicesPlugins {
		log.Debug("cleaning up device plugin", "plugin", devicePlugin)

		if devicePlugin.Methods.Cleanup != nil {
			if err := devicePlugin.Methods.Cleanup(devicePlugin); err != nil {
				log.Error("error cleaning up device plugin", "plugin", devicePlugin, "error", err.Error())
				return err
			}
		} else {
			log.Warn("no Cleanup method found", "client", devicePlugin.Client.ID)
		}
	}

	for _, pluginProvider := range LoadedPluginProviders {
		log.Debug("cleaning up plugin provider", "plugin", pluginProvider)

		if pluginProvider.Callbacks.OnCleaningUp != nil {
			if err := pluginProvider.Callbacks.OnCleaningUp(pluginProvider); err != nil {
				log.Error("error cleaning up plugin provider", "plugin", pluginProvider, "error", err.Error())
				return err
			}
		} else {
			log.Warn("no OnCleanup callback found", "provider", pluginProvider.Metadata.Name)
		}
	}

	return nil
}
