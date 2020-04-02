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
