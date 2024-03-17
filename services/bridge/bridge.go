package bridge

import (
	"fmt"

	"network/services/bridge/hooks"
	"network/utils/configuration"

	"github.com/charmbracelet/log"
	"github.com/hashicorp/mdns"
	"github.com/wind-c/comqtt/v2/mqtt"
	"github.com/wind-c/comqtt/v2/mqtt/listeners"
)

var Config *configuration.BridgeConfig
var MQTTServer *mqtt.Server
var MDNSServer *mdns.Server

func registerHooks() error {
	if err := MQTTServer.AddHook(new(hooks.AuthenticationHook), &hooks.AuthenticationHookConfig{
		Server: MQTTServer,
	}); err != nil {
		return err
	}

	if err := MQTTServer.AddHook(new(hooks.InitializeDeviceHook), &hooks.InitializeDeviceHookConfig{
		Server: MQTTServer,
	}); err != nil {
		return err
	}

	if err := MQTTServer.AddHook(new(hooks.AuthorizateHook), &hooks.AuthorizateHookConfig{
		Server: MQTTServer,
	}); err != nil {
		return err
	}

	return nil
}

func Initialize() error {
	log.Info("initializing bridge/v1")

	Config = &configuration.Config.Bridge

	MQTTServer = mqtt.New(&mqtt.Options{
		InlineClient: true,
	})

	tcp := listeners.NewTCP("t1", ":"+fmt.Sprint(Config.BrokerPort), nil)
	_ = MQTTServer.AddListener(tcp)

	registerHooks()

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
	log.Info("cleaning up bridge/v1")

	err := MQTTServer.Close()
	if err != nil {
		return err
	}

	err = MDNSServer.Shutdown()
	if err != nil {
		return err
	}

	return nil
}
