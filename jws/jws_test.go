package jws

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/square/go-jose.v2"
	"testing"
)

func TestJWE(t *testing.T) {
	msg := "Encrypted Message"
	key := "RandomKey"
	encrypted, err := EncryptJWE(key, msg)
	assert.Nil(t, err)
	decrypt, err := DecryptJWE(key, encrypted)
	var msgStr string
	err = json.Unmarshal(decrypt, &msgStr)
	assert.Nil(t, err)
	assert.Equal(t, msg, msgStr)

}

func TestDecryptError(t *testing.T) {
	msg := "Encrypted Message"
	key := "RandomKey"
	encrypted, err := EncryptJWE(key, msg)
	decrypt, err := DecryptJWE("", encrypted)
	assert.Nil(t, decrypt)
	assert.NotNil(t, err)
	assert.Equal(t, jose.ErrCryptoFailure, err)
}

func TestDecryptInvalid(t *testing.T) {
	decrypt, err := DecryptJWE("", "invalid-token")
	assert.Nil(t, decrypt)
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("square/go-jose: compact JWE format must have five parts"), err)
}
