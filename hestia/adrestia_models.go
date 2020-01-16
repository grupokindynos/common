package hestia

type AdrestiaOrder struct {
	ID           string         `firestore:"id" json:"id"`
	DualExchange bool           `firestore:"dual_exchange" json:"dual_exchange"`
	Time         int64          `firestore:"time" json:"time"`
	Status       AdrestiaStatus `firestore:"status" json:"status"`
	Amount       float64        `firestore:"amount" json:"amount"`
	BtcRate      float64        `firestore:"btc_rate" json:"btc_rate"`
	FromCoin     string         `firestore:"from_coin" json:"from_coin"`
	ToCoin       string         `firestore:"to_coin" json:"to_coin"`

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
	OrderId           string  `firestore:"order_id" json:"order_id"`
	Symbol            string  `firestore:"symbol" json:"symbol"`
	Side              string  `firestore:"side" json:"side"`
	Amount            float64 `firestore:"amount" json:"amount"`
	ListingAmount     float64 `firestore:"listingAmount" json:"listingAmount"`
	Timestamp         int64   `firestore:"time" json:"time"`
	Exchange          string  `firestore:"exchange" json:"exchange"`
	ListingCurrency   string  `firestore:"listingCurrency" json:"listingCurrency"`
	ReferenceCurrency string  `firestore:"referenceCurrency" json:"referenceCurrency"`
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
