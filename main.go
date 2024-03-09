package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"

	"network/common/bridger"
	"network/services/api"
	"network/services/bridge"
	"network/services/database"
	"network/services/provider"
	"network/utils/configuration"
)

func Initialize() {
	if _, err := configuration.Initialize(); err != nil {
		log.Fatal(err.Error())
	}

	if err := database.Initialize(); err != nil {
		log.Fatal(err.Error())
	}

	apiChan := api.Initialize()
	go func() {
		if err := <-apiChan; err != nil {
			// Fatal or Debug?
			log.Fatal(err.Error())
		}
	}()

	if err := bridge.Initialize(); err != nil {
		log.Fatal(err.Error())
	}

	if err := bridger.Initialize(bridge.MQTTServer, api.DevicesInteractionsGroup); err != nil {
		log.Fatal(err.Error())
	}

	_, err := provider.Initialize()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := database.InitializeDevices(); err != nil {
		log.Fatal(err.Error())
	}
}

func Cleanup() {
	if err := recover(); err != nil {
		log.Warn("panic", err)
	}

	log.Info("cleaning up...")
	if err := bridge.Cleanup(); err != nil {
		log.Error(err.Error())
	}

	if err := api.Cleanup(); err != nil {
		log.Error(err.Error())
	}

	if err := database.Cleanup(); err != nil {
		log.Error(err.Error())
	}
}

func WaitForTermination() {
	callChan := make(chan os.Signal, 1)
	signal.Notify(callChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	log.Debug("waiting for termination signal...")
	<-callChan
	log.Debug("termination signal received")
}

func main() {
	Initialize()

	defer Cleanup()
	defer WaitForTermination()
}
