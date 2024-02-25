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
	if err := configuration.Initialize(&config); err != nil {
		log.Fatalln(err.Error())
	}

	if err := database.Initialize(&config.Database); err != nil {
		log.Fatalln(err.Error())
	}

	if err := bridge.Initialize(&config.Bridge) ; err != nil {
		log.Fatalln(err.Error())
	}
	
	apiChan := make(chan error)
	go api.Initialize(&config.API, apiChan)

	go func() {
		if err := <-apiChan; err != nil {
			log.Println(err.Error()) 
		}
	}()

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

	}()

	callChan := make(chan os.Signal, 1)
	signal.Notify(callChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	log.Println("waiting for termination signal...")
	<-callChan
	log.Println("termination signal received")
}

