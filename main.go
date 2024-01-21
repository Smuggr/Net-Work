package main

import (
	"fmt"

	"github.com/spf13/viper"

	"overseer/services/apiv1"
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
	viper.SetConfigFile("config.toml")
	viper.SetConfigType("toml")


	var config Config
	loadConfig(&config)

	fmt.Println(config)

	apiv1.Initialize(config.WebServer.Port)

}