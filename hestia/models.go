package hestia

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

