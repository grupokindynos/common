package coins

// Coin is the basic coin structure to get the correct properties for each coin.
type Coin struct {
	Tag            string
	Name           string
	Rates          RatesSource
	BlockchainInfo BlockchainInfo
	RpcMethods     RPCMethods
	Keys           Keys
	ColdAddress    string
	BlockExplorer  string
	Token          bool
	TokenNetwork   string
	Contract       string
}

// Keys is the struct to wrap all credentials to hot-wallets
type Keys struct {
	RpcUser string
	RpcPass string
	RpcPort string
	Host    string
	Port    string
	User    string
	PrivKey string
}

// RatesSource is the prefered source of exchange for rates
type RatesSource struct {
	Exchange         string // The main exchange used to get rate information.
	FallBackExchange string // The fallback exchange to get if first fail.
}

// BlockchainInfo is a model to get information for a particular blockchain
type BlockchainInfo struct {
	BlockTime        float64 // Time between blocks in minutes.
	MinConfirmations int     // Minimum confirmations to mark a tx as usable.
	ExternalSource   string  // External source of information to double validate agains another trusted source
}

// RPCMethods is a model to wrap possible RPC commands for different coins
type RPCMethods struct {
	GetWalletInfo              string      // Get wallet information for balances.
	GetWalletAccounts          string      // Specific ETH method for accounts in wallet.
	GetBlockchainInfo          string      // Get blockchain information for sync status.
	GetNetworkInfo             string      // Get network information to know how many peers are connected.
	GetNewAddress              string      // Get a new address to pay to the hot-wallets
	SendToAddress              string      // Send a tx.
	ValidateAddress            string      // Validate an address to see if is ours.
	DecodeRawTransaction       string      // Get all the information of a raw transaction.
	GetRawTransaction          string      // Command to get a transaction
	GetRawTransactionVerbosity interface{} // Variable to get it encoded or decoded.
}
