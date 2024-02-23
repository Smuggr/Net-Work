package bridge

import (
	"log"
	"time"

	"network/data/configuration"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MainClient mqtt.Client


func Initialize(config *configuration.BridgeConfig) error {
	log.Println("initializing bridge/v1")

	options := mqtt.NewClientOptions()
	options.AddBroker(config.BrokerUrl)
	options.SetClientID(config.ClientId)
	options.SetKeepAlive(time.Duration(config.KeepAliveSeconds) * time.Second)
	options.SetAutoReconnect(true)
	options.SetCleanSession(false)

	MainClient = mqtt.NewClient(options)

	if token := MainClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}


func Cleanup(config *configuration.BridgeConfig) error {
	log.Println("cleaning up bridge/v1")
	MainClient.Disconnect(config.DisconnectMiliseconds)
	
	return nil
}