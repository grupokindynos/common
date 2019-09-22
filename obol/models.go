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
