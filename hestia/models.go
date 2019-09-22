package hestia

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

type Card struct {
	Address    string `bson:"address" json:"address"`
	CardCode   string `bson:"card_code" json:"card_code"`
	CardNumber string `bson:"card_number" json:"card_number"`
	City       string `bson:"city" json:"city"`
	Email      string `bson:"email" json:"email"`
	FirstName  string `bson:"firstname" json:"firstname"`
	LastName   string `bson:"lastname" json:"lastname"`
	UID        string `bson:"uid" json:"uid"`
}

type Pin struct {
	CardCode string `bson:"card_code" json:"card_code"`
	PinCode  string `bson:"pin_code" json:"pin_code"`
}

type Coin struct {
	Ticker            string   `bson:"ticker" json:"ticker"`
	ShiftAvailable    bool     `bson:"shift_available" json:"shift_available"`
	DepositAvailable  bool     `bson:"deposit_available" json:"deposit_available"`
	VouchersAvailable bool     `bson:"vouchers_available" json:"vouchers_available"`
	OrdersAvailable   bool     `bson:"orders_available" json:"orders_available"`
	Balances          Balances `bson:"balances" json:"balances"`
}

type Balances struct {
	HotWallet float64 `bson:"hot_wallet" json:"hot_wallet"`
	Exchanges float64 `bson:"exchanges" json:"exchanges"`
}

type Payment struct {
	Address       string `bson:"address" json:"address"`
	Amount        string `bson:"amount" json:"amount"`
	Coin          string `bson:"coin" json:"coin"`
	Txid          string `bson:"txid" json:"txid"`
	Confirmations string `bson:"confirmations" json:"confirmations"`
}

type Properties struct {
	FeePercentage int  `bson:"fee_percentage" json:"fee_percentage"`
	Available     bool `bson:"available" json:"available"`
}

type Config struct {
	Shift    Properties `bson:"shift" json:"shift"`
	Deposits Properties `bson:"deposits" json:"deposits"`
	Vouchers Properties `bson:"vouchers" json:"vouchers"`
	Orders   Properties `bson:"orders" json:"orders"`
}

type Order struct {
	ID                  string                     `bson:"id" json:"id"`
	UID                 string                     `bson:"uid" json:"uid"`
	Status              string                     `bson:"status" json:"status"`
	PaymentInfo         Payment                    `bson:"payment_info" json:"payment_info"`
	AddressInfo         AddressInformation         `bson:"address_info" json:"address_info"`
	Delivery            DeliveryOption             `bson:"delivery" json:"delivery"`
	PersonalizationData PersonalizationInformation `bson:"personalization_data" json:"personalization_data"`
}

type PersonalizationInformation struct {
	PersonalData PersonalInformation `bson:"personal_data" json:"personal_data"`
	AddressData  AddressInformation  `bson:"address_data" json:"address_data"`
}

type PersonalInformation struct {
	BirthDate   string `bson:"birth_date" json:"birth_date"`
	CivilState  string `bson:"civil_state" json:"civil_state"`
	FirstName   string `bson:"first_name" json:"first_name"`
	LastName    string `bson:"last_name" json:"last_name"`
	Sex         string `bson:"sex" json:"sex"`
	HomeNumber  string `bson:"home_number" json:"home_number"`
	Nationality string `bson:"nationality" json:"nationality"`
	PassportID  string `bson:"passport_id" json:"passport_id"`
}

type DeliveryOption struct {
	Type            string             `bson:"type" json:"type"`
	Service         string             `bson:"service" json:"service"`
	DeliveryAddress AddressInformation `bson:"delivery_address" json:"delivery_address"`
}

type AddressInformation struct {
	City       string `bson:"city" json:"city"`
	Country    string `bson:"country" json:"country"`
	PostalCode string `bson:"postal_code" json:"postal_code"`
	Email      string `bson:"email" json:"email"`
	Street     string `bson:"street" json:"street"`
}

type Shift struct {
	ID         string  `bson:"id" json:"id"`
	UID        string  `bson:"uid" json:"uid"`
	Status     string  `bson:"status" json:"status"`
	Timestamp  string  `bson:"timestamp" json:"timestamp"`
	Payment    Payment `bson:"payment" json:"payment"`
	Conversion Payment `bson:"conversion" json:"conversion"`
}

type User struct {
	ID       string         `bons:"id" json:"id"`
	Email    string         `bson:"email" json:"email"`
	KYCData  KYCInformation `bson:"kyc_data" json:"kyc_data"`
	Role     string         `bson:"role" json:"role"`
	Shifts   []string       `bson:"shifts" json:"shifts"`
	Vouchers []string       `bson:"vouchers" json:"vouchers"`
	Deposits []string       `bson:"deposits" json:"deposits"`
	Cards    []string       `bson:"cards" json:"cards"`
	Orders   []string       `bson:"orders" json:"orders"`
}

type KYCInformation struct{}

type Voucher struct {
	ID                string  `bson:"id" json:"id"`
	UID               string  `bson:"uid" json:"uid"`
	VoucherID         string  `bson:"voucher_id" json:"voucher_id"`
	VariantID         string  `bson:"variant_id" json:"variant_id"`
	FiatAmount        string  `bson:"fiat_amount" json:"fiat_amount"`
	Name              string  `bson:"name" json:"name"`
	PaymentData       Payment `bson:"payment_data" json:"payment_data"`
	BitcouPaymentData Payment `bson:"bitcou_payment_data" json:"bitcou_payment_data"`
	RedeemCode        string  `bson:"redeem_code" json:"redeem_code"`
	Status            string  `bson:"status" json:"status"`
	Timestamp         string  `bson:"timestamp" json:"timestamp"`
}

type Deposit struct {
	ID           string  `bson:"id" json:"id"`
	UID          string  `bson:"uid" json:"uid"`
	Payment      Payment `bson:"payment" json:"payment"`
	AmountInPeso string  `bson:"amount_in_peso" json:"amount_in_peso"`
	CardCode     string  `bson:"card_code" json:"card_code"`
	Status       string  `bson:"status" json:"status"`
	Timestamp    string  `bson:"timestamp" json:"timestamp"`
}
