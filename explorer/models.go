package explorer

type Status struct {
	Backend struct {
		BestBlockHash   string `json:"bestBlockHash"`
		Blocks          int    `json:"blocks"`
		Chain           string `json:"chain"`
		Difficulty      string `json:"difficulty"`
		Headers         int    `json:"headers"`
		ProtocolVersion string `json:"protocolVersion"`
		SizeOnDisk      int    `json:"sizeOnDisk"`
		Subversion      string `json:"subversion"`
		Version         string `json:"version"`
	} `json:"backend"`
	Blockbook struct {
		About           string `json:"about"`
		BestHeight      int    `json:"bestHeight"`
		BuildTime       string `json:"buildTime"`
		Coin            string `json:"coin"`
		DbSize          int    `json:"dbSize"`
		Decimals        int    `json:"decimals"`
		GitCommit       string `json:"gitCommit"`
		Host            string `json:"host"`
		InSync          bool   `json:"inSync"`
		InSyncMempool   bool   `json:"inSyncMempool"`
		InitialSync     bool   `json:"initialSync"`
		LastBlockTime   string `json:"lastBlockTime"`
		LastMempoolTime string `json:"lastMempoolTime"`
		MempoolSize     int    `json:"mempoolSize"`
		SyncMode        bool   `json:"syncMode"`
		Version         string `json:"version"`
	} `json:"blockbook"`
}

type Xpub struct {
	Address            string   `json:"address"`
	Balance            string   `json:"balance"`
	ItemsOnPage        int      `json:"itemsOnPage"`
	Page               int      `json:"page"`
	Tokens             []Tokens `json:"tokens"`
	TotalPages         int      `json:"totalPages"`
	TotalReceived      string   `json:"totalReceived"`
	TotalSent          string   `json:"totalSent"`
	Transactions       []Tx     `json:"transactions"`
	Txs                int      `json:"txs"`
	UnconfirmedBalance string   `json:"unconfirmedBalance"`
	UnconfirmedTxs     int      `json:"unconfirmedTxs"`
	UsedTokens         int      `json:"usedTokens"`
}

type Tokens struct {
	Balance       string `json:"balance"`
	Decimals      int    `json:"decimals"`
	Name          string `json:"name"`
	Path          string `json:"path"`
	TotalReceived string `json:"totalReceived"`
	TotalSent     string `json:"totalSent"`
	Transfers     int    `json:"transfers"`
	Type          string `json:"type"`
}

type Utxo struct {
	Address       string `json:"address"`
	Confirmations int    `json:"confirmations"`
	Height        int    `json:"height"`
	Path          string `json:"path"`
	Txid          string `json:"txid"`
	Value         string `json:"value"`
	Vout          int    `json:"vout"`
}

type TxVin struct {
	Addresses []string `json:"addresses"`
	Hex       string   `json:"hex"`
	N         int      `json:"n"`
	Sequence  int      `json:"sequence"`
	Txid      string   `json:"txid"`
	Value     string   `json:"value"`
	Vout      int      `json:"vout"`
}

type TxVout struct {
	Addresses []string `json:"addresses"`
	Hex       string   `json:"hex"`
	N         int      `json:"n"`
	Spent     bool     `json:"spent"`
	Value     string   `json:"value"`
}

type TokenTransfer struct {
	Type     string `json:"type"`
	From     string `json:"from"`
	To       string `json:"to"`
	Token    string `json:"token"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Value    string `json:"value"`
}

type Tx struct {
	BlockHash     string   `json:"blockHash"`
	BlockHeight   int      `json:"blockHeight"`
	BlockTime     int      `json:"blockTime"`
	Confirmations int      `json:"confirmations"`
	Fees          string   `json:"fees"`
	Hex           string   `json:"hex"`
	LockTime      int      `json:"lockTime"`
	Txid          string   `json:"txid"`
	Value         string   `json:"value"`
	ValueIn       string   `json:"valueIn"`
	Version       int      `json:"version"`
	Vin           []TxVin  `json:"vin"`
	Vout          []TxVout `json:"vout"`
	TokenTransfers []TokenTransfer `json:"tokenTransfers"`
}

type Fee struct {
	Result string `json:"result"`
}

type EthAddr struct {
	Address            string      `json:"address"`
	Balance            string      `json:"balance"`
	ItemsOnPage        int         `json:"itemsOnPage"`
	NonTokenTxs        int         `json:"nonTokenTxs"`
	Nonce              string      `json:"nonce"`
	Page               int         `json:"page"`
	Tokens             []EthTokens `json:"tokens"`
	TotalPages         int         `json:"totalPages"`
	Transactions       []EthTx     `json:"transactions"`
	Txs                int         `json:"txs"`
	UnconfirmedBalance string      `json:"unconfirmedBalance"`
	UnconfirmedTxs     int         `json:"unconfirmedTxs"`
}

type EthTokens struct {
	Balance   string `json:"balance"`
	Contract  string `json:"contract"`
	Decimals  int    `json:"decimals"`
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	Transfers int    `json:"transfers"`
	Type      string `json:"type"`
}

type EthTxVin struct {
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
	N         int      `json:"n"`
}

type EthTxVout struct {
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
	N         int      `json:"n"`
	Value     string   `json:"value"`
}

type EthTx struct {
	BlockHash        string `json:"blockHash"`
	BlockHeight      int    `json:"blockHeight"`
	BlockTime        int    `json:"blockTime"`
	Confirmations    int    `json:"confirmations"`
	EthereumSpecific struct {
		GasLimit int    `json:"gasLimit"`
		GasPrice string `json:"gasPrice"`
		GasUsed  int    `json:"gasUsed"`
		Nonce    int    `json:"nonce"`
		Status   int    `json:"status"`
	} `json:"ethereumSpecific"`
	Fees           string `json:"fees"`
	TokenTransfers []struct {
		Decimals int    `json:"decimals"`
		From     string `json:"from"`
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		To       string `json:"to"`
		Token    string `json:"token"`
		Type     string `json:"type"`
		Value    string `json:"value"`
	} `json:"tokenTransfers"`
	Txid  string      `json:"txid"`
	Value string      `json:"value"`
	Vin   []EthTxVin  `json:"vin"`
	Vout  []EthTxVout `json:"vout"`
}

type SendTx struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

type Address struct {
	Page               int      `json:"page"`
	TotalPages         int      `json:"totalPages"`
	ItemsOnPage        int      `json:"itemsOnPage"`
	Address            string   `json:"address"`
	Balance            string   `json:"balance"`
	TotalReceived      string   `json:"totalReceived"`
	TotalSent          string   `json:"totalSent"`
	UnconfirmedBalance string   `json:"unconfirmedBalance"`
	UnconfirmedTxs     int      `json:"unconfirmedTxs"`
	Txs                int      `json:"txs"`
	Txids              []string `json:"txids"`
}
