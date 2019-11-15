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
	VoucherStatusRefundTotal
	VoucherStatusRefundFee
)

var (
	VoucherStatusStr = map[VoucherStatus]string{
		0: "PENDING",
		1: "CONFIRMED",
		2: "CONFIRMING",
		3: "AWAITING PROVIDER",
		4: "ERROR",
		5: "COMPLETE",
		6: "REFUNDED",
		7: "REFUND_TOTAL",
		8: "REFUND_FEE",
	}
)

func GetVoucherStatusString(status VoucherStatus) string {
	value, ok := VoucherStatusStr[status]
	if !ok {
		return ""
	}
	return value
}
