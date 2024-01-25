package main

import (
	"fmt"

	"github.com/spf13/viper"

	"overseer/services/api"
	"overseer/services/database"
)


type TimeConfig struct {
	NTPServer string `mapstructure:"ntp_server"`
}

type WebServerConfig struct {
	Port int `mapstructure:"port"`
}

type Config struct {
	Time      TimeConfig      `mapstructure:"time"`
	WebServer WebServerConfig `mapstructure:"web_server"`
}


func loadConfig(config *Config) error {
	viper.SetConfigFile("config.toml")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return err
	}

	err := viper.Unmarshal(config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}



func main()  {
	var config Config
	loadConfig(&config)

	fmt.Println(config)

	database.Initialize(config.Database.DatabasePath)
	api.Initialize(config.WebServer.Port)
}