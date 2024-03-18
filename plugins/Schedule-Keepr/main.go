package main

import (
	"embed"
	"encoding/json"
	"net/http"

	"network/common/bridger"
	"network/common/pluginer"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

//go:embed static/*
var StaticDirectory embed.FS

var Plugin *pluginer.Plugin

type Callbacks struct{}
type Methods struct{}

func (p *Methods) Execute() error {
	log.Info("plugin Schedule-Keepr executed")

	return nil
}

func (p *Methods) Cleanup() error {
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

func NewPlugin(clientID string) (*pluginer.Plugin, error) {
	log.Info("initializing plugin Schedule-Keepr")

	// Tautological condition? WTF
	client, err := bridger.GetClient(clientID)
	if err != nil {
		return nil, err
	}

	var routes = make(map[string]interface{})
	routes["hello"] = bridger.BridgerRoute{Method: http.MethodGet, Callback: func(c *gin.Context) {
		log.Info("hello from Schedule-Keepr")
		c.JSON(http.StatusNotFound, gin.H{"amongus": "sussy bbaka"})
	}}

	routes["getters"] = make(map[string]interface{})
	routes["getters"].(map[string]interface{})["greet"] = bridger.BridgerRoute{
		Method: http.MethodGet,
		Callback: func(c *gin.Context) {
			log.Info("hello from Schedule-Keepr")
			c.JSON(http.StatusNotFound, gin.H{"amongus": "sussy bbaka"})
		},
	}

	log.Info("plugin Schedule-Keepr initialized", "client", clientID, "routes", routes)

	Plugin = &pluginer.Plugin{}
	Plugin.Methods = &Methods{}
	Plugin.Client = client
	Plugin.Routes = routes

	return Plugin, nil
}
