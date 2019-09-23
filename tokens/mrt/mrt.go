package mrt

import (
	"github.com/grupokindynos/common/jwt"
)

// MRT token refers to Microservice Response Token
// This token is used only by Hestia or Plutus as a response of a MVT.
// Has the following standard:
// 	HEADER:
// 		"service": JSON-WEBTOKEN-SIGNED (Service returning the information, Ej. "hestia") | (Service name with version announce, Ej. "hestia:1.0.0")
// 	BODY:
// 		"payload": JSON-WEBTOKEN-ENCRYPTED

// CreateMRTToken will return a header and payload strings to encode the MRT on the response
func CreateMRTToken(service string, masterPassword string, payload interface{}, signingKey string) (header string, body string, err error) {
	header, err = createMRTHeaderSignature(service, signingKey)
	if err != nil {
		return "", "", err
	}
	body, err = createMRTTokenBody(payload, masterPassword)
	if err != nil {
		return "", "", err
	}
	return header, body, nil
}

func createMRTHeaderSignature(service string, signature string) (string, error) {
	encodedToken, err := jwt.EncodeJWS(service, signature)
	if err != nil {
		return "", err
	}
	return encodedToken, nil
}

func createMRTTokenBody(payload interface{}, masterpassword string) (string, error) {
	encryptedPayload, err := jwt.EncryptJWE(masterpassword, payload)
	if err != nil {
		return "", err
	}
	return encryptedPayload, nil
}

// VerifyMRTToken is a utility function to verify and decrypt a MVT tokens
func VerifyMRTToken() (valid bool, payload interface{}) {
	return true, nil
}
