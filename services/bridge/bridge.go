package bridge

import (
	"log"
	"time"

	"overseer/data/configuration"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)


func Initialize(config *configuration.BridgeConfig) {
	log.Println("initializing bridge/v1")

	options := mqtt.NewClientOptions()
	options.AddBroker(config.BrokerUrl)
	options.SetClientID(config.ClientId)
	options.SetKeepAlive(time.Duration(config.KeepAliveSeconds) * time.Second)
	options.SetAutoReconnect(true)
	options.SetCleanSession(false)

	client := mqtt.NewClient(options)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client.Disconnect(config.DisconnectMiliseconds)
}
