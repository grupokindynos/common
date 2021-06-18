package explorer


/*
	BSC Scan Models
 */
type BSCResponse struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Result  string `json:"result"`
}

/*
	BSC API Models hosted at https://pp-bsc-api.herokuapp.com/
 */
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

type BSCAddressResponse struct {
	BSCBase
	Data struct{
		Address string `json:"address"`
	} `json:"data"`
}

type BSCBalanceResponse struct {
	BSCBase
	Data struct{
		Balance float64 `json:"balance"`
	} `json:"data"`
}

type BSCSwapInfo struct {
	CoinFrom string `json:"coin_from"`
	AmountIn int64 `json:"amount_in"`
	AmountOut int64 `json:"amount_out"`
}

type BSCSwapResponse struct {
	BSCBase
	Data struct{
		Path []string `json:"path"`
		TxID string `json:"tx_id"`
	} `json:"data"`
}

type BSCWithdrawResponse struct {
	BSCBase
	Data struct{
		TxID string `json:"tx_id"`
	} `json:"data"`
}

type BSCWithdrawInput struct {
	Address string `json:"address"`
	Asset string `json:"asset"`
	Amount float64 `json:"amount"`
}