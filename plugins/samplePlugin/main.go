package main

import (
	"log"

	"network/common/pluginer"
)

type SamplePlugin struct {}

func (p *SamplePlugin) Initialize() {
	log.Println("Sample plugin initialized")
}

func (p *SamplePlugin) Execute() {
	log.Println("Sample plugin executed")
}

func (p *SamplePlugin) Cleanup() {
	log.Println("Sample plugin cleaned up")
}

func NewPlugin() pluginer.Plugin {
	return &SamplePlugin{}
}