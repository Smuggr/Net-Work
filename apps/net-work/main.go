package main

import (
	"smuggr/net-work/common/configurator"
	"smuggr/net-work/common/logger"

	"smuggr/net-work/core/datastorer"
)

var Logger = logger.DefaultLogger

func main() {
	logger.Initialize()
	configurator.Initialize()

	datastorer.Logger.Log(datastorer.Initialize())
}
