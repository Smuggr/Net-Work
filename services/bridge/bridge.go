package bridge

import (
	"fmt"
	"log"

	"network/data/configuration"
	"network/services/bridge/hooks"

	"github.com/hashicorp/mdns"
	"github.com/wind-c/comqtt/v2/mqtt"
	"github.com/wind-c/comqtt/v2/mqtt/listeners"
)

var Config *configuration.BridgeConfig
var MQTTServer *mqtt.Server
var MDNSServer *mdns.Server

func Initialize() error {
	log.Println("initializing bridge/v1")

	Config = &configuration.Config.Bridge

	MQTTServer = mqtt.New(&mqtt.Options{
		InlineClient: true,
	})

	tcp := listeners.NewTCP("t1", ":" + fmt.Sprint(Config.BrokerPort), nil)
	_ = MQTTServer.AddListener(tcp)
	_ = MQTTServer.AddHook(new(hooks.AuthenticationHook), &hooks.AuthenticationHookConfig{
		Server: MQTTServer,
	})

	if err := InitializeMDNS(); err != nil {
		return err
	}

	go func() {
		err := MQTTServer.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}



func Cleanup() error {
	log.Println("cleaning up bridge/v1")

	MQTTServer.Close()
	MDNSServer.Shutdown()

	return nil
}