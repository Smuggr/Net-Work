package configurator

type DatastorerConfig struct{}

type BridgerConfig struct {
	BrokerHost              string `mapstructure:"broker_host"`
	BrokerPort              string `mapstructure:"broker_port"`
	MDNSServiceInstanceName string `mapstructure:"mdns_service_instance_name"`
	MDNSDomain              string `mapstructure:"mdns_domain"`
	MDNSHostName            string `mapstructure:"mdns_host_name"`
	ClientId                string `mapstructure:"client_id"`
}

type PluginerConfig struct {
	PluginsDirectory string `mapstructure:"plugins_directory"`
}

type APIConfig struct {
	Port               uint     `mapstructure:"port"`
	JWTLifespanMinutes uint     `mapstructure:"jwt_lifespan_minutes"`
	TrustedCORS        []string `mapstructure:"trusted_cors"`
}

type GlobalConfig struct {
	Datastorer DatastorerConfig `mapstructure:"datastorer"`
	Bridger    BridgerConfig    `mapstructure:"bridger"`
	Pluginer   PluginerConfig   `mapstructure:"pluginer"`
	API        APIConfig        `mapstructure:"api"`
}