package bridge

import (
	"os"
	"path/filepath"
	"plugin"

	"github.com/charmbracelet/log"

	"network/common/pluginer"
	"network/utils/errors"
)

var LoadedPluginProviders map[string]*pluginer.PluginProvider

func lookupProviders(p *plugin.Plugin, file string, pluginProvider *pluginer.PluginProvider) error {
	log.Debug("looking up provider", "file", file)

	newPluginSymbol, err := p.Lookup("NewPlugin")
	if err != nil {
		return err
	}

	log.Debugf("plugin symbol type %T", newPluginSymbol)

	NewPlugin, ok := newPluginSymbol.(func() (pluginer.Plugin, error))
	if !ok {
		return errors.ErrLookingUpPluginSymbol.Format(file)
	}

	getMetadataSymbol, err := p.Lookup("GetMetadata")
	if err != nil {
		return err
	}

	log.Debugf("plugin symbol type %T", getMetadataSymbol)

	GetMetadata, ok := getMetadataSymbol.(func() (*pluginer.PluginMetadata, error))
	if !ok {
		return errors.ErrLookingUpPluginSymbol.Format(file)
	}

	metadata, err := GetMetadata()
	if err != nil {
		return err
	}

	log.Debug("loaded", "metadata", metadata)

	pluginProvider.NewPlugin = NewPlugin

	pluginProvider.Info = &pluginer.PluginInfo{
		Directory: filepath.Dir(file),
		Metadata:  metadata,
	}

	return nil
}

func loadSOFile(file string, pluginProvider *pluginer.PluginProvider) error {
	log.Debug("loading plugin", "file", file)

	p, err := plugin.Open(file)
	if err != nil {
		return err
	}

	log.Debug("opened plugin", "file", file)
	if err := lookupProviders(p, file, pluginProvider); err != nil {
		return err
	}

	log.Debug("executing plugin", "file", file)

	return nil
}

func InitializeLoader() (map[string]error, error) {
	subdirs, err := os.ReadDir(Config.PluginsDirectory)
	if err != nil {
		return nil, err
	}

	LoadedPluginProviders = make(map[string]*pluginer.PluginProvider)

	failedPlugins := make(map[string]error)

	for _, subdir := range subdirs {
		if subdir.IsDir() {
			subdirName := subdir.Name()

			files, err := filepath.Glob(filepath.Join(Config.PluginsDirectory, subdirName, "*.so"))
			if err != nil {
				log.Fatal(err)
				return nil, err
			}

			// Allow max 1 file per plugin?
			for _, file := range files {
				var NewPluginProvider pluginer.PluginProvider

				if err := loadSOFile(file, &NewPluginProvider); err != nil {
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
