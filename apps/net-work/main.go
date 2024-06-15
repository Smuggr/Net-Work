package main

import (
	"smuggr.xyz/net-work/common/configurator"
	"smuggr.xyz/net-work/common/logger"

	api "smuggr.xyz/net-work/api/v1"
	"smuggr.xyz/net-work/core/datastorer"
)

var Logger = logger.DefaultLogger

func main() {
	logger.Initialize()
	configurator.Initialize()

	datastorer.Logger.Log(datastorer.Initialize())

	apiMsg, apiChan := api.Initialize()
	api.Logger.Log(apiMsg)

	go func() {
		if msg := <-apiChan; msg != nil {
			api.Logger.Log(msg)
		}
	}()

}
