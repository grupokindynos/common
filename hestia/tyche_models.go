package hestia

type ShiftStatus int

const (
	ShiftStatusPending ShiftStatus = iota
	ShiftStatusConfirming
	ShiftStatusConfirmed
	ShiftStatusError
	ShiftStatusRefund
	ShiftStatusRefunded
	ShiftStatusComplete
)

var (
	ShiftStatusStr = map[ShiftStatus]string{
		0: "PENDING",
		1: "CONFIRMING",
		2: "CONFIRMED",
		3: "ERROR",
		4: "REFUND",
		5: "REFUNDED",
		6: "COMPLETE",
	}
)

func GetShiftStatusString(status ShiftStatus) string {
	value, ok := ShiftStatusStr[status]
	if !ok {
		return ""
	}
	return value
}

type ShiftStatusV2 int

const (
	ShiftStatusV2Created ShiftStatusV2 = iota
	ShiftStatusV2Confirmed
	ShiftStatusV2ProcessingOrders
	ShiftStatusV2PartiallyCompleted
	ShiftStatusV2Complete
	ShiftStatusV2SentToUser
	ShiftStatusV2Error
	ShiftStatusV2Refund
	ShiftStatusV2Refunded
)

var (
	ShiftStatusV2Str = map[ShiftStatusV2]string{
		0: "CREATED",
		1: "CONFIRMED",
		2: "PROCESSING_ORDERS",
		3: "PARTIALLY_COMPLETED", // One of the two orders completed succesfully
		4: "COMPLETE",            //
		5: "SENT_TO_USER",
		6: "ERROR",
		7: "REFUNDED",
	}
)

func GetShiftStatusv2String(status ShiftStatusV2) string {
	value, ok := ShiftStatusV2Str[status]
	if !ok {
		return ""
	}
	return value
}

type ShiftV2TradeStatus int

const (
	ShiftV2TradeStatusInitialized ShiftV2TradeStatus = iota
	ShiftV2TradeStatusCreated
	ShiftV2TradeStatusPerforming
	ShiftV2TradeStatusCompleted
	ShiftV2TradeStatusWithdrawn         // just for outbound trade
	ShiftV2TradeStatusUserDeposit       // just for outbound trade
	ShiftV2TradeStatusWithdrawCompleted // just for outbound trade
)

var (
	ShiftV2TradeStatusStr = map[ShiftV2TradeStatus]string{
		0: "INITIALIZED",
		1: "CREATED",
		2: "PERFORMING",
		3: "COMPLETED",
		4: "WITHDRAWN",
		5: "USER_DEPOSIT",
		6: "WITHDRAW_COMPLETED",
	}
)

func GetShiftV2TradeStatusString(status ShiftV2TradeStatus) string {
	value, ok := ShiftV2TradeStatusStr[status]
	if !ok {
		return ""
	}
	return value
}
