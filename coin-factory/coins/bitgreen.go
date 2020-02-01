package coins

import (
	"github.com/eabz/btcutil/base58"
	"github.com/eabz/btcutil/chaincfg"
)

// Bitgreen coinfactory information
var Bitgreen = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAIAAAACACAMAAAD04JH5AAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAMAUExURUxpcYn/Zyq0SIzJPrz8o2lsRiaVK4LAaf///2//Kf3/+yK0RZLJP/7//XbOhVK+RsrlqiK0S3HEQ5nMQpLaop3coZ3NQki7R3bFQ8zqvL7fixuzSxWxTKHOQr/koW3DQFa9Rfb7+zu4RmnNher45WjIcji8YbjbdP7+/r7egur36mvCRX/NVjuzRtzv2NPvxoTIQrjhk5nWhajPQHHLcWXOW1nIe0O4RbHWXzO2RqPWcS22RpnVdLXjr47JQky7Ra7UWmfBQ43RZqPbh37HRFy/RS2nR7bioobOYIDLXMrruz2+aInIQJvWg2PAQpPVc5HXnaHYlKTdoqvae6vZcXfPZH3QiGTCRhOwS7vouJHWhByOQJPVcIrXl1LEZaHSWIrViaHajUXBa7raXoPFjheDO4zOYMfmoVPBWYDGQWfHYabUYGHAQTK3SaLZfJ7bi43Ra6XVZX7QeFnKd5/VW8PjklDBZmrMfrfeisHfdVK7bYbRaXDOgWDGbEjBYYrOWGHKeTrBbHHLb2bHZHnNaULEcrrdhU6xcDiuUs/psEelZJDXZYnNXKXUZFLGfZjaprLaaJDXay6bVGfDSWLBRmTCRpHKQ1/ARnDERV3AR2zDRXXFRXjFRSq2Siy2SmfCRpLLQ2nCRmjCRim1SmvDRia1SXPERVq/R6LdWFu/R5XLQz3FXznGVaPXU1e/R47KQ0G6SC63SjK3Sja4SXzGQxezSqfPQUa7SBy0Sjm6SDy5SYDIRIrJQ4TIREy9SKLPQYbVUETESozJQ27IQkK+RmPNUavaRHjRWFnNWYDQRofIRFvKSwyzSDWrRSedQnnRSo3XVjLBTKHYRHzUUE3LXR+WQW/NSVa/Qp3RPx67S6nTPkDIWnfLRYTWWUrHV7XdRDrCT4jSRZXWSpfPQJzcVCW+TjChRXHSUJXaVojKQ2PLSRy/UpvYUBS9UEXKX7/mTk7FShS6TI/URYLTSSi1RTTEUW7OXqHeSy/HU5rQSkC6VEWvQzPGaS+6Ule8TKrbThMApw4AAADndFJOUwAC/PwDAQIEAQEF/f0Sav9w/vz+RW7+/vwwP/3+/W3+/AL8qSHB8cAJqw79/v4XHfxNkPu+/sH94PzA/L4q/f3w/OR9/vz+WfrhQOL8bfugazdFiKzxg/36S3X+0lDz8qpg46Vv/sU/6fzw5f79m6yw17ao/XrZxGyVudOU0vjpu/Lf4KfLQZX+cMb+cn5qYtHy7P7//////////////////////////////////////////////////////////////////////////////////////////////////////////////jhGn8kAAAuISURBVHja7Jt5UBRXHoBf48x0zyCyNauBsUouxzgCDkdkiZwLyFkc4bQ4RAgqnoC3xrOSVGJ0U3G3RFEQUeTY4AFoSeGBA0hQKS4RcCOgIkKgKKW8jZWtfe/1DAyHME03+MfmK62ZefN+8/t4/a7ungHgT5hBkRBqaCkq5E9OfvpBOMp7E5xfZ3tc3GbRoGTQZnNc3HYLMPFtQPHlB589ePAsTkaotQFhvB0Vfm4hICZagATbn30GebYN8PoLecAKFz6IUyucGIR80ee/fnbp0uBcJPjxASy89OvfTSe6GwiBUuCXuYBQE5inFNCZFIFLH1sA5po1VOAXJHBpkgRmXb8+XAAVXv8/Ebj+8QVmzfoECggGC0AmSeDeJ5A7QwVQ4b0/BSZFYM29v0KGCtxBhY+XTqAAwUOQPJlKAD5XQfJUAnJVEUUJuf3LB8ZcfwuooxIQqS+cPJ6Qu5YHkviZkO9mxq95/DPk3kH4XMV3M+feQYWPl1rRBfE+clNjvFBz4wAX4ZBNZ5X8TPPurDrvlKWq17lL1xwM2RxP4Q7CGj4hCnkzVUUGnSpjqjrDCmFz3Hlzdv12K2P1wzfuXZAkd0oGMw4fzjg81elx79mDnuwbgQSB75wOjw/oMFeiNmLH2wKbpowz/2GnKe2btgEBu+mBNLZ7M0XFeRqnKR/GCXH+fH/N3GCWm1VKaBHS25GL6S3Dn1qWkftBOjp629szynxVDmW+vcEs+wGaBwx++Avkh3ntvmVlZU7n1+OXI2KwzS8hZP3S3ue+qCrCt4N1G/R3I4Nc36uQ53ZjRIikViGb2rtx5atXfXut2I4Fei0Q8Qw60Gf6PrcTDawFwyGxsCR46XOlQfcaHYKL8zYeMOiw7O7uvg1bQDDG7MmDE5BFSIflbVi927I9mJOzJixwGzKmAG41EhgHt+P6ty03SQDBicBvlteu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",
		Tag:         "BITG",
		Name:        "Bitgreen (BITG)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		Blockbook:   "https://bitg.polispay.com",
		Protocol:    "bitg",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     222,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18Bitgreen Signed Message:\n",
				Bech32:        "bg",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: 38,
				ScriptHash: 6,
				Wif:        46,
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "crex24",
		FallBackExchange: "sistemkoin",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        2,
		MinConfirmations: 1,
		ExternalSource:   "bitg.polispay.com",
	},
	NetParams: &chaincfg.Params{
		Bech32HRPSegwit:   "bg",
		PubKeyHashAddrID:  []byte{38},
		ScriptHashAddrID:  []byte{6},
		PrivateKeyID:      []byte{46},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        222,
		Base58CksumHasher: base58.Sha256D,
	},
}
