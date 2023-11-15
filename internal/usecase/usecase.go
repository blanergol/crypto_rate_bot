package usecase

import (
	"github.com/adshao/go-binance/v2"
)

var _ UseCases = (*UseCasesIml)(nil)

type UseCases interface {
	Tokens
	Price
}

type UseCasesIml struct {
	UseCases

	bc *binance.Client
}

func New(bc *binance.Client) *UseCasesIml {
	return &UseCasesIml{bc: bc}
}
