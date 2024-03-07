package pluginer

type PluginMetadata struct {
	APIVersion  string `json:"api_version"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Source      string `json:"source"`
}

// Error channels instead?
type Plugin interface {
	Initialize() error
	Execute() error
	Cleanup() error
}

// func GetMetadataFromFile(file fs.File) (*PluginMetadata, error) {
// 	metadataFile, err := fs.ReadFile(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &metadata, nil
// }
