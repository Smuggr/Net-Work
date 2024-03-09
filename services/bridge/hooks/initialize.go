package hooks

import (
	"bytes"
	"network/utils/errors"

	"github.com/charmbracelet/log"
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
	return "initialization"
}

func (h *InitializeDeviceHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnSessionEstablished,
		mqtt.OnDisconnect,
	}, []byte{b})
}

// Create device plugin in loader
func (h *InitializeDeviceHook) OnSessionEstablished(cl *mqtt.Client, pk packets.Packet) {
	log.Info("session established", "client", cl.ID)

	for _, cl := range InitDeviceHookConfig.Server.Clients.GetAll() {
		log.Debug("already connected", "client", cl.ID)
	}
}

// Remove device plugin from map in loader
func (h *InitializeDeviceHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	if err != nil {
		log.Info("disconnected", "client", cl.ID, "expire", expire, "error", err)
	} else {
		log.Info("disconnected", "client", cl.ID, "expire", expire)
	}
}

func (h *InitializeDeviceHook) Init(config any) error {
	initConfig, ok := config.(*InitializeDeviceHookConfig)
	if !ok {
		return errors.ErrInvalidHookConfig
	}

	InitDeviceHookConfig = initConfig

	log.Info("initialized hook /v1/initialization")
	return nil
}
