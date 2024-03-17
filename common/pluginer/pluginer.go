package pluginer

type PluginMetadata struct {
	APIVersion  string `json:"api_version"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Source      string `json:"source"`
}

type PluginInfo struct {
	Directory string          `json:"directory"`
	Metadata  *PluginMetadata `json:"metadata" gorm:"embedded"`
}

type PluginCallbacks interface {
	OnLoaded()     error
	OnCleaningUp() error
}

type PluginProvider struct {
	Info      *PluginInfo                  `json:"info"`
	NewPlugin func(string) (Plugin, error) `json:"-"`
	Callbacks PluginCallbacks              `json:"-"`
}

type Plugin interface {
	Execute() error
	Cleanup() error
}
