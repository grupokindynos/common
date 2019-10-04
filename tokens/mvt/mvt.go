package mvt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/grupokindynos/common/jwt"
	"github.com/gin-gonic/gin"
	"github.com/grupokindynos/common/errors"
	"github.com/grupokindynos/common/jwt"
	"io/ioutil"
	"os"
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
		return nil, errors.ErrorUnknownMethod
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
func VerifyMVTToken(tokenHeader string, tokenBody string, servicePubKey string, masterPassword string) (valid bool, payload []byte) {
	_, err := jwt.DecodeJWS(tokenHeader, servicePubKey)
	if err != nil {
		return false, nil
	}

	if tokenBody != "" {
		payload, err = jwt.DecryptJWE(masterPassword, tokenBody)
		if err != nil {
			return false, nil
		}
	}
	return true, payload
}

// VerifyRequest is a wrapper around VerifyMRTToken to get information correctly from the GIN context
func VerifyRequest(c *gin.Context) (payload []byte, err error) {
	headerSignature := c.GetHeader("service")
	if headerSignature == "" {
		return nil, errors.ErrorNoHeaderSignature
	}
	decodedHeader, err := jwt.DecodeJWSNoVerify(headerSignature)
	if err != nil {
		return nil, errors.ErrorSignatureParse
	}
	var serviceStr string
	err = json.Unmarshal(decodedHeader, &serviceStr)
	if err != nil {
		return nil, errors.ErrorUnmarshal
	}
	// Check which service the request is announcing
	var pubKey string
	switch serviceStr {
	case "ladon":
		pubKey = os.Getenv("LADON_PUBLIC_KEY")
	case "tyche":
		pubKey = os.Getenv("TYCHE_PUBLIC_KEY")
	case "adrestia":
		pubKey = os.Getenv("ADRESTIA_PUBLIC_KEY")
	default:
		return nil, errors.ErrorWrongMessage
	}
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	var reqToken string
	if len(reqBody) > 0 {
		err := json.Unmarshal(reqBody, &reqToken)
		if err != nil {
			return nil, errors.ErrorUnmarshal
		}
	}
	valid, payload := VerifyMVTToken(headerSignature, reqToken, pubKey, os.Getenv("MASTER_PASSWORD"))
	if !valid {
		return nil, errors.ErrorInvalidPassword
	}
	return payload, nil
}
