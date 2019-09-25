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
func VerifyPPATToken(service, masterpassword string, tokenHeader string, tokenBody []byte, hestiaAuthUser string, hestiaAuthPassword string, serviceSigningPrivKey string, hestiaPubKey string) (valid bool, payload []byte) {
	valid, uid, _ := hestia.VerifyToken(service, masterpassword, tokenHeader, hestiaAuthUser, hestiaAuthPassword, serviceSigningPrivKey, hestiaPubKey)
	if !valid {
		return false, nil
	}
	if len(tokenBody) > 0 {
		var token string
		err := json.Unmarshal(tokenBody, &token)
		if err != nil {
			return false, nil
		}
		payload, err = jwt.DecryptJWE(uid, token)
		if err != nil {
			return false, nil
		}
	}
	return true, payload
}
