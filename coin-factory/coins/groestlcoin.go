package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var Groestlcoin = Coin{
	Info: CoinInfo{
		Icon:        "iVBORw0KGgoAAAANSUhEUgAAAIAAAACACAMAAAD04JH5AAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAMAUExURUxpcWWrv2PB1IKlrWWtwWTH2mW+0WKou5K3ua/z/WW1yWGzyGW5zb3n7Y/V4m+ntnnN3pPV42qtwIzO3Gmpu3equG+mtoHS4mu2yYrU5GuktWitv3PJ22qounSntHHJ22unuXPA0Wu7z2urvGu5y2y+0ZTI1HTG2HPA0XrT5HLL3GmsvnrN3XbK3H3V5WvB1XTS5GiyxG7C1W/I2n3N3mynuGzB1P////7+/nfO3/3//7DW4P///////////4y6yOPr7Mfq8f///wCevgCcuwCYtwCPrwCaugCSsgCfvwCKqwCJqgCcvACTswCOrgCQsACFpgCRsv7//wCVtgCIqQCiwQCMrAB9ngCNrQB4mQB7nACCowCZuQCYuACHqACauQCgwAB5mgCUtQCGpwCEpQB/oQB+nwB9nwCjwgB2lwCevQB/oACBogCBoQB7nQB8nQCXtgB3mAB5m/7//wB6mwCLrPz+/wCDpACZuAB1lgB+oP3+/gCRsQGDowGlw2XE2GaxxgBzlQCFpf/+/mXJ3GW7z2/M3mvD123G2WTH2mu1yfv9/my5zWizx2m3y2zA02rJ3G7J3Gyww2avxABxkgCNsAOTsmy90AWJqWnG2W2zxgBtjxyrx2KqvmzQ5GfB1AKWtx2YtNXr8Pf7/O33+XHV6WS3y+bz9jurxGnN4MPW29zv8yqtyEaSplKcsCqMpDS2z83p71a1y2eqvA6JpwuEol6frw95lg6SsWvT51Cnvp3Q3Zi9xpHN2wWMrC2ivDuMotXg41CVqIfH1vXz86XV4P77+rnP1D6XrV+luDObtCCQrB2EoLXe5xJ/nK3Iz0eqwHe7yxeNqsDi6hx9l/r394i1v1yuww+pxgl8m1zE2JHE0XDY7D63zxt3kUOetQdri67Y4ufs7Diiu0m70glzkV65zk+uxSmTrhKcuu7x8R+hvS2Em1W+06DI0SaBmhWlwq/P1wChwcvj6XrD1M7b3namsoO/zWujsUi0yoq9y77e5t7m6HituabByNni5NxhsKAAAACBdFJOUwD9/Q/+/P4DAQL9//wPUFG+NP4rzSddlZVdlunR4DTnvU7epLvOGpVh4PH1nqzO5/rV3Pl/gPi5Tdgv1NeJisniXf///////////////////////////////////////////////////////////////////////////////////me0GC4AABVKSURBVHjatJh7UNNXFsd/owNJKLAwtrJMGdluW62OWqtuZ/z38ggkkAcEyAqkmAAyKy5M04DjLxAXQkgIj2TZBNKJIbw0KCuPFReUx7Q+cLUsPkZLqYyzylDR1dlxtdNt63bP/SWBACEEmj0zwAyE3+d7vufcc3/3EsQqgk6jUT+DQsP3bg+JiNy2Rw2xZ1tkRMj2veGhQdQfaTQ68X8JOgPTGRvCw0Ii9RaLpaNDqVHYQ6Pr6LCo1JEhYeEbGFgDw/caKHpQcNg+VXu7RalYGOuoUJg6LBZ1xPbgzZQG3+Iho6DwHdvaBzo0Cj9AmkxqvV6vdonKyiasoslkUUXuCA7CjvkOD1+hYVvbB5R+AMdok+Lp9FR3Z6eZis7Ozqnpp+sqQQUlQtGhitwS6vhHn+DpwSGq9g6ga3R6vebpVOdESnxUgj2imUxmBgQz4V9m2/S6SuyEv79CVRkSDLbRfFF7RvjO9gENRVdPTplL46OiouLj8fffQVAiorGMzBhrZvQTW29TZV2Tv7+/ThURzvi5vUCHDIJ34uQ1Jr1+stucAuSsrPgF/Gg7/yOIIzFWceFsZ28dJcGkiggGE+g/y/3QEAqvU2mmzXnxWSkpWcvyMzKOZGZmxsSIxTFPptapK+tAgj4kdO11gPSDwnYN2PHdE1nPU1JW5MfExFiZVqGkpLH5xgyYUKHbs2X9Gk3A7m+9p/HTaABfmpWS5xW/MMMq6Xo1furU6LOvX8zMQCVUu99akwk4fXDfgc/L85LP5HSN30H2uNJycaauoqJJj01Yvf2hO+9h9/XTE8/zvOUnM1nnLiFpK4RWSyJEXugFBRWqiNBVloFOJ8J3WfyUStWkmcJ7x4/5iHPuilRrz5+UIrIIDdlAQWDlh+/hZ66i+xlh95S4+aZK80q958eIu+4A/vTYhQtjbQgVISmJXmv2Bwn++i107wcjgwgKuUdV3/y8dBX8QmbZKYR6emflBgN39kUrIlGRtLW5v7cBTFDt8LoRGMRmKD/YPw3pryb/GP79cjQkNiSKIeQy20lQoEUXRFwbVqB/d7N3HtCIDVsHML87b3X8wgzBMzRk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",
		Tag:         "GRS",
		Name:        "GroestlCoin (GRS)",
		Trezor:      true,
		Ledger:      true,
		Segwit:      true,
		Masternodes: false,
		Token:       false,
		Blockbook:   "https://grs.polispay.com",
		Protocol:    "grs",
		TxVersion:   1,
		TxBuilder:   "groestljs",
		HDIndex:     17,
		Networks: map[string]CoinNetworkInfo{
			"P2SHInP2WPKH": {
				MessagePrefix: "\x18Groestlcoin Signed Message:\n",
				Bech32:        "grs",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x049d7cb2,
					Private: 0x049d7878,
				},
				PubKeyHash: 36,
				ScriptHash: 5,
				Wif:        80,
			},
			"P2WPKH": {
				MessagePrefix: "\x18Groestlcoin Signed Message:\n",
				Bech32:        "grs",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x04b24746,
					Private: 0x04b2430c,
				},
				PubKeyHash: 36,
				ScriptHash: 5,
				Wif:        80,
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "grs.polispay.com",
		BlockTime:        10,
		MinConfirmations: 2,
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "grs",
		ScriptHashAddrID:  []byte{5},
		PubKeyHashAddrID:  []byte{36},
		PrivateKeyID:      []byte{80},
		Base58CksumHasher: base58.Groestl512D,
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        17,
	},
}
