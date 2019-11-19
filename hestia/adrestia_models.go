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
}

type AdrestiaOrderUpdate struct {
	ID     string `firestore:"id" json:"id"`
	Status string `firestore:"status" json:"status"`
}

type AdrestiaStatus int

const (
	AdrestiaStatusSentAmount AdrestiaStatus = iota
	AdrestiaStatusCreated
	AdrestiaStatusPartiallyFulfilled
	AdrestiaStatusPendingWidthdrawal
	AdrestiaStatusCompleted
	AdrestiaStatusError
)

var (
	AdrestiaStatusStr = map[AdrestiaStatus]string{
		0: "SENTAMOUNT",
		1: "CREATED",
		2: "PARTIALLYFULFILLED",
		3: "PENDINGWITHDRAWAL",
		4: "COMPLETED",
		5: "ERROR",
	}
)

func GetAdrestiaStatusString(status AdrestiaStatus) string {
	value, ok := AdrestiaStatusStr[status]
	if !ok {
		return ""
	}
	return value
}
