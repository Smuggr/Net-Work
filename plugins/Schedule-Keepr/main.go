package main

import (
	"embed"
	"encoding/json"

	// "network/common/bridger"
	"network/common/pluginer"
	"network/smugwork"

	"github.com/charmbracelet/log"
)

//go:embed static/*
var StaticDirectory embed.FS

func getMetadata() (*pluginer.PluginMetadata, error) {
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

func Execute(plugin *pluginer.Plugin) error {
	log.Info("plugin Schedule-Keepr executed")

	r := plugin.Router

	gettersGroup := r.Group("getters", pluginer.HandlersChain{
		"GET": { func(ctx *pluginer.Context) {
			ctx.JSON(200, "fab!")
		},},
	})
	{
		barGettersGroup := gettersGroup.Group("bar", nil)
		{
			group := barGettersGroup.GET("fab", func(ctx *pluginer.Context) {
				ctx.JSON(200, "fab!")
			})

			group.GET("fem", func(ctx *pluginer.Context) {
				ctx.JSON(200, "fem!")
			})

			barGettersGroup.GET("fam", func(ctx *pluginer.Context) {
				ctx.JSON(200, "fam!")
			})
		}

		farGettersGroup := gettersGroup.Group(":far", nil)
		{
			farGettersGroup.GET("fob", func(ctx *pluginer.Context) {
				ctx.JSON(200, "fob!")
			})
		}
	}

	return nil
}

func Cleanup(plugin *pluginer.Plugin) error {
	log.Info("plugin Schedule-Keepr cleaned up")

	return nil
}

func OnLoaded(pluginProvider *pluginer.PluginProvider) error {
	log.Info("loaded sample plugin provider Schedule-Keepr callback")

	return nil
}

func OnCleaningUp(pluginProvider *pluginer.PluginProvider) error {
	log.Info("cleaning up plugin provider Schedule-Keepr callback")

	return nil
}

func NewPlugin(clientID string) (*pluginer.Plugin, error) {
	log.Info("initializing plugin Schedule-Keepr")

	return smugwork.InitializePlugin(clientID, &pluginer.PluginMethods{
		Execute: Execute,
		Cleanup: Cleanup,
	})
}

func NewPluginProvider() (*pluginer.PluginProvider, error) {
	log.Info("initializing plugin provider for Schedule-Keepr")

	var newPluginFactory pluginer.PluginFactory = NewPlugin

	metadata, err := getMetadata()
	if err != nil {
		return nil, err
	}

	return smugwork.InitializePluginProvider(newPluginFactory, metadata, pluginer.PluginCallbacks{
		OnLoaded:     OnLoaded,
		OnCleaningUp: OnCleaningUp,
	})
}
