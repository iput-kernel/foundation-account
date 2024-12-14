package config

import "time"

type EmailConfig struct {
	Name     string `mapstructure:"NAME"`
	Address  string `mapstructure:"ADDRESS"`
	Password string `mapstructure:"PASSWORD"`
}

type TokenConfig struct {
	AccessDuration  time.Duration `mapstructure:"ACCESS_DURATION"`
	RefreshDuration time.Duration `mapstructure:"REFRESH_DURATION"`
}

type CredConfig struct {
	DefaultCredit    int `mapstructure:"DEFAULT"`
	Level1Thereshold int `mapstructure:"LEVEL1"`
	Level2Thereshold int `mapstructure:"LEVEL2"`
	Level3Thereshold int `mapstructure:"LEVEL3"`
	Level4Thereshold int `mapstructure:"LEVEL4"`
	Level5Thereshold int `mapstructure:"LEVEL5"`
}
