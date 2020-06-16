package explorer

import (
	"fmt"
	coinfactory "github.com/grupokindynos/common/coin-factory"
	"testing"
)

func TestExplorerFactory(t *testing.T) {
	fmt.Println("HOla")
	ef := NewExplorerFactory()
	nuls, _ := coinfactory.GetCoin("NULS")
	explorer, _ := ef.GetExplorerByCoin(*nuls)
	_, err := explorer.GetAddress("NULSd6HggApZFb5bvWkksdgXEy8WV4nHxjzE")
	fmt.Println(err)
}

func TestNoCoin(t *testing.T) {
	/*newCoin, err := GetCoin("NONEXISTINGCOIN")
	assert.NotNil(t, err)
	assert.Equal(t, "coin not available", err.Error())
	assert.Nil(t, newCoin)*/
}
