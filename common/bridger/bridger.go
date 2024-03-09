package bridger

import (
	"network/utils/errors"

	"github.com/wind-c/comqtt/v2/mqtt"
	"github.com/wind-c/comqtt/v2/mqtt/packets"
)

var MQTTServer *mqtt.Server

func GetClient(clientID string) (*mqtt.Client, *errors.ErrorWrapper) {
	client, ok := MQTTServer.Clients.Get(clientID)
	if !ok {
		return nil, errors.ErrClientNotFound.Format(clientID)
	}

	return client, nil
}

func DisconnectClient(clientID string) error {
	client, err := GetClient(clientID)
	if err != nil {
		return err
	}

	MQTTServer.DisconnectClient(client, packets.ErrNotAuthorized)

	return nil
}

func Initialize(mqttServer *mqtt.Server) error {
	MQTTServer = mqttServer

	return nil
}