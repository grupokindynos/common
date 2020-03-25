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

func GetShiftStatusV2String(status ShiftStatus) string {
	value, ok := ShiftStatusStr[status]
	if !ok {
		return ""
	}
	return value
}

type ShiftStatusv2 int

const (
	ShiftStatusV2Pending ShiftStatusv2 = iota
	ShiftStatusV2Confirming
	ShiftStatusV2Confirmed
	ShiftStatusV2Error
	ShiftStatusV2Refund
	ShiftStatusV2Refunded
	ShiftStatusV2Complete
)

var (
	ShiftStatusV2Str = map[ShiftStatus]string{
		0: "CREATED",
		1: "CONFIRMED",
		2: "PROCESSING_ORDERS",
		3: "PARTIALLY_COMPLETED", // One of the two orders completed succesfully
		4: "COMPLETED",           //
		5: "SENT_TO_USER",
	}
)

func GetShiftStatusString(status ShiftStatus) string {
	value, ok := ShiftStatusStr[status]
	if !ok {
		return ""
	}
	return value
}
