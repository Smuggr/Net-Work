package bridge

import (
	"fmt"
	"log"

	"network/data/configuration"
	"network/services/bridge/hooks"

	"github.com/wind-c/comqtt/v2/mqtt"
	"github.com/wind-c/comqtt/v2/mqtt/hooks/auth"
	"github.com/wind-c/comqtt/v2/mqtt/listeners"
)

var Config *configuration.BridgeConfig
var Server *mqtt.Server

func Initialize() error {
	log.Println("initializing bridge/v1")

	Config = &configuration.Config.Bridge

	Server = mqtt.New(&mqtt.Options{
		InlineClient: true,
	})

	_ = Server.AddListener(listeners.NewTCP("t1", ":" + fmt.Sprint(Config.BrokerPort), nil))
	_ = Server.AddHook(new(auth.AllowHook), nil)
	_ = Server.AddHook(new(hooks.AuthenticationHook), nil)

	go func() {
		err := Server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}



func Cleanup() error {
	log.Println("cleaning up bridge/v1")

	Server.Close()

	return nil
}