package hestia

type AdrestiaOrder struct {
	ID              string  `firestore:"id" json:"id"`
	Exchange        string  `firestore:"exchange" json:"exchange"`
	Time            int64   `firestore:"time" json:"time"`
	Status          string  `firestore:"status" json:"status"`
	Amount          float64 `firestore:"amount" json:"amount"`
	FromCoin        string  `firestore:"from_coin" json:"from_coin"`
	ToCoin          string  `firestore:"to_coin" json:"to_coin"`
	WithdrawAddress string  `firestore:"withdraw_address" json:"withdraw_address"`
	Message         string  `firestore:"message" json:"message"`
	OrderId         string  `firestore:"order_id" json:"order_id"`
	TxId            string  `firestore:"tx_id" json:"tx_id"`
	ExchangeAddress string  `firestore:"exchange_address" json:"exchange_address"`
}

type AdrestiaOrderUpdate struct {
	ID     string `firestore:"id" json:"id"`
	Status string `firestore:"status" json:"status"`
}

type AdrestiaStatus int

const (
	AdrestiaStatusCreated AdrestiaStatus = iota
	AdrestiaStatusFirstExchange
	AdrestiaStatusFirstConversion
	AdrestiaStatusSecondExchange
	AdrestiaStatusSecondConversion
	AdrestiaStatusExchangeComplete
	AdrestiaStatusCompleted
	AdrestiaStatusError
)

var (
	AdrestiaStatusStr = map[AdrestiaStatus]string{
		0: "CREATED",
		1: "FIRST_EXCHANGE",
		2: "FIRST_CONVERSION",
		3: "SECOND_EXCHANGE",
		4: "SECOND_CONVERSION",
		5: "EXCHANGE_COMPLETE",
		6: "COMPLETED",
		7: "ERROR",
	}
)

func GetAdrestiaStatusString(status AdrestiaStatus) string {
	value, ok := AdrestiaStatusStr[status]
	if !ok {
		return ""
	}
	return value
}
