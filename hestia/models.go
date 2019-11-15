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

type Properties struct {
	FeePercentage float64 `firestore:"fee_percentage" json:"fee_percentage"`
	Available     bool    `firestore:"available" json:"available"`
}

type Config struct {
	Shift    Available `firestore:"shift" json:"shift"`
	Deposits Available `firestore:"deposits" json:"deposits"`
	Vouchers Available `firestore:"vouchers" json:"vouchers"`
	Orders   Available `firestore:"orders" json:"orders"`
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
	ID         string  `firestore:"id" json:"id"`
	UID        string  `firestore:"uid" json:"uid"`
	Status     string  `firestore:"status" json:"status"`
	Timestamp  int64   `firestore:"timestamp" json:"timestamp"`
	Payment    Payment `firestore:"payment" json:"payment"`
	FeePayment Payment `firestore:"fee_payment" json:"fee_payment"`
	ToCoin     string  `firestore:"to_coin" json:"to_coin"`
	ToAmount   int64   `firestore:"to_amount" json:"to_amount"`
	ToAddress  string  `firestore:"pay_address" json:"pay_address"`
}

type User struct {
	ID       string         `firestore:"id" json:"id"`
	Email    string         `firestore:"email" json:"email"`
	KYCData  KYCInformation `firestore:"kyc_data" json:"kyc_data"`
	Role     string         `firestore:"role" json:"role"`
	Shifts   []string       `firestore:"shifts" json:"shifts"`
	Vouchers []string       `firestore:"vouchers" json:"vouchers"`
	Deposits []string       `firestore:"deposits" json:"deposits"`
	Cards    []string       `firestore:"cards" json:"cards"`
	Orders   []string       `firestore:"orders" json:"orders"`
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
	BitcouID             string  `firestore:"bitcou_id" json:"bitcou_id"`
	RedeemCode           string  `firestore:"redeem_code" json:"redeem_code"`
	Status               string  `firestore:"status" json:"status"`
	Timestamp            int64   `firestore:"timestamp" json:"timestamp"`
	RedeemTimestamp      int64   `firestore:"redeem_timestamp" json:"redeem_timestamp"`
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
