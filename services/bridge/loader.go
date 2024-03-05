package bridge

import (
	"log"
	"os"
	"path/filepath"
	"plugin"

	"network/common/pluginer"
	"network/data/errors"
)

var NewPluginsLoaded map[string]func() pluginer.Plugin

func loadSOFile(file string) (func() pluginer.Plugin, error) {
	log.Println("loading SO file:", file)

	p, err := plugin.Open(file)
	if err != nil {
		return nil, err
	}

	log.Println("plugin file opened", p)

	newPluginSymbol, err := p.Lookup("NewPlugin")
	if err != nil {
		return nil, err
	}

	NewPlugin, ok := newPluginSymbol.(func() pluginer.Plugin)
	if !ok {
		return nil, errors.ErrLookingUpPluginSymbol.Format(file)
	}

	log.Println("executing plugin file")

	plugin := NewPlugin()
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

	NewPluginsLoaded = make(map[string]func() pluginer.Plugin)
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
				} else {
					NewPluginsLoaded[file] = NewPlugin
				}
			}
		}
	}

	log.Println("plugins loaded")
	return failedPlugins, nil
}
