package blockbook

type blockBookInterface interface {
	GetXpub(url string, xpub string)
	GetEthAddress(url string, addr string)
	GetUtxo(url string, xpub string)
	GetTx(url string, txid string)
	GetTxEth(url string, txid string)
	SendTx(url string, rawTx string)
}
