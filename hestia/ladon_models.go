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
	VoucherStatusRefundedPartially
	VoucherStatusRefunded
)

var (
	VoucherStatusStr = map[VoucherStatus]string{
		0: "PENDING",
		1: "CONFIRMED",
		2: "CONFIRMING",
		3: "AWAITING PROVIDER",
		4: "ERROR",
		5: "COMPLETE",
		6: "REFUND_TOTAL",
		7: "REFUND_FEE",
		8: "REFUNDED",
		9: "REFUNDED_PARTIALLY",
	}
)

func GetVoucherStatusString(status VoucherStatus) string {
	value, ok := VoucherStatusStr[status]
	if !ok {
		return ""
	}
	return value
}
