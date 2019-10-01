package hestia

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

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
	Ticker            string   `firestore:"ticker" json:"ticker"`
	ShiftAvailable    bool     `firestore:"shift_available" json:"shift_available"`
	DepositAvailable  bool     `firestore:"deposit_available" json:"deposit_available"`
	VouchersAvailable bool     `firestore:"vouchers_available" json:"vouchers_available"`
	OrdersAvailable   bool     `firestore:"orders_available" json:"orders_available"`
	Balances          Balances `firestore:"balances" json:"balances"`
}

type Balances struct {
	HotWallet float64 `firestore:"hot_wallet" json:"hot_wallet"`
	Exchanges float64 `firestore:"exchanges" json:"exchanges"`
}

type Payment struct {
	Address       string  `firestore:"address" json:"address"`
	Amount        float64 `firestore:"amount" json:"amount"`
	Coin          string  `firestore:"coin" json:"coin"`
	RawTx         string  `firestore:"rawtx" json:"rawtx"`
	Txid          string  `firestore:"txid" json:"txid"`
	Confirmations int32   `firestore:"confirmations" json:"confirmations"`
}

type Properties struct {
	FeePercentage int  `firestore:"fee_percentage" json:"fee_percentage"`
	Available     bool `firestore:"available" json:"available"`
}

type Config struct {
	Shift    Properties `firestore:"shift" json:"shift"`
	Deposits Properties `firestore:"deposits" json:"deposits"`
	Vouchers Properties `firestore:"vouchers" json:"vouchers"`
	Orders   Properties `firestore:"orders" json:"orders"`
}

type Order struct {
	ID                  string                     `firestore:"id" json:"id"`
	UID                 string                     `firestore:"uid" json:"uid"`
	Status              string                     `firestore:"status" json:"status"`
	PaymentInfo         Payment                    `firestore:"payment_info" json:"payment_info"`
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
	Timestamp  string  `firestore:"timestamp" json:"timestamp"`
	Payment    Payment `firestore:"payment" json:"payment"`
	Conversion Payment `firestore:"conversion" json:"conversion"`
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
	ID                string  `firestore:"id" json:"id"`
	UID               string  `firestore:"uid" json:"uid"`
	VoucherID         int     `firestore:"voucher_id" json:"voucher_id"`
	VariantID         string  `firestore:"variant_id" json:"variant_id"`
	FiatAmount        float64 `firestore:"fiat_amount" json:"fiat_amount"`
	Name              string  `firestore:"name" json:"name"`
	PaymentData       Payment `firestore:"payment_data" json:"payment_data"`
	BitcouPaymentData Payment `firestore:"bitcou_payment_data" json:"bitcou_payment_data"`
	RedeemCode        string  `firestore:"redeem_code" json:"redeem_code"`
	Status            string  `firestore:"status" json:"status"`
	Timestamp         string  `firestore:"timestamp" json:"timestamp"`
}

type Deposit struct {
	ID           string  `firestore:"id" json:"id"`
	UID          string  `firestore:"uid" json:"uid"`
	Payment      Payment `firestore:"payment" json:"payment"`
	AmountInPeso string  `firestore:"amount_in_peso" json:"amount_in_peso"`
	CardCode     string  `firestore:"card_code" json:"card_code"`
	Status       string  `firestore:"status" json:"status"`
	Timestamp    string  `firestore:"timestamp" json:"timestamp"`
}

type BodyReq struct {
	Payload string `json:"payload"`
}

type TokenVerification struct {
	Valid bool   `json:"valid"`
	UID   string `json:"uid"`
}
