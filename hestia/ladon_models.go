package hestia

type VoucherStatus int

const (
	VoucherStatusPending VoucherStatus = iota
	VoucherStatusConfirmed
	VoucherStatusConfirming
	VoucherStatusAwaitingProvider
	VoucherStatusError
	VoucherStatusComplete
	VoucherStatusRefundTotal
	VoucherStatusRefundFee
	VoucherStatusRefunded
	VoucherStatusRefundedPartially
	VoucherStatusBitcouRefunded
)

var (
	VoucherStatusStr = map[VoucherStatus]string{
		0:  "PENDING",
		1:  "CONFIRMED",
		2:  "CONFIRMING",
		3:  "AWAITING_PROVIDER",
		4:  "ERROR",
		5:  "COMPLETE",
		6:  "REFUND_TOTAL",
		7:  "REFUND_FEE",
		8:  "REFUNDED",
		9:  "REFUNDED_PARTIALLY",
		10: "BITCOU_REFUNDED",
	}
)

func GetVoucherStatusString(status VoucherStatus) string {
	value, ok := VoucherStatusStr[status]
	if !ok {
		return ""
	}
	return value
}
