package config

type ServerConfig struct {
	Domain   string `mapstructure:"DOMAIN"`
	Host     string `mapstructure:"HOST"`
	GRPCPort string `mapstructure:"GRPC_PORT"`
	GWPort   string `mapstructure:"GW_PORT"`
}
