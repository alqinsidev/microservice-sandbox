package config

import (
	"time"

	"github.com/spf13/viper"
)

type JWTConfig struct {
	AccessSecretKey  string
	RefreshSecretKey string
	AccessLifetime   time.Duration
	RefreshLifeTime  time.Duration
}

func LoadJWTConfig() JWTConfig {
	return JWTConfig{
		AccessSecretKey:  viper.GetString("JWT_ACCESS_SECRET_KEY"),
		RefreshSecretKey: viper.GetString("JWT_REFRESH_SECRET_KEY"),
		AccessLifetime:   viper.GetDuration("JWT_ACCESS_LIFETIME") * time.Second,
		RefreshLifeTime:  viper.GetDuration("JWT_REFRESH_LIFETIME") * time.Second,
	}
}
