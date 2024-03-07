package bridge

import (
	"os"
	"path/filepath"
	"plugin"

	"github.com/charmbracelet/log"

	"network/common/pluginer"
	"network/utils/errors"
)

var LoadedPluginConstructors map[string]func() (pluginer.Plugin, error)

func loadSOFile(file string) ((func() (pluginer.Plugin, error)), error) {
	log.Debug("loading plugin", "file", file)

	p, err := plugin.Open(file)
	if err != nil {
		return nil, err
	}

	log.Debug("opened plugin", "file", file)

	newPluginSymbol, err := p.Lookup("NewPlugin")
	if err != nil {
		return nil, err
	}

	log.Debugf("type %T", newPluginSymbol)

	NewPlugin, ok := newPluginSymbol.(func() (pluginer.Plugin, error))
	if !ok {
		return nil, errors.ErrLookingUpPluginSymbol.Format(file)
	}

	log.Debug("executing plugin", "file", file)

	plugin, err := NewPlugin()
	if err != nil {
		return nil, err
	}

	plugin.Initialize()
	plugin.Execute()
	plugin.Cleanup()

	return NewPlugin, nil
}

func InitializeLoader() (map[string]error, error) {
	subdirs, err := os.ReadDir(Config.PluginsDirectory)
	if err != nil {
		return nil, err
	}

	LoadedPluginConstructors = make(map[string]func() (pluginer.Plugin, error))
	failedPlugins := make(map[string]error)

	for _, subdir := range subdirs {
		if subdir.IsDir() {
			files, err := filepath.Glob(filepath.Join(Config.PluginsDirectory, subdir.Name(), "*.so"))
			if err != nil {
				log.Fatal(err)
				return nil, err
			}

			for _, file := range files {
				NewPlugin, err := loadSOFile(file)
				if err != nil {
					failedPlugins[file] = err
					log.Error("failed to load plugin", "file", file, "error", err)
				} else {
					LoadedPluginConstructors[file] = NewPlugin
					log.Debug("successfully loaded plugin", "file", file)
				}
			}
		}
	}

	log.Info("plugins:", "loaded", len(LoadedPluginConstructors), "failed", len(failedPlugins), "out of", len(subdirs))
	for key := range LoadedPluginConstructors {
		log.Debug("Key:", key)
	}

	return failedPlugins, nil
}
