package usecase

import (
	"context"

	"github.com/adshao/go-binance/v2"

	"github.com/blanergol/crypto_rate_bot/internal/models"
)

var _ Price = (*UseCasesIml)(nil)

type Price interface {
	GetCurrentPriceForListTokens(ctx context.Context, tokens []models.Token) (*[]models.CurrentPrice, error)
	GetPriceForListTokens(ctx context.Context, tokens []models.Token, time string) (*[]models.Price, error)
}

func (u UseCasesIml) GetCurrentPriceForListTokens(ctx context.Context, tokens []models.Token) (*[]models.CurrentPrice, error) {
	var currentPrices []models.CurrentPrice
	symbols := models.MakeTokenSymbols(tokens)

	binanceCurrentPrices, err := u.bc.NewListPricesService().Symbols(symbols).Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, price := range binanceCurrentPrices {
		currentPrices = append(currentPrices, models.MakeCurrentPrice(price))
	}

	return &currentPrices, nil
}

func (u UseCasesIml) GetPriceForListTokens(ctx context.Context, tokens []models.Token, time string) (*[]models.Price, error) {
	var prices []models.Price
	var allBinancePrices []*binance.SymbolTicker

	symbols := models.MakeTokensChunkSlice(tokens, 100)

	for _, s := range symbols {
		binancePrices, err := u.bc.NewListSymbolTickerService().Symbols(models.MakeTokenSymbols(s)).WindowSize(time).Do(ctx)
		if err != nil {
			return nil, err
		}
		allBinancePrices = append(allBinancePrices, binancePrices...)
	}

	for _, price := range allBinancePrices {
		prices = append(prices, models.MakePrice(price))
	}

	return &prices, nil
}
