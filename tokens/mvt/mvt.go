package mvt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/grupokindynos/common/jwt"
	"net/http"
)

// MVT token refers to Microservice Verification Token
// This token is used only by microservices talking to Hestia or Plutus.
// Has the following standard:
// 	BASIC-AUTH:
// 		USER: HESTIA_AUTH_USERNAME | PLUTUS_AUTH_USERNAME
// 		PASS: HESTIA_AUTH_PASSWORD | PLUTUS_AUTH_PASSWORD
// 	HEADER:
// 		"service": JSON-WEBTOKEN-SIGNED (Service name asking for the information, Ej. "ladon") | (Service name with version announce, Ej. "ladon:1.0.0")
// 	BODY:
// 		"payload": JSON-WEBTOKEN-ENCRYPTED

// CreateMVTToken will return a http.Request encoded using the MVT format
func CreateMVTToken(method string, reqUrl string, service string, masterPassword string, payload interface{}, basicUser string, basicPassword string, signingKey string) (req *http.Request, err error) {
	switch method {
	case "GET":
		req, err = http.NewRequest(method, reqUrl, nil)
		if err != nil {
			return nil, err
		}
	case "POST":
		tokenBytes, err := createMVTTokenBody(payload, masterPassword)
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
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(basicUser+":"+basicPassword)))
	headerSig, err := createMVTHeaderSignature(service, signingKey)
	if err != nil {
		return nil, err
	}
	req.Header.Set("service", headerSig)
	return req, nil
}

func createMVTHeaderSignature(service string, signature string) (string, error) {
	encodedToken, err := jwt.EncodeJWS(service, signature)
	if err != nil {
		return "", err
	}
	return encodedToken, nil
}

func createMVTTokenBody(payload interface{}, masterpassword string) ([]byte, error) {
	encryptedPayload, err := jwt.EncryptJWE(masterpassword, payload)
	if err != nil {
		return nil, err
	}
	return json.Marshal(encryptedPayload)
}

// VerifyMVTToken is a utility function to verify and decrypt a MVT tokens
func VerifyMVTToken(tokenHeader string, tokenBody []byte, servicePubKey string, masterPassword string) (valid bool, payload interface{}) {
	_, err := jwt.DecodeJWS(tokenHeader, servicePubKey)
	if err != nil {
		return false, nil
	}
	if len(tokenBody) > 0 {
		var token string
		err = json.Unmarshal(tokenBody, &token)
		if err != nil {
			return false, nil
		}
		payload, err = jwt.DecryptJWE(masterPassword, token)
		if err != nil {
			return false, nil
		}
	}
	return true, payload
}
