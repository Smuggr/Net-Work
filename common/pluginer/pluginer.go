package pluginer

import (
	"encoding/json"
	"os"
)

type Plugin interface {
	Initialize()
	Execute()
	Cleanup()
	NewPlugin() Plugin
}

type PluginMetadata struct {
	APIVersion  string `json:"api_version"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Source      string `json:"source"`
}

type PluginBase struct {
	Metadata PluginMetadata
}

func NewPlugin() (*PluginBase, error) {
	metadataFile, err := os.Open("/home/karol/Documents/Repositories/Test/network/common/pluginer/metadata.json")
	if err != nil {
		return nil, err
	}
	defer metadataFile.Close()

	var metadata PluginMetadata
	err = json.NewDecoder(metadataFile).Decode(&metadata)
	if err != nil {
		return nil, err
	}

	return &PluginBase{
		Metadata: metadata,
	}, nil
}
