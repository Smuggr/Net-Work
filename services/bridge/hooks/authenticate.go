package hooks

import (
	"bytes"
	"network/services/database"
	"network/utils/errors"

	"github.com/charmbracelet/log"
	"github.com/wind-c/comqtt/v2/mqtt"
	"github.com/wind-c/comqtt/v2/mqtt/packets"
)

type AuthenticationHookConfig struct {
	Server *mqtt.Server
}

type AuthenticationHook struct {
	mqtt.HookBase
}

var AuthHookConfig *AuthenticationHookConfig

func (h *AuthenticationHook) ID() string {
	return "authentication"
}

func (h *AuthenticationHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnectAuthenticate,
	}, []byte{b})
}

func (h *AuthenticationHook) Init(config any) error {
	authConfig, ok := config.(*AuthenticationHookConfig)
	if !ok {
		return errors.ErrInvalidHookConfig
	}

	AuthHookConfig = authConfig

	log.Info("initialized hook /v1/authentication")
	return nil
}

func (h *AuthenticationHook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	log.Info("authenticating", "client", cl.ID, "username", string(pk.Connect.Username))

	_notAuthenticated := func() {
		log.Warn("authentication failed", "client", cl.ID, "username", string(pk.Connect.Username))
	}

	device := database.GetDevice(cl.ID)
	if device == nil {
		_notAuthenticated()
		return false
	}

	device = database.GetDeviceByUsername(string(pk.Connect.Username))
	if device == nil {
		_notAuthenticated()
		return false
	}

	if err := database.AuthenticateDevicePassword(device, string(pk.Connect.Password)); err != nil {
		_notAuthenticated()
		return false
	}

	log.Info("authenticated", "client", cl.ID, "username", string(pk.Connect.Username))

	return true
}
