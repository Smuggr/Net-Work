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
	Client    string          `json:"client"`
}

type PluginProvider struct {
	Info      *PluginInfo            `json:"info"`
	NewPlugin func() (Plugin, error) `json:"-"`
	OnLoaded  func() error           `json:"-"`
}

type Plugin interface {
	Initialize() error
	Execute() error
	Cleanup() error
}