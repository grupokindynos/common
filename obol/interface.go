package obol

type ObolService interface {
	// TODO(Oedipus): Change to return map instead of array
	GetCoinRates(coin string) ([]Rate, error)
	GetCoin2CoinRates(fromCoin string, toCoin string) (float64, error)
	GetCoin2CoinRatesWithAmount(fromCoin string, toCoin string, amount string) (CoinToCoinWithAmountResponse, error)
}
