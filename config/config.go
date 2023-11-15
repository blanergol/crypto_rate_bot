package config

import "github.com/spf13/viper"

const (
	envBinanceApiKey    = "BINANCE_API_KEY"
	envBinanceSecretKey = "BINANCE_API_SECRET"
	envTelegramToken    = "TELEGRAM_TOKEN"
	envPriceTaskNotify  = "PRICE_TASK_NOTIFY"

	priceTaskNotify = 3.0 // percent
)

const (
	envPrefix = "CRYPTO_RATE"
)

type Config struct {
	BinanceApiKey    string
	BinanceSecretKey string
	TelegramToken    string
	PriceTaskNotify  float64
}

func NewConfig() *Config {
	v := viper.New()
	v.SetEnvPrefix(envPrefix)
	v.SetDefault(envPriceTaskNotify, priceTaskNotify)
	v.AutomaticEnv()

	return &Config{
		BinanceApiKey:    v.GetString(envBinanceApiKey),
		BinanceSecretKey: v.GetString(envBinanceSecretKey),
		TelegramToken:    v.GetString(envTelegramToken),
		PriceTaskNotify:  v.GetFloat64(envPriceTaskNotify),
	}
}
