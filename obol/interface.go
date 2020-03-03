package obol

type ObolService interface {
	GetCoinRatesV2(coin string) (map[string]RateV2, error)
	GetCoin2FIATRate(fromCoin string, toCoin string) (float64, error)
	GetCoinRates(coin string) ([]Rate, error)
	GetCoin2CoinRates(fromCoin string, toCoin string) (float64, error)
	GetCoin2CoinRatesWithAmount(fromCoin string, toCoin string, amount string) (CoinToCoinWithAmountResponse, error)
}
