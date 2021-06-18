package explorer

type BSCResponse struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Result  string `json:"result"`
}

type BSCBase struct {
	Status int32 `json:"status"`
	Error string `json:"error"`
}

type BSCTxInfoResponse struct {
	BSCBase
	Data struct{
		TxInfo struct {
			Block int64 `json:"block"`
			Confirmations int64 `json:"confirmations"`
			IsServiceTx bool `json:"service_tx"`
			TransactionHash string `json:"tx_hash"`
			ReceivedAmount float64 `json:"received_amount"`
		}`json:"tx_info"`
	} `json:"data"`
}