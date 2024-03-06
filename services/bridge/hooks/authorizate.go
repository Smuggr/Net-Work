package hooks

import (
	"bytes"
	"network/data/errors"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/wind-c/comqtt/v2/mqtt"
)

type AuthorizateHookConfig struct {
	Server *mqtt.Server
}

type AuthorizateHook struct {
	mqtt.HookBase
}

var AuthorHookConfig *AuthorizateHookConfig

func (h *AuthorizateHook) ID() string {
	return "authorizate"
}

func (h *AuthorizateHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnACLCheck,
	}, []byte{b})
}

func (h *AuthorizateHook) Init(config any) error {
	authorConfig, ok := config.(*AuthorizateHookConfig)
	if !ok {
		return errors.ErrInvalidHookConfig
	}

	AuthorHookConfig = authorConfig

	log.Info("initialized hook /v1/authorizate")
	return nil
}

func (h *AuthorizateHook) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	log.Debug("client", cl.ID, "wanted to authorizate", topic, "for write", write)

	if strings.Contains(topic, "/v1/device/" + cl.ID + "/") {
		return true
	}

	if strings.HasPrefix(topic, "/v1/devices/") {
		return true
	}

	return false
}
