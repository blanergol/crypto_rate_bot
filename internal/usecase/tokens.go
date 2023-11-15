package usecase

import (
	"context"

	"github.com/blanergol/crypto_rate_bot/internal/models"
)

var _ Tokens = (*UseCasesIml)(nil)

type Tokens interface {
	GetListTokensWithCurrentPrice(ctx context.Context, tokensFilter []string) (*[]models.Token, error)
}

func (u UseCasesIml) GetListTokensWithCurrentPrice(ctx context.Context, tokensFilter []string) (*[]models.Token, error) {
	var listTokens []models.Token
	var symbolsFilter []string

	predict := u.bc.NewExchangeInfoService()
	if len(tokensFilter) > 0 {
		for _, f := range tokensFilter {
			symbolsFilter = append(symbolsFilter, f+models.USDT.String())
		}
		predict.Symbols(symbolsFilter...)
	}
	symbols, err := predict.Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, symbol := range symbols.Symbols {
		if symbol.QuoteAsset == models.USDT.String() {
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
