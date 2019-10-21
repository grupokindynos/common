package hestia

import "errors"

type ShiftStatus int

const (
	ShiftStatusPending ShiftStatus = iota
	ShiftStatusConfirming
	ShiftStatusConfirmed
	ShiftStatusError
	ShiftStatusRefunded
	ShiftStatusComplete
)

var (
	ShiftStatusStr = map[int]string{
		0: "PENDING",
		1: "CONFIRMING",
		2: "CONFIRMED",
		3: "ERROR",
		4: "REFUNDED",
		5: "COMPLETE",
	}
)

func GetShiftStatusString(status int) (string, error) {
	value, ok := ShiftStatusStr[status]
	if !ok {
		return "", errors.New("unknown status")
	}
	return value, nil
}
