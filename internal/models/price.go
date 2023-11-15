package models

import (
	"github.com/adshao/go-binance/v2"
)

type CurrentPrice struct {
	Symbol string
	Price  string
}

type Price struct {
	Symbol             string
	PriceChange        string
	PriceChangePercent string
}

func MakeCurrentPrice(price *binance.SymbolPrice) CurrentPrice {
	return CurrentPrice{
		Symbol: price.Symbol,
		Price:  price.Price,
	}
}

func MakePrice(price *binance.SymbolTicker) Price {
	return Price{
		Symbol:             price.Symbol,
		PriceChange:        price.PriceChange,
		PriceChangePercent: price.PriceChangePercent,
	}
}

func MakeCurrentPriceMap(prices []CurrentPrice) map[string]CurrentPrice {
	priceMap := map[string]CurrentPrice{}
	for _, price := range prices {
		priceMap[price.Symbol] = price
	}
	return priceMap
}
