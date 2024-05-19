package configurator

import (
	"os"

	"smuggr/net-work/common/logger"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var MyLogger = logger.NewCustomLogger("conf")
var Config GlobalConfig

func loadEnv() *logger.MessageWrapper {
	err := godotenv.Load()
	if err != nil {
		return logger.ErrReadingEnvFile
	}

	return logger.MsgEnvFileLoaded
}

func loadConfig(config *GlobalConfig) *logger.MessageWrapper {
	if err := viper.ReadInConfig(); err != nil {
		MyLogger.Error("reading config", "error", err.Error())
		return logger.ErrReadingConfigFile
	}

	err := viper.Unmarshal(config)
	if err != nil {
		return logger.ErrFormattingConfigFile
	}

	return logger.MsgConfigFileLoaded.Format(viper.ConfigFileUsed())
}

func Initialize() {
	MyLogger.Info("initializing configurator")
	MyLogger.Log(loadEnv())

	viper.AddConfigPath(os.Getenv("CONFIG_PATH"))
	viper.SetConfigType(os.Getenv("CONFIG_TYPE"))

	MyLogger.Log(loadConfig(&Config))
}
