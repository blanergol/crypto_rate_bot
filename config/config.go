package config

import "github.com/spf13/viper"

const (
	envApiKey    = "API_KEY"
	envSecretKey = "API_SECRET"
)

const (
	envPrefix = "CRYPTO_RATE"
)

type Config struct {
	ApiKey    string
	SecretKey string
}

func NewConfig() *Config {
	v := viper.New()
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	return &Config{
		ApiKey:    v.GetString(envApiKey),
		SecretKey: v.GetString(envSecretKey),
	}
}
