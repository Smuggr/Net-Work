package hooks

import (
	"bytes"
	"network/utils/errors"

	"github.com/charmbracelet/log"
	"github.com/wind-c/comqtt/v2/mqtt"
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

	log.Info("initialized hook /v1/initialization")
	return nil
}
