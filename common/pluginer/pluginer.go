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

type PluginProvider struct {
	Info       *PluginInfo
	NewPlugin  func() (Plugin, error)
}

// Error channels instead?
type Plugin interface {
	Initialize() error
	Execute() error
	Cleanup() error
}
