package bridger

import (
	"network/common/pluginer"
	"network/utils/errors"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/wind-c/comqtt/v2/mqtt"
	"github.com/wind-c/comqtt/v2/mqtt/packets"
)

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

func InteractionsHandler(c *gin.Context) {
	group, _ := c.Get("group")
	group.(*pluginer.RouterGroup).Execute(c.Request.Method, c)
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