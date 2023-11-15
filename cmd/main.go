package main

import (
	"context"
	"time"

	"github.com/blanergol/crypto_rate_bot/config"
	"github.com/blanergol/crypto_rate_bot/internal/handlers"

	"github.com/adshao/go-binance/v2"

	tele "gopkg.in/telebot.v3"

	"github.com/blanergol/crypto_rate_bot/internal/usecase"
)

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()

	binanceClient := binance.NewClient(cfg.BinanceApiKey, cfg.BinanceSecretKey)

	useCases := usecase.New(binanceClient)

	pref := tele.Settings{
		Token:  cfg.TelegramToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		return
	}

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Бот помогает отслеживать курсы криптовалют и отправлять оповещения об изменении цены.")
	})

	bot.Handle("/current", func(c tele.Context) error {
		return handlers.TokensHandler(ctx, c, useCases)
	})

	bot.Handle("/price", func(c tele.Context) error {
		return handlers.PriceHandler(ctx, c, useCases)
	})

	bot.Start()
}
