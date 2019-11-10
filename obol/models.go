package obol

type Rate struct {
	Code string  `json:"code"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

// CoinToCoinWithAmountResponse is the response of a CoinToCoinWithAmount result
type CoinToCoinWithAmountResponse struct {
	Rates        map[float64]float64 `json:"rates"`
	AveragePrice float64             `json:"average_price"`
}