package main

import (
	"network/common/pluginer"

	"github.com/charmbracelet/log"
)

type SamplePlugin struct {
	pluginer.PluginBase
}

func (p *SamplePlugin) Initialize() {
	log.Info("sample plugin initialized")
}

func (p *SamplePlugin) Execute() {
	log.Info("sample plugin executed")
}

func (p *SamplePlugin) Cleanup() {
	log.Info("sample plugin cleaned up")
}

func (p *SamplePlugin) NewPlugin() pluginer.Plugin {
	return &SamplePlugin{}
}