package configuration

import (
	"os"

	"network/utils/errors"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host string `mapstructure:"host"`
	Port uint   `mapstructure:"port"`
}

type BridgeConfig struct {
	BrokerHost              string `mapstructure:"broker_host"`
	BrokerPort              uint   `mapstructure:"broker_port"`
	MDNSServiceInstanceName string `mapstructure:"mdns_service_instance_name"`
	MDNSDomain              string `mapstructure:"mdns_domain"`
	MDNSHostName            string `mapstructure:"mdns_host_name"`
	ClientId                string `mapstructure:"client_id"`
}

type ProviderConfig struct {
	PluginsDirectory string `mapstructure:"plugins_directory"`
}

type APIConfig struct {
	Port               uint `mapstructure:"port"`
	JWTLifespanMinutes uint `mapstructure:"jwt_lifespan_minutes"`
}

type GlobalConfig struct {
	Database DatabaseConfig `mapstructure:"database"`
	Bridge   BridgeConfig   `mapstructure:"bridge"`
	Provider ProviderConfig `mapstructure:"provider"`
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
	viper.SetConfigFile(os.Getenv("CONFIG_FILE"))
	viper.SetConfigType(os.Getenv("CONFIG_TYPE"))

	if err := viper.ReadInConfig(); err != nil {
		log.Error("reading config", "error", err.Error())
		return errors.ErrReadingConfigFile
	}

	err := viper.Unmarshal(config)
	if err != nil {
		return errors.ErrFormattingConfigFile
	}

	log.Info("config loaded", "file", viper.ConfigFileUsed())

	return nil
}

func setupLogging() {
	log.SetLevel(log.DebugLevel)

	styles := log.DefaultStyles()

	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("[DEBUG]").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("#1E90FF")).
		Bold(true)

	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("[INFOR]").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("#CCCCCC")).
		Bold(true)

	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("[WARNI]").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("#FFA500")).
		Bold(true)

	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("[ERROR]").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true)

	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("[FATAL]").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("#8B0000")).
		Bold(true).
		Blink(true)

	log.SetStyles(styles)
}

func Initialize() (*GlobalConfig, error) {
	setupLogging()

	if err := loadEnv(); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if err := loadConfig(&Config); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &Config, nil
}
