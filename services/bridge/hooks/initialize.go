package hooks

import (
	"bytes"
	"log"
	"network/data/errors"

	"github.com/wind-c/comqtt/v2/mqtt"
	"github.com/wind-c/comqtt/v2/mqtt/packets"
)

type InitializeDeviceHookConfig struct {
	Server *mqtt.Server
}

type InitializeDeviceHook struct {
	mqtt.HookBase
}

var InitDeviceHookConfig *InitializeDeviceHookConfig


func (h *InitializeDeviceHook) ID() string {
	return "authentication"
}

func (h *InitializeDeviceHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnect,
		mqtt.OnDisconnect,
	}, []byte{b})
}

func (h *InitializeDeviceHook) Init(config any) error {
	initConfig, ok := config.(*InitializeDeviceHookConfig)
	if !ok {
		return errors.ErrInvalidHookConfig
	}

	InitDeviceHookConfig = initConfig

	log.Println("initialized hook /v1/initialization")
	return nil
}

func (h *InitializeDeviceHook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	log.Println("client", cl.ID, "connected")
	
	for clientID, client := range InitDeviceHookConfig.Server.Clients.GetAll() {
		log.Println("Client ID:", clientID)
		log.Println("Client:", client.Properties.Username)
	}

	return nil
}

func (h *InitializeDeviceHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	if err != nil {
		log.Println("client", cl.ID, "disconnected", "expire", expire, "error", err)
	} else {
		log.Println("client", cl.ID, "disconnected", "expire", expire)
	}
}