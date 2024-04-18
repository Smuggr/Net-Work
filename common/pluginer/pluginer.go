package pluginer

import (
	"github.com/wind-c/comqtt/v2/mqtt"
)

type PluginMetadata struct {
	APIVersion  string `json:"api_version"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Source      string `json:"source"`
}

type PluginCallbacks struct {
	OnLoaded     func(*PluginProvider) error
	OnCleaningUp func(*PluginProvider) error
}

type PluginFactory func(string) (*Plugin, error)

type PluginProvider struct {
	Metadata  *PluginMetadata `json:"metadata" gorm:"embedded"`
	Callbacks PluginCallbacks `json:"-"`
	Factory   PluginFactory   `json:"-"`
}

type PluginMethods struct {
	Execute func(*Plugin) error
	Cleanup func(*Plugin) error
}

type Plugin struct {
	Methods  *PluginMethods
	Client   *mqtt.Client
	Router   *Router
	Metadata *PluginMetadata
}
