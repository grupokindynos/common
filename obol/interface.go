package obol

type ObolService interface {
	GetCoinRates(obolURL string, coin string) ([]Rate, error)
	GetCoin2Rates(obolURL string, fromCoin string, toCoin string) (float64, error)
	GetCoin2RatesWithAmount(obolURL string, fromCoin string, toCoin string, amount string) (CoinToCoinWithAmountResponse, error)
}
