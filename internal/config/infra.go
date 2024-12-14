package config

import "fmt"

type DBConfig struct {
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Name     string `mapstructure:"NAME"`
	Port     string `mapstructure:"PORT"`
	SSLMode  string `mapstructure:"SSLMODE"`
}

type RedisConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Password string `mapstructure:"PASSWORD"`
	DB       int    `mapstructure:"DB"`
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.DB.Host, c.DB.User, c.DB.Password, c.DB.Name, c.DB.Port, c.DB.SSLMode,
	)
}
