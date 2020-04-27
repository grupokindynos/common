package errors

import "errors"

var (
	ErrorNoAuth            = errors.New("you are not authorized")
	ErrorHeaderToken       = errors.New("no token found in header")
	ErrorFbInitializeAuth  = errors.New("unable to initialize auth client")
	ErrorNoUserInformation = errors.New("unable to get user information")
	ErrorMissingID         = errors.New("missing id param")
	ErrorInfoDontMatchUser = errors.New("information requested doesn't match for this user")
	ErrorCoinDataGet       = errors.New("unable to get coin information")
	ErrorConfigDataGet     = errors.New("unable to get config information")
	ErrorDecryptJWE        = errors.New("unable to decrypt jwe")
	ErrorDBStore           = errors.New("unable to store information to database")
	ErrorNotFound          = errors.New("information not found")
	ErrorAlreadyExists     = errors.New("object already exists")
	ErrorVoucherLimit      = errors.New("user daily voucher purchase exceeded")
	ErrorNotEnoughDash     = errors.New("not enough dash to pay vouchers")

	ErrorWrongMessage      = errors.New("signed message is not on known hosts")
	ErrorInvalidPassword   = errors.New("could not decrypt using master password")
	ErrorUnmarshal         = errors.New("unable to unmarshal object")
	ErrorSignatureParse    = errors.New("could not parse header signature")
	ErrorNoHeaderSignature = errors.New("no signature found in header")
	ErrorUnknownMethod     = errors.New("missing known method, currently support are GET and POST")

	ErrorShiftMinimumAmount = errors.New("minimum amount for a shift should be above $18.00 USD")
	// Open Shift
	ErrorServiceUnavailable        = errors.New("the service you requested is currently unavailable")
	ErrorAssetUnavailable          = errors.New("the asset you requested is currently unavailable for this service")
	ErrorObtainingRates            = errors.New("rates for your conversion could not be calculated")
	ErrorFillingPaymentInformation = errors.New("payment information could not be filled")
	ErrorShiftNotFound             = errors.New("we could not match the provided id with our registry. likely due to the shift time expiring")
)
