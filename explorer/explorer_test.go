package explorer

import (
	coinfactory "github.com/grupokindynos/common/coin-factory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExplorerFactory(t *testing.T) {
	ef := NewExplorerFactory()
	// nuls
	nuls, _ := coinfactory.GetCoin("NULS")
	explorer, _ := ef.GetExplorerByCoin(*nuls)
	res, err := explorer.GetAddress("NULSd6HggApZFb5bvWkksdgXEy8WV4nHxjzED")
	assert.Nil(t, err)
	assert.Equal(t, "NULSd6HggApZFb5bvWkksdgXEy8WV4nHxjzED", res.Address)
	txdata, err := explorer.GetTx("d405a19d7e603144ba6a0643db7d47383eb79ae1c7a6a6063346b9e41f170b5b")
	assert.Nil(t, err)
	assert.Equal(t, 2341860, txdata.BlockHeight)
	fee, err := explorer.GetFee("123")
	assert.Nil(t, err)
	assert.NotNil(t, fee.Result)
	_, err = explorer.SendTx("02005d76e15e00008c01170100bf314ba28b38d554bad918ff63c6f033501000100a0ebce1d00000000000000000000000000000000000000000000000000000000083c3434371322edda000117010001fe64a23f125613b5c7c5e6bbf31f053f30bbb529010001000065cd1d0000000000000000000000000000000000000000000000000000000000000000000000006a21030fdbbf38148cdf9d43a28421ae4ba5f47acc2a01b72516bb48b48b9748be0d5b473045022100b9934e28c58e79cc3a2ac21f135a004191c5ecb9d2c91e55dd674bf0fbfd210f02201527342821e345e2a3e2ba12298e920e752953f31d9c60ac73d02533344e541b")
	assert.NotNil(t, err)
	assert.Equal(t, "Deserialize error;Deserialize error", err.Error())
	txid, err := explorer.FindDepositTxId("NULSd6HgjWe7C7chnLMLHbXVSuCDPLR5S8h1t", 600000000)
	assert.Nil(t, err)
	assert.Equal(t, "d4bdc0ad6408b6dfa569d34547efe0291b49e2bb5f83c3764c3a11c7247a2e8b", txid)
	// blockbook
	polis, _ := coinfactory.GetCoin("POLIS")
	explorer, _ = ef.GetExplorerByCoin(*polis)
	res, err = explorer.GetAddress("PQooE9ecq4mi3Us3vVigCZehuDJ3nTDrtf")
	assert.Nil(t, err)
	assert.Equal(t, "PQooE9ecq4mi3Us3vVigCZehuDJ3nTDrtf", res.Address)
	txdata, err = explorer.GetTx("2c56355e9925a945ffef411fa40fc9755892933bc854b4aa55fcf34f21842a49")
	assert.Nil(t, err)
	assert.Equal(t, 622459, txdata.BlockHeight)
	txid, err = explorer.FindDepositTxId("PHwa4qU9o5TPT5bSdr9A1dXjyF3qEGPbio", 1847829923)
	assert.Nil(t, err)
	assert.Equal(t, "2c56355e9925a945ffef411fa40fc9755892933bc854b4aa55fcf34f21842a49", txid)

}
