package obol

import "github.com/shopspring/decimal"

type Rate struct {
	Code string  `json:"code"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}

type RateV2 struct {
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

// CoinToCoinWithAmountResponse is the response of a CoinToCoinWithAmount result
type CoinToCoinWithAmountResponse struct {
	AveragePrice decimal.Decimal `json:"average_price"`
	Amount       decimal.Decimal `json:"amount"`
}
