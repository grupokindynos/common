package explorer

type Explorer interface {
	GetXpub(xpub string) (response Xpub, err error)
	GetAddress(address string) (response Address, err error)
	GetEthAddress(addr string) (response EthAddr, err error)
	GetUtxo(xpub string, confirmed bool) (response []Utxo, err error)
	GetTx(txid string) (response Tx, err error)
	GetTxEth(txid string) (response EthTx, err error)
	GetFee(nBlocks string) (response Fee, err error)
	SendTx(rawTx string) (response string, err error)
	SendTxWithMessage(rawTx string) (string, error, string)
	FindDepositTxId(address string, amount int64) (string, error)
}
