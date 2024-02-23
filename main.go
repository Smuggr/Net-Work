package main

import (
	"log"
	"network/data/configuration"
	"network/services/api"
	"network/services/bridge"
	"network/services/database"
	"os"
	"os/signal"
	"syscall"
)



func main() {
	var config configuration.Config
	configuration.Initialize(&config)

	go database.Initialize(&config.Database)
	go bridge.Initialize(&config.Bridge)
	go api.Initialize(&config.API)

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}

		log.Println("cleaning up...")
		if err := bridge.Cleanup(&config.Bridge); err != nil {
			log.Println(err.Error())
		}

		if err := api.Cleanup(&config.API); err != nil {
			log.Println(err.Error())
		}

		if err := database.Cleanup(&config.Database); err != nil {
			log.Println(err.Error())
		}

		os.Exit(1)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	log.Println("waiting for termination signal...")
	<-c
	log.Println("termination signal received")
}
