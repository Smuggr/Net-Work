package main

import (
	"fmt"

	"overseer/services/api"
	"overseer/services/database"

	"github.com/spf13/viper"
	"github.com/joho/godotenv"
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

func loadEnv() {
	err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }
}


func main()  {
	var config Config
	loadConfig(&config)
	loadEnv()

	fmt.Println(config)

	database.Initialize()
	api.Initialize(config.WebServer.Port)
}