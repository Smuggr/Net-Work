package configuration

import (
	"network/data/errors"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
}

type BridgeConfig struct {
	BrokerHost              string `mapstructure:"broker_host"`
	BrokerPort              uint   `mapstructure:"broker_port"`
	MDNSServiceInstanceName string `mapstructure:"mdns_service_instance_name"`
	MDNSDomain              string `mapstructure:"mdns_domain"`
	MDNSHostName            string `mapstructure:"mdns_host_name"`
	ClientId                string `mapstructure:"client_id"`
	KeepAliveSeconds        uint   `mapstructure:"keep_alive_seconds"`
	DisconnectMiliseconds   uint   `mapstructure:"disconnect_miliseconds"`
}

type APIConfig struct {
	Port               uint `mapstructure:"port"`
	JWTLifespanMinutes uint `mapstructure:"jwt_lifespan_minutes"`
}

type GlobalConfig struct {
	Database DatabaseConfig `mapstructure:"database"`
	Bridge   BridgeConfig   `mapstructure:"bridge"`
	API      APIConfig      `mapstructure:"api"`
}


var Config GlobalConfig

func loadEnv() error {
	err := godotenv.Load()
    if err != nil {
        return errors.ErrReadingEnvFile
    }

	return nil
}

func loadConfig(config *GlobalConfig) error {
	viper.SetConfigFile("config.toml")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return errors.ErrReadingConfigFile
	}

	err := viper.Unmarshal(config)
	if err != nil {
		return errors.ErrFormattingConfigFile
	}

	return nil
}

func Initialize() (*GlobalConfig, error) {
	if err := loadConfig(&Config); err != nil {
		return nil, err
	}
	
	if err := loadEnv(); err != nil {
		return nil, err
	}

	return &Config, nil
}