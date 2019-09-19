package coinfactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoinFactory(t *testing.T) {
	for _, coin := range CoinFactory {
		newCoin, err := GetCoin(coin.Tag)
		assert.Nil(t, err)
		assert.IsType(t, &Coin{}, newCoin)
	}
}

func TestNoCoin(t *testing.T) {
	newCoin, err := GetCoin("NONEXISTINGCOIN")
	assert.NotNil(t, err)
	assert.Equal(t, "coin not available", err.Error())
	assert.Nil(t, newCoin)
}
