package hestia

// Represents either a withdrawal from a exchange to our hot wallet
// or a deposit from our hot wallet to a exchange
type SimpleTx struct {
	Id string `firestore:"id" json:"id"`
	TxId string `firestore:"txid" json:"txid"`
	BalancerId string `firestore:"balancer_id" json:"balancer_id"`
	Exchange string `firestore:"exchange" json:"exchange"`
	Address string `firestore:"address" json:"address"`
	Currency string `firestore:"currency" json:"currency"`
	Amount float64 `firestore:"amount" json:"amount"`
	Status SimpleTxStatus `firestore:"status" json:"status"`
	Timestamp int64 `firestore:"timestamp" json:"timestamp"`
}

type Balancer struct {
	Id string `firestore:"id" json:"id"`
	Status BalancerStatus `firestore:"status" json:"status"`
	CreatedTime int64 `firestore:"created_time" json:"created_time"`
	FulfilledTime int64 `firestore:"fulfilled_time" json:"fulfilled_time"`
}

type Trade struct {
	Id string `firestore:"id" json:"id"`
	OrderId string `firestore:"order_id" json:"order_id"`
	Amount float64 `firestore:"amount" json:"amount"`
	ReceivedAmount float64 `firestore:"received_amount" json:"received_amount"`
	FromCoin string `firestore:"from_coin" json:"from_coin"`
	ToCoin string `firestore:"to_coin" json:"to_coin"`
	Symbol string `firestore:"symbol" json:"symbol"`
	Side string `firestore:"side" json:"side"`
	CreatedTime int64 `firestore:"created_time" json:"created_time"`
	FulfilledTime int64 `firestore:"fulfilled_time" json:"fulfilled_time"`
}

type BalancerOrder struct {
	Id string `firestore:"id" json:"id"`
	BalancerId string `firestore:"balancer_id" json:"balancer_id"`
	FromCoin string `firestore:"from_coin" json:"from_coin"`
	ToCoin string `firestore:"to_coin" json:"to_coin"`
	DualConversion bool `firestore:"dual_conversion" json:"dual_conversion"`
	OriginalAmount float64 `firestore:"original_amount" json:"original_amount"`
	FirstTrade Trade `firestore:"first_trade" json:"first_trade"`
	SecondTrade Trade `firestore:"second_trade" json:"second_trade"`
	Status BalancerOrderStatus `firestore:"status" json:"status"`
	ReceivedAmount float64 `firestore:"received_amount" json:"received_amount"`
	DepositTxId string `firestore:"deposit_txid" json:"deposit_txid"`
	WithdrawTxId string `firestore:"withdraw_txid" json:"withdraw_txid"`
	DepositAddress string `firestore:"deposit_address" json:"deposit_address"`
	WithdrawAddress string `firestore:"withdraw_address" json:"withdraw_address"`
	CreatedTime int64 `firestore:"created_time" json:"created_time"`
	FulfilledTime int64 `firestore:"fulfilled_time" json:"fulfilled_time"`
}

type SimpleTxStatus int // For deposits or withdrawals to exchanges
type BalancerStatus int
type BalancerOrderStatus int


const (
	SimpleTxStatusCreated SimpleTxStatus = iota
	SimpleTxStatusPerformed
	SimpleTxStatusCompleted
	SimpleTxStatusPlutusDeposit
)

var (
	SimpleTxStatusStr = map[SimpleTxStatus]string {
		0: "CREATED",
		1: "PERFORMED",
		2: "COMPLETED",
		3: "PLUTUS_DEPOSIT", // Status only for withdrawals
	}
)

const (
	BalancerStatusCreated BalancerStatus = iota
	BalancerStatusWithdrawal
	BalancerStatusTradeOrders
	BalancerStatusCompleted
)

var (
	BalancerStatusStr = map[BalancerStatus]string {
		0: "CREATED",
		1: "WITHDRAWAL",
		2: "TRADE_ORDERS",
		3: "COMPLETED",
	}
)

const (
	BalancerOrderStatusCreated BalancerOrderStatus = iota
	BalancerOrderStatusExchangeDepositSent
	BalancerOrderStatusFirstTrade
	BalancerOrderStatusSecondTrade
	BalancerOrderStatusWithdrawal
	BalancerOrderStatusPlutusDeposit
	BalancerOrderStatusCompleted
)

var (
	BalancerOrderStatusStr = map[BalancerOrderStatus]string {
		0: "CREATED",
		1: "EXCHANGE_DEPOSIT_SENT",
		2: "FIRST_TRADE",
		3: "SECOND_TRADE",
		4: "WITHDRAWAL",
		5: "PLUTUS_DEPOSIT",
		6: "COMPLETED",
	}
)

func GetSimpleTxStatusString(status SimpleTxStatus) string {
	value, ok := SimpleTxStatusStr[status]
	if !ok {
		return ""
	}
	return value
}

func GetBalancerStatusString(status BalancerStatus) string {
	value, ok := BalancerStatusStr[status]
	if !ok {
		return ""
	}
	return value
}

func GetBalancerOrderStatusString(status BalancerOrderStatus) string {
	value, ok := BalancerOrderStatusStr[status]
	if !ok {
		return ""
	}
	return value
}
