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
	DefaultCredit    int64 `mapstructure:"DEFAULT"`
	Level1Thereshold int64 `mapstructure:"LEVEL1"`
	Level2Thereshold int64 `mapstructure:"LEVEL2"`
	Level3Thereshold int64 `mapstructure:"LEVEL3"`
	Level4Thereshold int64 `mapstructure:"LEVEL4"`
	Level5Thereshold int64 `mapstructure:"LEVEL5"`
}
