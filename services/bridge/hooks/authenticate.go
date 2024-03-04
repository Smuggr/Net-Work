package hooks

import (
	"bytes"
	"log"
	"network/data/errors"
	"network/services/database"

	"github.com/wind-c/comqtt/v2/mqtt"
	"github.com/wind-c/comqtt/v2/mqtt/packets"
)

type AuthenticationHookConfig struct {
	Server *mqtt.Server
}

type AuthenticationHook struct {
	mqtt.HookBase
}

var Config *AuthenticationHookConfig


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

	Config = authConfig

	log.Println("initialized hook /v1/authentication")
	return nil
}

func (h *AuthenticationHook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	log.Println("client", cl.ID, "connected")

	for clientID, client := range Config.Server.Clients.GetAll() {
		log.Println("Client ID:", clientID)
		log.Println("Client:", client.Properties.Username)
	}

	return nil
}

func (h *AuthenticationHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	if err != nil {
		log.Println("client", cl.ID, "disconnected", "expire", expire, "error", err)
	} else {
		log.Println("client", cl.ID, "disconnected", "expire", expire)
	}
}

func (h *AuthenticationHook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	log.Println("client", cl.ID, "wanted to authenticate as", string(pk.Connect.Username))

	device := database.GetDevice(database.DB, string(pk.Connect.Username))
	if device == nil {
		return false
	}

	if err := database.AuthenticateDevicePassword(device, string(pk.Connect.Password)); err != nil {
		return false
	}

	return true
}