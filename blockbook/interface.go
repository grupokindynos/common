package blockbook

type BlockBookInterface interface {
	GetXpub(url string, xpub string)
	GetEthAddress(url string, addr string)
	GetUtxo(url string, xpub string)
	GetTx(url string, txid string)
	GetTxEth(url string, txid string)
	SendTx(url string, rawTx string)
	FindDepositTxId(address string, amount int64) (string, error)
	SendTxWithMessage(url string, rawTx string)
}
