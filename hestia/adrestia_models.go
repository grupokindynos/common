package hestia

type AdrestiaOrder struct {
	ID            string         `firestore:"id" json:"id"`
	DualExchange  bool           `firestore:"dual_exchange" json:"dual_exchange"`
	CreatedTime   int64          `firestore:"createdTime" json:"createdTime"`
	FulfilledTime int64          `firestore:"fulfilledTime" json:"fulfilledTime"`
	Status        AdrestiaStatus `firestore:"status" json:"status"`
	Amount        float64        `firestore:"amount" json:"amount"`
	BtcRate       float64        `firestore:"btc_rate" json:"btc_rate"`
	FromCoin      string         `firestore:"from_coin" json:"from_coin"`
	ToCoin        string         `firestore:"to_coin" json:"to_coin"`

	Message string `firestore:"message" json:"message"`

	FirstOrder ExchangeOrder `firestore:"first_order" json:"first_order"`
	FinalOrder ExchangeOrder `firestore:"final_order" json:"final_order"`

	HETxId          string `firestore:"he_tx_id" json:"he_tx_id"`
	EETxId          string `firestore:"ee_tx_id" json:"ee_tx_id"`
	EHTxId          string `firestore:"eh_tx_id" json:"eh_tx_id"`
	FirstExAddress  string `firestore:"f_ex_address" json:"f_ex_address"`
	SecondExAddress string `firestore:"s_ex_address" json:"s_ex_address"`
	WithdrawAddress string `firestore:"withdraw_address" json:"withdraw_address"`
}

type ExchangeOrder struct {
	OrderId          string  `firestore:"order_id" json:"order_id"`
	Symbol           string  `firestore:"symbol" json:"symbol"`
	Side             string  `firestore:"side" json:"side"`
	Amount           float64 `firestore:"amount" json:"amount"`
	ReceivedAmount   float64 `firestore:"receivedAmount" json:"receivedAmount"`
	CreatedTime      int64   `firestore:"createdTime" json:"createdTime"`
	FulfilledTime    int64   `firestore:"fulfilledTime" json:"fulfilledTime"`
	Exchange         string  `firestore:"exchange" json:"exchange"`
	ReceivedCurrency string  `firestore:"receivedCurrency" json:"receivedCurrency"`
	SoldCurrency     string  `firestore:"soldCurrency" json:"soldCurrency"`
}

type OrderStatus struct {
	Status          ExchangeStatus
	AvailableAmount float64
}

type AdrestiaOrderUpdate struct {
	ID     string         `firestore:"id" json:"id"`
	Status AdrestiaStatus `firestore:"status" json:"status"`
}

type AdrestiaStatus int
type ExchangeStatus int

const (
	AdrestiaStatusCreated AdrestiaStatus = iota
	AdrestiaStatusFirstExchange
	AdrestiaStatusFirstConversion
	AdrestiaStatusSecondExchange
	AdrestiaStatusSecondConversion
	AdrestiaStatusFirstWithdrawal
	AdrestiaStatusSecondWithdrawal
	AdrestiaStatusExchangeComplete
	AdrestiaStatusCompleted
	AdrestiaStatusError
)

var (
	AdrestiaStatusStr = map[AdrestiaStatus]string{
		0: "CREATED",
		1: "FIRST_EXCHANGE",
		2: "FIRST_CONVERSION",
		3: "FIRST_WITHDRAWAL",
		4: "SECOND_EXCHANGE",
		5: "SECOND_CONVERSION",
		6: "SECOND_WITHDRAWAL",
		7: "EXCHANGE_COMPLETE",
		8: "COMPLETED",
		9: "ERROR",
	}
)

const (
	ExchangeStatusOpen ExchangeStatus = iota
	ExchangeStatusCompleted
	ExchangeStatusError
)

var (
	ExchangeStatusStr = map[ExchangeStatus]string{
		0: "OPEN",
		1: "COMPLETED",
		2: "ERROR",
	}
)

// Returns listing and reference currencies
func (eo *ExchangeOrder) GetTradingPair() (string, string) {
	if eo.Side == "buy" {
		return eo.ReceivedCurrency, eo.SoldCurrency
	} else {
		return eo.SoldCurrency, eo.ReceivedCurrency
	}
}

func GetAdrestiaStatusString(status AdrestiaStatus) string {
	value, ok := AdrestiaStatusStr[status]
	if !ok {
		return ""
	}
	return value
}

func GetExchangeStatusString(status ExchangeStatus) string {
	value, ok := ExchangeStatusStr[status]
	if !ok {
		return ""
	}
	return value
}
