package errors

import "errors"

var (
	ErrorNoAuth            = errors.New("you are not authorized")
	ErrorFbInitializeAuth  = errors.New("unable to initialize auth client")
	ErrorDbInitialize      = errors.New("unable to connect to database")
	ErrorNoUserInformation = errors.New("unable to get user information")
	ErrorUnmarshal         = errors.New("unable to unmarshal object")
	ErrorMissingID         = errors.New("missing id param")
	ErrorInfoDontMatchUser = errors.New("information requested doesn't match for this user")
	ErrorCoinDataGet       = errors.New("unable to get coin information")
	ErrorConfigDataGet     = errors.New("unable to get config information")
	ErrorDecryptJWE        = errors.New("unable to decrypt jwe")
	ErrorDBStore           = errors.New("unable to store information to database")
	ErrorNotFound          = errors.New("information not found")
	ErrorAlreadyExists     = errors.New("object already exists")
)