package mvt

import (
	"encoding/json"
	"github.com/grupokindynos/common/jwt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const (
	TestURL          = "localhost/test-url"
	TestService      = "service-test"
	TestMasterPass   = "1234567890"
	TestAuthUser     = "test-user"
	TestAuthPassword = "test-password"
	TestSignPrivKey  = "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"
	TestSignPubKey   = "MIICCgKCAgEAk3BM5fWCydl5ME++Y2RvE3ZgI53B1gZImCsQnyothOzUVHFxiLnScJWcPsP5+K2RPRXY8A6RhkOcnc9uSop51i800pq0fNVSrKvRofgl61o73UdYA5BdNueT7MVge0I6ioscgwIp/brJoSV2Vy5U0LdNvwEhkYF5W4FVK2VrVne2cnoWYOyx9fTMpPRq9YBVprwnU7wIKsnIjqa0xiH8kussphl0xpMzlJDVHCH3gp6HaCiHj5XEbWTe2y+tuJbynLMdIjvyjj4Jo+ThTyw2PRFqjIt+O/jm4FMWU2tgSg5pRj3NSMsuvutXPqo2m51SOOk4ECCgy/gNatwZdl38XFsJnEzJWt/TQvrf61bICcrMzrAkVDycbAYaJGNPILo309giQJUhs6hRf+Qdan+XI9hlFazzfqdFlhxgR4MUfE9RjUm2ypRWmym6JvbLeU4qsoiF/g/XZqumG+PGx21LMWCilg8SSemYBBKbHS8zUHUZiQ+yuW7r7FI9WEHEg47KNiVeDoc0IRMYPSBzMgYh2XjoivLWjZokzVzYBUhrHlSIqaZVAwFQ9AAkHvL4QAZRP3MG4emiDaZ3bX+hDRbvhs6L4Jqia1UCy7F9B2I42lpF/3gMrUlWCzkHUfD2m7x0ry7Sdbua3WVt9SW9HpiDckzTdmrCQsuXVO9t4H43gz0CAwEAAQ=="
)

var (
	TestPayload = []string{"string1", "string2"}
)

func TestCreateMVTPOSTToken(t *testing.T) {
	req, err := CreateMVTToken("POST", TestURL, TestService, TestMasterPass, TestPayload, TestAuthUser, TestAuthPassword, TestSignPrivKey)
	assert.Nil(t, err)
	url := req.URL
	assert.Equal(t, TestURL, url.Path)
	headerSignature := req.Header.Get("service")
	headerSignatureBytes, err := jwt.DecodeJWSNoVerify(headerSignature)
	assert.Nil(t, err)
	var serviceStr string
	err = json.Unmarshal(headerSignatureBytes, &serviceStr)
	assert.Nil(t, err)
	assert.Equal(t, TestService, serviceStr)
	var serviceStrValidating string
	headerSignatureBytesVerified, err := jwt.DecodeJWS(headerSignature, TestSignPubKey)
	assert.Nil(t, err)
	err = json.Unmarshal(headerSignatureBytesVerified, &serviceStrValidating)
	assert.Nil(t, err)
	body, _ := ioutil.ReadAll(req.Body)
	var bodyToken string
	err = json.Unmarshal(body, &bodyToken)
	assert.Nil(t, err)
	decryptedBody, err := jwt.DecryptJWE(TestMasterPass, bodyToken)
	assert.Nil(t, err)
	var bodyFormat []string
	err = json.Unmarshal(decryptedBody, &bodyFormat)
	assert.Nil(t, err)
	assert.Equal(t, TestPayload, bodyFormat)
}

func TestCreateMVTGETToken(t *testing.T) {
	req, err := CreateMVTToken("GET", TestURL, TestService, TestMasterPass, TestPayload, TestAuthUser, TestAuthPassword, TestSignPrivKey)
	assert.Nil(t, err)
	url := req.URL
	assert.Equal(t, TestURL, url.Path)
	headerSignature := req.Header.Get("service")
	headerSignatureBytes, err := jwt.DecodeJWSNoVerify(headerSignature)
	assert.Nil(t, err)
	var serviceStr string
	err = json.Unmarshal(headerSignatureBytes, &serviceStr)
	assert.Nil(t, err)
	assert.Equal(t, TestService, serviceStr)
	var serviceStrValidating string
	headerSignatureBytesVerified, err := jwt.DecodeJWS(headerSignature, TestSignPubKey)
	assert.Nil(t, err)
	err = json.Unmarshal(headerSignatureBytesVerified, &serviceStrValidating)
	assert.Nil(t, err)
}

func TestVerifyMVTPOSTToken(t *testing.T) {
	req, err := CreateMVTToken("POST", TestURL, TestService, TestMasterPass, TestPayload, TestAuthUser, TestAuthPassword, TestSignPrivKey)
	assert.Nil(t, err)
	headerSignature := req.Header.Get("service")
	body, err := ioutil.ReadAll(req.Body)
	assert.Nil(t, err)
	var bodyStr string
	err = json.Unmarshal(body, &bodyStr)
	assert.Nil(t, err)
	valid, payload := VerifyMVTToken(headerSignature, bodyStr, TestSignPubKey, TestMasterPass)
	assert.Equal(t, true, valid)
	var bodyFormat []string
	err = json.Unmarshal(payload, &bodyFormat)
	assert.Nil(t, err)
	assert.Equal(t, TestPayload, bodyFormat)
}

func TestVerifyMVTGETToken(t *testing.T) {
	req, err := CreateMVTToken("GET", TestURL, TestService, TestMasterPass, nil, TestAuthUser, TestAuthPassword, TestSignPrivKey)
	assert.Nil(t, err)
	headerSignature := req.Header.Get("service")
	assert.Nil(t, err)
	valid, payload := VerifyMVTToken(headerSignature, "", TestSignPubKey, TestMasterPass)
	assert.Equal(t, true, valid)
	assert.Equal(t, []byte(nil), payload)
}
