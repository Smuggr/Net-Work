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
		mqtt.OnConnect,
		mqtt.OnDisconnect,
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

func (h *AuthenticationHook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	log.Debug("connected", "client", cl.ID)

	for _, cl := range AuthHookConfig.Server.Clients.GetAll() {
		log.Debug("already connected", "client", cl.ID)
	}

	return nil
}

func (h *AuthenticationHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	if err != nil {
		log.Debug("disconnected", "client", cl.ID, "expire", expire, "error", err)
	} else {
		log.Debug("disconnected", "client", cl.ID, "expire", expire)
	}
}

func (h *AuthenticationHook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	log.Debug("authenticating", "client", cl.ID, "username", string(pk.Connect.Username))

	device := database.GetDevice(database.DB, cl.ID)
	if device == nil {
		return false
	}

	device = database.GetDeviceByUsername(database.DB, string(pk.Connect.Username))
	if device == nil {
		return false
	}

	if err := database.AuthenticateDevicePassword(device, string(pk.Connect.Password)); err != nil {
		return false
	}

	log.Debug("authenticated", "client", cl.ID, "username", string(pk.Connect.Username))

	return true
}
