package hestia

type VoucherStatus int

const (
	VoucherStatusPending VoucherStatus = iota
	VoucherStatusConfirmed
	VoucherStatusConfirming
	VoucherStatusAwaitingProvider
	VoucherStatusError
	VoucherStatusComplete
	VoucherStatusRefunded
)

var (
	VoucherStatusStr = map[int]string{
		0: "PENDING",
		1: "CONFIRMED",
		2: "CONFIRMING",
		3: "AWAITING PROVIDER",
		4: "ERROR",
		5: "COMPLETE",
		6: "REFUNDED",
	}
)

func GetVoucherStatusString(status int) string {
	value, ok := VoucherStatusStr[status]
	if !ok {
		return ""
	}
	return value
}
