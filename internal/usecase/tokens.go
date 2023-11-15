package usecase

import (
	"context"

	"github.com/blanergol/crypto_rate_bot/internal/models"
)

var _ Tokens = (*UseCasesIml)(nil)

type Tokens interface {
	GetListTokensWithCurrentPrice(ctx context.Context, token models.TokenName) (*[]models.Token, error)
}

func (u UseCasesIml) GetListTokensWithCurrentPrice(ctx context.Context, token models.TokenName) (*[]models.Token, error) {
	var listTokens []models.Token

	symbols, err := u.bc.NewExchangeInfoService().Symbols().Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, symbol := range symbols.Symbols {
		if symbol.QuoteAsset == token.String() {
			listTokens = append(listTokens, models.MakeToken(symbol, "0"))
		}
	}

	prices, err := u.GetCurrentPriceForListTokens(ctx, listTokens)
	if err != nil {
		return &listTokens, err
	}

	pricesMap := models.MakeCurrentPriceMap(*prices)
	for i, tk := range listTokens {
		listTokens[i].Price = pricesMap[tk.Symbols].Price
	}
	return &listTokens, nil
}
