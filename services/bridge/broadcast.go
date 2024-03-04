package bridge

import (
	"github.com/hashicorp/mdns"
)

func InitializeMDNS() error {
	service, err := mdns.NewMDNSService(Config.MDNSServiceInstanceName, "_mqtt._tcp", Config.MDNSDomain, Config.MDNSHostName, int(Config.BrokerPort), nil, nil)
	if err != nil {
		return err
	}

	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return err
	}

	MDNSServer = server

	return nil
}