package blockbook

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
	Address		string	`json:"address"`
	Balance		string	`json:"balance"`
	ItemsOnPage	int	`json:"itemsOnPage"`
	Page		int	`json:"page"`
	Tokens		[]struct {
		Balance		string	`json:"balance"`
		Decimals	int	`json:"decimals"`
		Name		string	`json:"name"`
		Path		string	`json:"path"`
		TotalReceived	string	`json:"totalReceived"`
		TotalSent	string	`json:"totalSent"`
		Transfers	int	`json:"transfers"`
		Type		string	`json:"type"`
	}	`json:"tokens"`
	TotalPages	int	`json:"totalPages"`
	TotalReceived	string	`json:"totalReceived"`
	TotalSent	string	`json:"totalSent"`
	Transactions	[]struct {
		BlockHash	string	`json:"blockHash"`
		BlockHeight	int	`json:"blockHeight"`
		BlockTime	int	`json:"blockTime"`
		Confirmations	int	`json:"confirmations"`
		Fees		string	`json:"fees"`
		Hex		string	`json:"hex"`
		LockTime	*int	`json:"lockTime,omitempty"`
		Txid		string	`json:"txid"`
		Value		string	`json:"value"`
		ValueIn		string	`json:"valueIn"`
		Version		int	`json:"version"`
		Vin		[]struct {
			Addresses	[]string	`json:"addresses"`
			Hex		*string		`json:"hex,omitempty"`
			N		int		`json:"n"`
			Sequence	int		`json:"sequence"`
			Txid		string		`json:"txid"`
			Value		string		`json:"value"`
			Vout		*int		`json:"vout,omitempty"`
		}	`json:"vin"`
		Vout	[]struct {
			Addresses	[]string	`json:"addresses"`
			Hex		string		`json:"hex"`
			N		int		`json:"n"`
			Spent		*bool		`json:"spent,omitempty"`
			Value		string		`json:"value"`
		}	`json:"vout"`
	}	`json:"transactions"`
	Txs			int	`json:"txs"`
	UnconfirmedBalance	string	`json:"unconfirmedBalance"`
	UnconfirmedTxs		int	`json:"unconfirmedTxs"`
	UsedTokens		int	`json:"usedTokens"`
}


type Tx struct {
	BlockHash	string	`json:"blockHash"`
	BlockHeight	int	`json:"blockHeight"`
	BlockTime	int	`json:"blockTime"`
	Confirmations	int	`json:"confirmations"`
	Fees		string	`json:"fees"`
	Hex		string	`json:"hex"`
	LockTime	int	`json:"lockTime"`
	Txid		string	`json:"txid"`
	Value		string	`json:"value"`
	ValueIn		string	`json:"valueIn"`
	Version		int	`json:"version"`
	Vin		[]struct {
		Addresses	[]string	`json:"addresses"`
		Hex		string		`json:"hex"`
		N		int		`json:"n"`
		Sequence	int		`json:"sequence"`
		Txid		string		`json:"txid"`
		Value		string		`json:"value"`
	}	`json:"vin"`
	Vout	[]struct {
		Addresses	[]string	`json:"addresses"`
		Hex		string		`json:"hex"`
		N		int		`json:"n"`
		Spent		bool		`json:"spent"`
		Value		string		`json:"value"`
	}	`json:"vout"`
}