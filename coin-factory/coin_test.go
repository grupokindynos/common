package coinfactory

import (
	"github.com/grupokindynos/common/coin-factory/coins"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoinFactory(t *testing.T) {
	for _, coin := range Coins {
		newCoin, err := GetCoin(coin.Tag)
		assert.Nil(t, err)
		assert.IsType(t, &coins.Coin{}, newCoin)
	}
}

func TestNoCoin(t *testing.T) {
	newCoin, err := GetCoin("NONEXISTINGCOIN")
	assert.NotNil(t, err)
	assert.Equal(t, "coin not available", err.Error())
	assert.Nil(t, newCoin)
}
