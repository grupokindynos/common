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
