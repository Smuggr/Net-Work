package bridger

import (
	"network/utils/errors"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/wind-c/comqtt/v2/mqtt"
	"github.com/wind-c/comqtt/v2/mqtt/packets"
)

// Indexed by directory
type BridgerRoute struct {
	Method    string
	Callback  func(c *gin.Context)
}

var MQTTServer *mqtt.Server
var InteractionsGroup *gin.RouterGroup

func GetClient(clientID string) (*mqtt.Client, *errors.ErrorWrapper) {
	client, ok := MQTTServer.Clients.Get(clientID)
	if !ok {
		return nil, errors.ErrClientNotFound.Format(clientID)
	}

	return client, nil
}

func GetAllClients() map[string]*mqtt.Client {
	return MQTTServer.Clients.GetAll()
}

func DisconnectClient(clientID string) error {
	client, err := GetClient(clientID)
	if err != nil {
		return err
	}

	MQTTServer.DisconnectClient(client, packets.ErrNotAuthorized)

	return nil
}

func InteractionsGETHandler(c *gin.Context) {

}

func InteractionsPOSTHandler(c *gin.Context) {

}

func Initialize(mqttServer *mqtt.Server, interactionsGroup *gin.RouterGroup) error {
	log.Info("initializing bridger/v1")

	MQTTServer = mqttServer
	InteractionsGroup = interactionsGroup

	return nil
}

func Cleanup() error {
	log.Info("cleaning up bridger/v1")

	return nil
}