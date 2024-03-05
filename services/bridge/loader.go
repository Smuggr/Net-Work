package bridge

import (
	"log"
	"os"
	"path/filepath"
	"plugin"

	"network/common/pluginer"
)

func loadSOFile(file string) error {
	log.Println("loading SO file:", file)

	p, err := plugin.Open(file)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("plugin file opened", p)

	newPluginSymbol, err := p.Lookup("NewPlugin")
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Type of sym: %T\n", newPluginSymbol)

	log.Println("plugin instance found")

	NewPlugin, ok := newPluginSymbol.(func() pluginer.Plugin)
	if !ok {
		log.Fatalln("PluginInstance is not of type func() Plugin")
	}

	log.Println("executing plugin file")

	plugin := NewPlugin()
	plugin.Initialize()

	return nil
}

func InitializeLoader() error {
	subdirs, err := os.ReadDir(Config.PluginsDirectory)
	if err != nil {
		return err
	}

	for _, subdir := range subdirs {
		if subdir.IsDir() {
			files, err := filepath.Glob(filepath.Join(Config.PluginsDirectory, subdir.Name(), "*.so"))
			if err != nil {
				log.Fatal(err)
				return err
			}

			for _, file := range files {
				err := loadSOFile(file)
				if err != nil {
					log.Printf("Error loading plugin from file %s: %v", file, err)
				}
			}
		}
	}

	log.Println("plugins loaded")

	return nil
}
