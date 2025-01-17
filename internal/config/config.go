package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment string       `mapstructure:"ENV"`
	SecretKey   string       `mapstructure:"SECRET_KEY"`
	Server      ServerConfig `mapstructure:"SERVER"`
	Token       TokenConfig  `mapstructure:"TOKEN"`
	Cred        CredConfig   `mapstructure:"CRED"`
	DB          DBConfig     `mapstructure:"DB"`
	Redis       RedisConfig  `mapstructure:"REDIS"`
	EmailSender EmailConfig  `mapstructure:"EMAIL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
