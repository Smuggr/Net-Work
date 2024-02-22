package main

import (
	"overseer/data/configuration"
	"overseer/services/api"
	"overseer/services/bridge"
	"overseer/services/database"
)



func main() {
	var config configuration.Config
	configuration.Initialize(&config)

	database.Initialize(&config.Database)
	bridge.Initialize(&config.Bridge)
	api.Initialize(&config.API)
}

