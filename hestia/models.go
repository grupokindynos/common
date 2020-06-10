package hestia

type Card struct {
	Address    string `firestore:"address" json:"address"`
	CardCode   string `firestore:"card_code" json:"card_code"`
	CardNumber string `firestore:"card_number" json:"card_number"`
	City       string `firestore:"city" json:"city"`
	Email      string `firestore:"email" json:"email"`
	FirstName  string `firestore:"firstname" json:"firstname"`
	LastName   string `firestore:"lastname" json:"lastname"`
	UID        string `firestore:"uid" json:"uid"`
}

type Pin struct {
	CardCode string `firestore:"card_code" json:"card_code"`
	PinCode  string `firestore:"pin_code" json:"pin_code"`
}

type Coin struct {
	Ticker   string        `firestore:"ticker" json:"ticker"`
	Shift    Properties    `firestore:"shift" json:"shift"`
	Deposits Properties    `firestore:"deposits" json:"deposits"`
	Vouchers Properties    `firestore:"vouchers" json:"vouchers"`
	Orders   Properties    `firestore:"orders" json:"orders"`
	Balances BalanceLimits `firestore:"balances" json:"balances"`
	Adrestia AdrestiaInfo  `firestore:"adrestia" json:"adrestia"`
}

type AdrestiaInfo struct {
	Available bool    `firestore:"available" json:"available"`
	CoinUsage float64 `firestore:"coin_usage" json:"coin_usage"`
}

type CoinBalances struct {
	Ticker  string  `firestore:"ticker" json:"ticker"`
	Balance float64 `firestore:"balance" json:"balance"`
	Status  string  `firestore:"status" json:"status"`
}

type BalanceLimits struct {
	HotWallet float64 `firestore:"hot_wallet" json:"hot_wallet"`
	Exchanges float64 `firestore:"exchanges" json:"exchanges"`
}

type Payment struct {
	Address       string `firestore:"address" json:"address"`
	Amount        int64  `firestore:"amount" json:"amount"`
	Coin          string `firestore:"coin" json:"coin"`
	Txid          string `firestore:"txid" json:"txid"`
	Confirmations int32  `firestore:"confirmations" json:"confirmations"`
}

type PaymentWithFee struct {
	Address       string `firestore:"address" json:"address"`
	Amount        int64  `firestore:"amount" json:"amount"`
	Fee           int64  `firestore:"amount" json:"fee"`
	Usable        int64  `firestore:"amount" json:"usable"`
	Coin          string `firestore:"coin" json:"coin"`
	Txid          string `firestore:"txid" json:"txid"`
	Confirmations int32  `firestore:"confirmations" json:"confirmations"`
}

type LightPayment struct {
	Address string `firestore:"address" json:"address"`
	Coin    string `firestore:"coin" json:"coin"`
	Txid    string `firestore:"txid" json:"txid"`
	Amount  int64  `firestore:"amount" json:"amount"`
}

type Properties struct {
	FeePercentage float64 `firestore:"fee_percentage" json:"fee_percentage"`
	Available     bool    `firestore:"available" json:"available"`
}

type Config struct {
	Shift    Available `firestore:"shift" json:"shift"`
	Deposits Available `firestore:"deposits" json:"deposits"`
	Vouchers Available `firestore:"vouchers" json:"vouchers"`
	Orders   Available `firestore:"orders" json:"orders"`
	Adrestia Available `firestore:"adrestia" json:"adrestia"`
}

type Available struct {
	Service   bool `firestore:"service" json:"service"`
	Processor bool `firestore:"processor" json:"processor"`
}

type Order struct {
	ID                  string                     `firestore:"id" json:"id"`
	UID                 string                     `firestore:"uid" json:"uid"`
	Status              string                     `firestore:"status" json:"status"`
	PaymentInfo         Payment                    `firestore:"payment_info" json:"payment_info"`
	FeePayment          Payment                    `firestore:"fee_payment" json:"fee_payment"`
	AddressInfo         AddressInformation         `firestore:"address_info" json:"address_info"`
	Delivery            DeliveryOption             `firestore:"delivery" json:"delivery"`
	PersonalizationData PersonalizationInformation `firestore:"personalization_data" json:"personalization_data"`
}

type PersonalizationInformation struct {
	PersonalData PersonalInformation `firestore:"personal_data" json:"personal_data"`
	AddressData  AddressInformation  `firestore:"address_data" json:"address_data"`
}

type PersonalInformation struct {
	BirthDate   string `firestore:"birth_date" json:"birth_date"`
	CivilState  string `firestore:"civil_state" json:"civil_state"`
	FirstName   string `firestore:"first_name" json:"first_name"`
	LastName    string `firestore:"last_name" json:"last_name"`
	Sex         string `firestore:"sex" json:"sex"`
	HomeNumber  string `firestore:"home_number" json:"home_number"`
	Nationality string `firestore:"nationality" json:"nationality"`
	PassportID  string `firestore:"passport_id" json:"passport_id"`
}

type DeliveryOption struct {
	Type            string             `firestore:"type" json:"type"`
	Service         string             `firestore:"service" json:"service"`
	DeliveryAddress AddressInformation `firestore:"delivery_address" json:"delivery_address"`
}

type AddressInformation struct {
	City       string `firestore:"city" json:"city"`
	Country    string `firestore:"country" json:"country"`
	PostalCode string `firestore:"postal_code" json:"postal_code"`
	Email      string `firestore:"email" json:"email"`
	Street     string `firestore:"street" json:"street"`
}

type Shift struct {
	ID             string  `firestore:"id" json:"id"`
	UID            string  `firestore:"uid" json:"uid"`
	Status         string  `firestore:"status" json:"status"`
	Timestamp      int64   `firestore:"timestamp" json:"timestamp"`
	Payment        Payment `firestore:"payment" json:"payment"`
	FeePayment     Payment `firestore:"fee_payment" json:"fee_payment"`
	RefundAddr     string  `firestore:"refund_addr" json:"refund_addr"`
	ToCoin         string  `firestore:"to_coin" json:"to_coin"`
	ToAmount       int64   `firestore:"to_amount" json:"to_amount"`
	ToAddress      string  `firestore:"to_address" json:"to_address"`
	PaymentProof   string  `firestore:"payment_proof" json:"payment_proof"`
	ProofTimestamp int64   `firestore:"proof_timestamp" json:"proof_timestamp"`
	Message        string  `firestore:"message" json:"message"`
}

type ShiftV2 struct {
	ID                 string           `firestore:"id" json:"id"`
	UID                string           `firestore:"uid" json:"uid"`
	Status             ShiftStatusV2    `firestore:"status" json:"status"`
	InboundTrade       DirectionalTrade `firestore:"inbound_trade" json:"inbound_trade"`
	OutboundTrade      DirectionalTrade `firestore:"outbound_trade" json:"outbound_trade"`
	Timestamp          int64            `firestore:"timestamp" json:"timestamp"`
	Payment            PaymentWithFee   `firestore:"payment" json:"payment"`
	RefundAddr         string           `firestore:"refund_addr" json:"refund_addr"`
	ToCoin             string           `firestore:"to_coin" json:"to_coin"`
	ToAmount           int64            `firestore:"to_amount" json:"to_amount"`
	UserReceivedAmount float64          `firestore:"user_received_amount" json:"user_received_amount"`
	ToAddress          string           `firestore:"to_address" json:"to_address"`
	PaymentProof       string           `firestore:"payment_proof" json:"payment_proof"`
	ProofTimestamp     int64            `firestore:"proof_timestamp" json:"proof_timestamp"`
	Message            string           `firestore:"message" json:"message"`
	OriginalUsdRate    float64          `firestore:"original_rate" json:"original_rate"`
}

type LightShift struct {
	ID                 string       `firestore:"id" json:"id"`
	UID                string       `firestore:"uid" json:"uid"`
	Status             string       `firestore:"status" json:"status"`
	Timestamp          int64        `firestore:"timestamp" json:"timestamp"`
	Payment            LightPayment `firestore:"payment" json:"payment"`
	RefundAddr         string       `firestore:"refund_addr" json:"refund_addr"`
	ToCoin             string       `firestore:"to_coin" json:"to_coin"`
	ToAmount           int64        `firestore:"to_amount" json:"to_amount"`
	UserReceivedAmount float64      `firestore:"user_received_amount" json:"user_received_amount"`
	ToAddress          string       `firestore:"to_address" json:"to_address"`
	PaymentProof       string       `firestore:"payment_proof" json:"payment_proof"`
	ProofTimestamp     int64        `firestore:"proof_timestamp" json:"proof_timestamp"`
}

type DirectionalTrade struct {
	Conversions    []Trade            `firestore:"conversions" json:"conversions"`
	Status         ShiftV2TradeStatus `firestore:"status" json:"status"`
	Exchange       string             `firestore:"exchange" json:"exchange"`
	WithdrawAmount float64            `firestore:"withdraw" json:"withdraw"`
}

type User struct {
	ID         string         `firestore:"id" json:"id"`
	Email      string         `firestore:"email" json:"email"`
	KYCData    KYCInformation `firestore:"kyc_data" json:"kyc_data"`
	Role       string         `firestore:"role" json:"role"`
	Shifts     []string       `firestore:"shifts" json:"shifts"`
	Vouchers   []string       `firestore:"vouchers" json:"vouchers"`
	Deposits   []string       `firestore:"deposits" json:"deposits"`
	Cards      []string       `firestore:"cards" json:"cards"`
	Orders     []string       `firestore:"orders" json:"orders"`
	ShiftV2    []string       `firestore:"shift2" json:"shift2"`
	VouchersV2 []string       `firestore:"vouchers2" json:"vouchers2"`
}

type KYCInformation struct{}

type Voucher struct {
	ID                   string  `firestore:"id" json:"id"`
	UID                  string  `firestore:"uid" json:"uid"`
	VoucherID            int     `firestore:"voucher_id" json:"voucher_id"`
	VariantID            string  `firestore:"variant_id" json:"variant_id"`
	Name                 string  `firestore:"name" json:"name"`
	PaymentData          Payment `firestore:"payment_data" json:"payment_data"`
	FeePayment           Payment `firestore:"fee_payment" json:"fee_payment"`
	BitcouPaymentData    Payment `firestore:"bitcou_payment_data" json:"bitcou_payment_data"`
	BitcouFeePaymentData Payment `firestore:"bitcou_fee_payment_data" json:"bitcou_fee_payment_data"`
	RefundFeeAddr        string  `firestore:"refund_fee_addr" json:"refund_fee_addr"`
	RefundAddr           string  `firestore:"refund_addr" json:"refund_addr"`
	BitcouID             string  `firestore:"bitcou_id" json:"bitcou_id"`
	RedeemCode           string  `firestore:"redeem_code" json:"redeem_code"`
	Status               string  `firestore:"status" json:"status"`
	Timestamp            int64   `firestore:"timestamp" json:"timestamp"`
	RedeemTimestamp      int64   `firestore:"redeem_timestamp" json:"redeem_timestamp"`
	AmountEuro           string  `firestore:"amount_euro" json:"amount_euro"`
	AmountFeeEuro        string  `firestore:"amount_fee_euro" json:"amount_fee_euro"`
	Image                string  `firestore:"image" json:"image"`
	PhoneNumber          int64   `firestore:"phone_nb" json:"phone_nb"`
	ProviderId           int32   `firestore:"provider_id" json:"provider_id"`
	Valid                int32   `firestore:"valid" json:"valid"`
	Message              string  `firestore:"message" json:"message"`
	BitcouRefundData     Payment `firestore:"bitcou_refund_data" json:"bitcou_refund_data"`
	NewType              bool    `firestore:"new_type" json:"new_type"`
}

type VoucherV2 struct {
	Id             string                `firestore:"id" json:"id"`
	CreatedTime    int64                 `firestore:"created_time" json:"created_time"`
	AmountEuro     int64                 `firestore:"amount_euro" json:"amount_euro"`
	UserPayment    Payment               `firestore:"user_payment" json:"user_payment"`
	Status         VoucherStatusV2       `firestore:"status" json:"status"`
	RefundAddress  string                `firestore:"refund_address" json:"refund_address"`
	VoucherId      int                   `firestore:"voucher_id" json:"voucher_id"`
	VariantId      int                   `firestore:"variant_id" json:"variant_id"`
	BitcouTxId     string                `firestore:"bitcou_txid" json:"bitcou_txid"`
	UserId         string                `firestore:"user_id" json:"user_id"`
	RefundTxId     string                `firestore:"refund_txid" json:"refund_txid"`
	FulfilledTime  int64                 `firestore:"fulfilled_time" json:"fulfilled_time"`
	VoucherName    string                `firestore:"voucher_name" json:"voucher_name"`
	PhoneNumber    int64                 `firestore:"phone_number" json:"phone_number"`
	ProviderId     string                `firestore:"provider_id" json:"provider_id"`
	RedeemCode     string                `firestore:"redeem_code" json:"redeem_code"`
	Conversion     DirectionalTrade      `firestore:"conversion" json:"conversion"`
	ReceivedAmount float64               `firestore:"received_amount" json:"received_amount"`
	Email          string                `firestore:"email" json:"email"`
	ShippingMethod VoucherShippingMethod `firestore:"shipping_method" json:"shipping_method"`
	Message        string                `firestore:"message" json:"message"`
}

type LightVoucher struct {
	Id             string                `firestore:"id" json:"id"`
	VoucherId      int                   `firestore:"voucher_id" json:"voucher_id"`
	VariantId      int                   `firestore:"variant_id" json:"variant_id"`
	Name           string                `firestore:"name" json:"name"`
	Timestamp      int64                 `firestore:"timestamp" json:"timestamp"`
	Amount         float64               `firestore:"amount" json:"amount"`
	PaymentTxId    string                `firestore:"payment_txid" json:"payment_txid"`
	PaymentCoin    string                `firestore:"payment_coin" json:"payment_coin"`
	RefundTxId     string                `firestore:"refund_txid" json:"refund_txid"`
	Status         string                `firestore:"status" json:"status"`
	ProviderId     string                `firestore:"provider_id" json:"provider_id"`
	RedeemCode     string                `firestore:"redeem_code" json:"redeem_code"`
	ShippingMethod VoucherShippingMethod `firestore:"shipping_method" json:"shipping_method"`
}

type Deposit struct {
	ID           string  `firestore:"id" json:"id"`
	UID          string  `firestore:"uid" json:"uid"`
	Payment      Payment `firestore:"payment" json:"payment"`
	FeePayment   Payment `firestore:"fee_payment" json:"fee_payment"`
	AmountInPeso string  `firestore:"amount_in_peso" json:"amount_in_peso"`
	CardCode     string  `firestore:"card_code" json:"card_code"`
	Status       string  `firestore:"status" json:"status"`
	Timestamp    string  `firestore:"timestamp" json:"timestamp"`
}

type ExchangeInfo struct {
	Id                  string  `firestore:"id" json:"id"`
	Name                string  `firestore:"name" json:"name"`
	StockCurrency       string  `firestore:"stock_currency" json:"stock_currency"`
	StockExpectedAmount float64 `firestore:"stock_expected_amount" json:"stock_expected_amount"`
	StockMinimumAmount  float64 `firestore:"stock_minimum_amount" json:"stock_minimum_amount"`
	StockMaximumAmount  float64 `firestore:"stock_maximum_amount" json:"stock_maximum_amount"`
	ApiPublicKey        string  `firestore:"api_public_key" json:"api_public_key"`
	ApiPrivateKey       string  `firestore:"api_private_key" json:"api_private_key"`
}

type ExchangeOrderInfo struct {
	Status         ExchangeOrderStatus `json:"status"`
	ReceivedAmount float64             `json:"received_amount"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

type BodyReq struct {
	Payload string `json:"payload"`
}

type TokenVerification struct {
	Valid bool   `json:"valid"`
	UID   string `json:"uid"`
}

type ExchangeOrderStatus int

const (
	ExchangeOrderStatusOpen ExchangeOrderStatus = iota
	ExchangeOrderStatusCompleted
	ExchangeOrderStatusError
)

var (
	ExchangeOrderStatusStr = map[ExchangeOrderStatus]string{
		0: "OPEN",
		1: "COMPLETED",
		2: "ERROR",
	}
)
