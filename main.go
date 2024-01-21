package main

import (
	"fmt"

	"github.com/spf13/viper"

	"overseer/services/apiv1"
)


type Config struct {
	Time struct {
		NtpServer string `mapstructure:"ntp_server"`
	}
}


func main()  {
	viper.SetConfigFile("config.toml")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Unable to decode config into struct:", err)
		return
	}

	fmt.Println(config.Time.NtpServer)

	apiv1.Initialize()

}