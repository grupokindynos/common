package obol

type ObolService interface {
	GetCoinRates(obolURL string, coin string) ([]Rate, error)
	GetCoin2CoinRates(obolURL string, fromCoin string, toCoin string) (float64, error)
	GetCoin2CoinRatesWithAmount(obolURL string, fromCoin string, toCoin string, amount string) (CoinToCoinWithAmountResponse, error)
}
