package bridge

import (
	"log"
	"time"

	"network/data/configuration"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MainClient mqtt.Client
var Config *configuration.BridgeConfig

func Initialize() error {
	log.Println("initializing bridge/v1")

	Config = &configuration.Config.Bridge

	options := mqtt.NewClientOptions()
	options.AddBroker(Config.BrokerUrl)
	options.SetClientID(Config.ClientId)
	options.SetKeepAlive(time.Duration(Config.KeepAliveSeconds) * time.Second)
	options.SetAutoReconnect(true)
	options.SetCleanSession(false)

	MainClient = mqtt.NewClient(options)

	if token := MainClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}


func Cleanup() error {
	log.Println("cleaning up bridge/v1")
	MainClient.Disconnect(Config.DisconnectMiliseconds)
	
	return nil
}