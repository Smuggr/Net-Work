package main

import (
	"embed"
	"encoding/json"

	"network/common/bridger"
	"network/common/pluginer"

	"github.com/charmbracelet/log"
	"github.com/wind-c/comqtt/v2/mqtt"
)

//go:embed static/*
var StaticDirectory embed.FS

type Callbacks struct{}

type Plugin struct {
	Client *mqtt.Client
}

func (p *Plugin) Execute() error {
	log.Info("plugin Schedule-Keepr executed")

	return nil
}

func (p *Plugin) Cleanup() error {
	log.Info("plugin Schedule-Keepr cleaned up")

	return nil
}

func (p *Callbacks) OnLoaded() error {
	log.Info("loaded sample plugin provider Schedule-Keepr")

	return nil
}

func (p *Callbacks) OnCleaningUp() error {
	log.Info("cleaning up plugin provider Schedule-Keepr")

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

func GetCallbacks() (pluginer.PluginCallbacks, error) {
	return &Callbacks{}, nil
}

func NewPlugin(clientID string) (pluginer.Plugin, error) {
	log.Info("initializing plugin Schedule-Keepr")

	// Tautological condition? WTF
	client, err := bridger.GetClient(clientID)
	if err != nil {
		return nil, err
	}

	log.Info("plugin Schedule-Keepr initialized", "client", clientID)

	return &Plugin{
		Client: client,
	}, nil
}
