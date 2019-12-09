package obol

import (
	"net/http"
)

type ObolService interface {
	// TODO(Oedipus): Change to return map instead of array
	GetCoinRates(obolURL string, coin string) ([]Rate, error)
	GetCoin2CoinRates(obolURL string, fromCoin string, toCoin string) (float64, error)
	GetCoin2CoinRatesWithAmount(obolURL string, fromCoin string, toCoin string, amount string) (CoinToCoinWithAmountResponse, error)
	GetProductionURL() (string, error)
	GetHTTPClient() (http.Client, error)
}
