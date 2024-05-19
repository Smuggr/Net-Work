package main

import (
	"smuggr/net-work/common/configurator"
	"smuggr/net-work/common/logger"

	"smuggr/net-work/core/datastorer"
)

var MyLogger = logger.DefaultLogger

func main() {
	logger.Initialize()
	configurator.Initialize()

	if err := datastorer.Initialize(); err != nil && err.IsError() {
		datastorer.MyLogger.Log(err)
	}
}
