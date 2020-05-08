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


type VoucherStatusV2 int

const (
	VoucherStatusV2PaymentProcessing VoucherStatusV2 = iota
	VoucherStatusV2Redeemed
	VoucherStatusV2PerformingTrade
	VoucherStatusV2Complete
	VoucherStatusV2Refunded
	VoucherStatusV2Error
	VoucherStatusV2NeedsRefund
	VoucherStatusV2WaitingRefundTxId
)

var(
	VoucherStatusV2Str = map[VoucherStatusV2]string{
		0: "PAYMENT_PROCESSING",
		1: "REDEEMED",
		2: "PERFORMING_TRADE",
		3: "COMPLETE",
		4: "REFUNDED",
		5: "ERROR",
		6: "NEEDS_REFUND",
		7: "WAITING_REFUND_TXID",
	}
)

func GetVoucherStatusV2String(status VoucherStatusV2) string {
	value, ok := VoucherStatusV2Str[status]
	if !ok {
		return ""
	}
	return value
}