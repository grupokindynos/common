package herodotus

import "github.com/grupokindynos/common/hestia"

type VoucherV2Filters struct {
	Status map[hestia.VoucherStatusV2]bool `json:"status"`
	FromTimestamp int64 `json:"from_timestamp"`
	ToTimestamp int64 `json:"to_timestamp"`
	UserId map[string]bool `json:"user_id"`
	VoucherId map[int]bool `json:"voucher_id"`
	ProviderId map[string]bool `json:"provider_id"`
	Coin map[string]bool `json:"coin"`
}
