package handlers

import (
	"context"
	"fmt"

	tele "gopkg.in/telebot.v3"

	"github.com/blanergol/crypto_rate_bot/internal/helpers"
	"github.com/blanergol/crypto_rate_bot/internal/usecase"
)

func TokensHandler(ctx context.Context, c tele.Context, useCases *usecase.UseCasesIml) error {
	tokensFilter := c.Args()

	tokens, err := useCases.GetListTokensWithCurrentPrice(ctx, tokensFilter)
	if err != nil {
		return c.Send("Error get list tokens")
	}

	var respList []string
	for _, token := range *tokens {
		str := fmt.Sprintf("- Token: %s; Symbols: %s/%s; CurrentPrice: %s\n", token.Name, token.Name, token.Symbol, token.Price)
		respList = append(respList, str)
	}

	return helpers.SendTelegramWith400Words(c, respList)
}
