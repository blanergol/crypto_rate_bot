package handler

import (
	"context"
	"fmt"
	"strconv"

	tele "gopkg.in/telebot.v3"

	"github.com/blanergol/crypto_rate_bot/internal/helpers"
)

type Pricer interface {
	Price(ctx context.Context, c tele.Context) error
	PriceTask(ctx context.Context, c tele.Context) error
}

func (h Handlers) Price(ctx context.Context, c tele.Context) error {
	listPriceFilter := c.Args()
	if len(listPriceFilter) == 1 {
		return c.Send("Error args")
	}

	listTokens, err := h.usecase.GetListTokensWithCurrentPrice(ctx, listPriceFilter[:len(listPriceFilter)-1])
	if err != nil {
		return c.Send("Error get list tokens")
	}

	prices, err := h.usecase.GetPriceForListTokens(ctx, *listTokens, listPriceFilter[len(listPriceFilter)-1])
	if err != nil {
		return c.Send("Error get prices")
	}

	var respList []string
	for _, p := range *prices {
		percent, err := strconv.ParseFloat(p.PriceChangePercent, 64)
		if err != nil {
			return err
		}
		smile := "&#128185;"
		if percent < 0 {
			smile = "&#128219;"
		}

		str := fmt.Sprintf("%s<strong>Symbols</strong>: %s; <strong>Price Change</strong>: %s; <strong>Percent</strong>: %s%%;\n", smile, p.Symbol, p.PriceChange, p.PriceChangePercent)
		respList = append(respList, str)
	}

	return helpers.SendTelegramMessage(c, respList)
}

func (h Handlers) PriceTask(ctx context.Context, c tele.Context) error {
	listTokens, err := h.usecase.GetListTokensWithCurrentPrice(ctx, []string{})
	if err != nil {
		return c.Send("Error get list tokens")
	}

	prices, err := h.usecase.GetPriceForListTokens(ctx, *listTokens, "5m")
	if err != nil {
		return c.Send("Error get prices")
	}

	var respList []string
	for _, p := range *prices {
		priceChangePercent, err := strconv.ParseFloat(p.PriceChangePercent, 64)
		if err != nil {
			return err
		}
		if priceChangePercent > h.cfg.PriceTaskNotify {
			str := fmt.Sprintf("&#128185;<strong>Symbols</strong>: %s; <strong>Price</strong>: %s; <strong>Percent</strong>: %s%%;\n", p.Symbol, p.PriceChange, p.PriceChangePercent)
			respList = append(respList, str)
		}
	}

	return helpers.SendTelegramMessage(c, respList)
}
