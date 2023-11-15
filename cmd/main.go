package main

import (
	"context"
	"time"

	"github.com/blanergol/crypto_rate_bot/config"
	"github.com/blanergol/crypto_rate_bot/internal/handler"

	"github.com/adshao/go-binance/v2"

	tele "gopkg.in/telebot.v3"

	"github.com/blanergol/crypto_rate_bot/internal/usecase"
)

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()

	binanceClient := binance.NewClient(cfg.BinanceApiKey, cfg.BinanceSecretKey)

	teleConf := tele.Settings{
		Token:  cfg.TelegramToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(teleConf)
	if err != nil {
		return
	}

	useCases := usecase.New(binanceClient)

	h := handler.NewHandlers(binanceClient, useCases, cfg)

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Бот помогает отслеживать курсы криптовалют и отправлять оповещения об изменении их цены.")
	})

	bot.Handle("/current", func(c tele.Context) error {
		return h.Tokens(ctx, c)
	})

	bot.Handle("/price", func(c tele.Context) error {
		return h.Price(ctx, c)
	})

	bot.Handle("/price_background", func(c tele.Context) error {
		go GetPriceTask(h, ctx, c)
		return nil
	})

	bot.Start()
}

func GetPriceTask(h handler.Handlers, ctx context.Context, c tele.Context) {
	for {
		time.Sleep(time.Minute * 1)
		h.PriceTask(ctx, c)
	}
}
