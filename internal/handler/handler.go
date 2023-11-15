package handler

import (
	"github.com/adshao/go-binance/v2"

	"github.com/blanergol/crypto_rate_bot/config"
	"github.com/blanergol/crypto_rate_bot/internal/usecase"
)

type Handlers struct {
	Pricer
	Tokener

	bc      *binance.Client
	usecase *usecase.UseCasesIml
	cfg     *config.Config
}

func NewHandlers(bc *binance.Client, useCase *usecase.UseCasesIml, cfg *config.Config) Handlers {
	return Handlers{
		bc:      bc,
		usecase: useCase,
		cfg:     cfg,
	}
}
