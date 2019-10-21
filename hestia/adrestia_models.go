package hestia

import "errors"

type AdrestiaOrder struct {
	ID              string  `firestore:"id" json:"id"`
	UID             string  `firestore:"uid" json:"uid"`
	Exchange        string  `firestore:"exchange" json:"exchange"`
	Time            int64   `firestore:"time" json:"time"`
	Status          string  `firestore:"status" json:"status"`
	Amount          float64 `firestore:"amount" json:"amount"`
	FromCoin        string  `firestore:"from_coin" json:"from_coin"`
	ToCoin          string  `firestore:"to_coin" json:"to_coin"`
	WithdrawAddress string  `firestore:"withdraw_address" json:"withdraw_address"`
	Message         string  `firestore:"message" json:"message"`
}

type AdrestiaStatus int

const (
	AdrestiaStatusPartiallyFulfilled AdrestiaStatus = iota
	AdrestiaStatusCreated
	AdrestiaStatusPendingWidthdrawal
	AdrestiaStatusCompleted
	AdrestiaStatusError
)

var (
	AdrestiaStatusStr = map[int]string{
		0: "PARTIALLYFULFILLED",
		1: "CREATED",
		2: "PENDINGWITHDRAWAL",
		3: "COMPLETED",
		4: "ERROR",
	}
)

func GetAdrestiaStatusString(status int) (string, error) {
	value, ok := AdrestiaStatusStr[status]
	if !ok {
		return "", errors.New("unknown status")
	}
	return value, nil
}
