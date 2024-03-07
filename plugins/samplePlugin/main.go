package main

import (
	"embed"
	"encoding/json"

	"network/common/pluginer"

	"github.com/charmbracelet/log"
)

//go:embed resources/*
var ResourcesDirectory embed.FS


type SamplePlugin struct {
	Metadata *pluginer.PluginMetadata
}

func (p *SamplePlugin) Initialize() error {
	log.Info("sample plugin initialized")

	return nil
}

func (p *SamplePlugin) Execute() error {
	log.Info("sample plugin executed")

	return nil
}

func (p *SamplePlugin) Cleanup() error {
	log.Info("sample plugin cleaned up")

	return nil
}

func NewPlugin() (pluginer.Plugin, error) {
	metadataFile, err := ResourcesDirectory.Open("resources/metadata.json")
	if err != nil {
		return nil, err
	}
	defer metadataFile.Close()

	var metadata pluginer.PluginMetadata
	err = json.NewDecoder(metadataFile).Decode(&metadata)
	if err != nil {
		return nil, err
	}

	// metadata, err := pluginer.GetMetadataFromFile(metadataFile)
	// if err != nil {
	// 	return nil, err
	// }

	log.Debug("metadata: ", metadata)

	return &SamplePlugin{
		Metadata: &metadata,
	}, nil
}