package herodotus

import "github.com/grupokindynos/common/hestia"

type VoucherV2Filters struct {
	Status hestia.VoucherStatusV2 `json:"status"`
	FromTimestamp int64 `json:"from_timestamp"`
	ToTimestamp int64 `json:"to_timestamp"`
	UserId string `json:"user_id"`
	VoucherId int `json:"voucher_id"`
	ProviderId string `json:"provider_id"`
	Coin string `json:"coin"`
}
