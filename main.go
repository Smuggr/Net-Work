package main

import (
	"fmt"

	"overseer/services/api"
	"overseer/services/database"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)


type TimeConfig struct {
	NTPServer string `mapstructure:"ntp_server"`
}

type WebServerConfig struct {
	Port int `mapstructure:"port"`
}

type APIConfig struct {
	JWTLifespanMinutes int `mapstructure:"jwt_lifespan_minutes"`
}

type Config struct {
	Time      TimeConfig      `mapstructure:"time"`
	WebServer WebServerConfig `mapstructure:"web_server"`
	API APIConfig             `mapstructure:"api"`
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