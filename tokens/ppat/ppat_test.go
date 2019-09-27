package ppat

import (
	"encoding/json"
	"github.com/grupokindynos/common/jwt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const (
	TestURL        = "localhost/test-url"
	TestMasterPass = "1234567890"
	TestFBToken    = "firebase-jwt"
)

var (
	TestPayload = []string{"string1", "string2"}
)

func TestCreatePPATPOSTToken(t *testing.T) {
	req, err := CreatePPATToken("POST", TestURL, TestFBToken, TestMasterPass, TestPayload)
	assert.Nil(t, err)
	token := req.Header.Get("token")
	url := req.URL
	assert.Equal(t, TestFBToken, token)
	assert.Equal(t, TestURL, url.Path)
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

func TestCreatePPATGETToken(t *testing.T) {
	req, err := CreatePPATToken("GET", TestURL, TestFBToken, TestMasterPass, nil)
	assert.Nil(t, err)
	token := req.Header.Get("token")
	url := req.URL
	assert.Equal(t, TestFBToken, token)
	assert.Equal(t, TestURL, url.Path)
}
