package main

import (
	"context"
	"fmt"
	"github.com/blanergol/crypto_rate_bot/config"

	"github.com/adshao/go-binance/v2"

	"github.com/blanergol/crypto_rate_bot/internal/models"
	"github.com/blanergol/crypto_rate_bot/internal/usecase"
)

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()

	binanceClient := binance.NewClient(cfg.ApiKey, cfg.SecretKey)

	useCases := usecase.New(binanceClient)

	tokens, err := useCases.GetListTokensWithCurrentPrice(ctx, models.USDT)
	if err != nil {
		return
	}

	for _, v := range *tokens {
		fmt.Printf("Token: %s; Symbols: %s/%s; CurrentPrice: %s\n", v.Name, v.Name, v.Symbol, v.Price)
	}

	fmt.Println("______________________________________")

	prices, err := useCases.GetPriceForListTokens(ctx, *tokens, "5m")
	if err != nil {
		return
	}

	for _, v := range *prices {
		fmt.Printf("Symbols: %s; PriceChange: %s; PriceChangePercent: %s%%;\n", v.Symbol, v.PriceChange, v.PriceChangePercent)
	}
}
