package models

import "github.com/adshao/go-binance/v2"

type TokenName string

const (
	USDT TokenName = "USDT"
	BTC  TokenName = "BTC"
	ETC  TokenName = "ETC"
	BNB  TokenName = "BNB"
	TON  TokenName = "TON"
)

type Token struct {
	Name    string
	Symbol  string
	Symbols string
	Price   string
}

func MakeToken(token binance.Symbol, price string) Token {
	return Token{
		Name:    token.BaseAsset,
		Symbol:  token.QuoteAsset,
		Symbols: token.Symbol,
		Price:   price,
	}
}

func MakeTokenSymbols(tokens []Token) []string {
	var symbols []string
	for _, token := range tokens {
		symbols = append(symbols, token.Symbols)
	}
	return symbols
}

func MakeTokensMap(tokens []Token) map[string]Token {
	tokensMap := map[string]Token{}
	for _, token := range tokens {
		tokensMap[token.Symbols] = token
	}
	return tokensMap
}

func MakeTokensChunkSlice(tokens []Token, chunkSize int) [][]Token {
	var chunks [][]Token
	for i := 0; i < len(tokens); i += chunkSize {
		end := i + chunkSize
		if end > len(tokens) {
			end = len(tokens)
		}

		chunks = append(chunks, tokens[i:end])
	}

	return chunks
}

func (t TokenName) String() string {
	return string(t)
}
