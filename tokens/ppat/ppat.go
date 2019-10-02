package ppat

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/grupokindynos/common/hestia"
	"github.com/grupokindynos/common/jwt"
)

// PPAT token refers to PolisPay Authentication Token.
// Is only used from the PolisPay app to talk with Hestia or another microservice.
// Has the following standard:
// 	HEADER:
// 		"token": FBT
// 	BODY:
// 		"payload": JSON-WEBTOKEN-ENCRYPTED

// CreatePPATToken will return a http.Request encoded using the PPAT format
func CreatePPATToken(method string, reqUrl string, fbt string, uid string, payload interface{}) (req *http.Request, err error) {
	switch method {
	case "GET":
		req, err = http.NewRequest(method, reqUrl, nil)
		if err != nil {
			return nil, err
		}
	case "POST":
		tokenBytes, err := createPPATTokenBody(payload, uid)
		if err != nil {
			return nil, err
		}
		payload := bytes.NewBuffer(tokenBytes)
		req, err = http.NewRequest(method, reqUrl, payload)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("missing known method, currently support are GET and POST")
	}
	req.Header.Set("token", fbt)
	return req, nil
}

func createPPATTokenBody(payload interface{}, uid string) ([]byte, error) {
	encryptedPayload, err := jwt.EncryptJWE(uid, payload)
	if err != nil {
		return nil, err
	}
	return json.Marshal(encryptedPayload)
}

// VerifyPPATToken is a utility function to verify and decrypt a PPAT token (only must be used for external microservices, Hestia can verify itself)
func VerifyPPATToken(productionURL string, service string, masterpassword string, tokenHeader string, tokenBody string, hestiaAuthUser string, hestiaAuthPassword string, serviceSigningPrivKey string, hestiaPubKey string) (valid bool, payload []byte, uid string, err error) {
	valid, uid, err = hestia.VerifyToken(productionURL, service, masterpassword, tokenHeader, hestiaAuthUser, hestiaAuthPassword, serviceSigningPrivKey, hestiaPubKey)
	if !valid {
		return false, nil, "", err
	}
	if tokenBody != "" {
		payload, err = jwt.DecryptJWE(uid, tokenBody)
		if err != nil {
			return false, nil, "", errors.New("unable to decrypt token")
		}
	}
	return true, payload, uid, nil
}
