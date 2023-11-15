package handlers

import (
	"context"
	"fmt"

	tele "gopkg.in/telebot.v3"

	"github.com/blanergol/crypto_rate_bot/internal/helpers"
	"github.com/blanergol/crypto_rate_bot/internal/usecase"
)

func PriceHandler(ctx context.Context, c tele.Context, useCases *usecase.UseCasesIml) error {
	listPriceFilter := c.Args()
	if len(listPriceFilter) == 1 {
		return c.Send("Error args")
	}

	listTokens, err := useCases.GetListTokensWithCurrentPrice(ctx, listPriceFilter[:len(listPriceFilter)-1])
	if err != nil {
		return c.Send("Error get list tokens")
	}

	prices, err := useCases.GetPriceForListTokens(ctx, *listTokens, listPriceFilter[len(listPriceFilter)-1])
	if err != nil {
		return c.Send("Error get prices")
	}

	var respList []string
	for _, p := range *prices {
		str := fmt.Sprintf("- Symbols: %s; PriceChange: %s; PriceChangePercent: %s%%;\n", p.Symbol, p.PriceChange, p.PriceChangePercent)
		respList = append(respList, str)
	}

	return helpers.SendTelegramWith400Words(c, respList)
}
