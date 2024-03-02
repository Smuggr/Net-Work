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

func Initialize() {
	if _, err := configuration.Initialize(); err != nil {
		log.Fatalln(err.Error())
	}

	if err := database.Initialize(); err != nil {
		log.Fatalln(err.Error())
	}

	apiChan := api.Initialize()
	go func() {
		if err := <-apiChan; err != nil {
			log.Println(err.Error())
		}
	}()

	if err := bridge.Initialize(); err != nil {
		log.Fatalln(err.Error())
	}
}

func Cleanup() {
	if err := recover(); err != nil {
		log.Println("panic occurred:", err)
	}

	log.Println("cleaning up...")
	if err := bridge.Cleanup(); err != nil {
		log.Println(err.Error())
	}

	if err := api.Cleanup(); err != nil {
		log.Println(err.Error())
	}

	if err := database.Cleanup(); err != nil {
		log.Println(err.Error())
	}
}

func WaitForTermination() {
	callChan := make(chan os.Signal, 1)
	signal.Notify(callChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	log.Println("waiting for termination signal...")
	<-callChan
	log.Println("termination signal received")
}

func main() {
	Initialize()

	defer Cleanup()
	defer WaitForTermination()
}