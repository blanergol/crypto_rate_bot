package handler

import (
	"context"
	"fmt"

	tele "gopkg.in/telebot.v3"

	"github.com/blanergol/crypto_rate_bot/internal/helpers"
)

type Tokener interface {
	Tokens(ctx context.Context) error
}

func (h Handlers) Tokens(ctx context.Context, c tele.Context) error {
	tokensFilter := c.Args()

	tokens, err := h.usecase.GetListTokensWithCurrentPrice(ctx, tokensFilter)
	if err != nil {
		return c.Send("Error get list tokens")
	}

	var respList []string
	for _, token := range *tokens {
		str := fmt.Sprintf("<strong>Token</strong>: %s; <strong>Symbols</strong>: <a href='https://www.binance.com/en/trade/%s?type=spot'>%s%s</a>; <strong>Price</strong>: %s\n", token.Name, token.Name+token.Symbol, token.Name, token.Symbol, token.Price)
		respList = append(respList, str)
	}

	return helpers.SendTelegramMessage(c, respList)
}
