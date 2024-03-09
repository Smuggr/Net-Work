package main

import (
	"embed"
	"encoding/json"

	// "network/common/bridger"
	"network/common/pluginer"

	"github.com/charmbracelet/log"
)

//go:embed static/*
var StaticDirectory embed.FS

// put into PluginBase struct?
type SamplePlugin struct {
	Metadata *pluginer.PluginMetadata
}

func (p *SamplePlugin) Initialize() error {
	log.Info("initializing sample plugin")

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

func GetMetadata() (*pluginer.PluginMetadata, error) {
	metadataFile, err := StaticDirectory.Open("static/metadata.json")
	if err != nil {
		return nil, err
	}
	defer metadataFile.Close()

	var metadata pluginer.PluginMetadata
	err = json.NewDecoder(metadataFile).Decode(&metadata)
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}

func OnLoaded() error {
	log.Info("loaded sample plugin provider")

	return nil
}

func NewPlugin() (pluginer.Plugin, error) {
	metadata, err := GetMetadata()
	if err != nil {
		return nil, err
	}

	log.Debug("loaded", "metadata", metadata)

	return &SamplePlugin{
		Metadata: metadata,
	}, nil
}
