package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var Ravencoin = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAIAAAACACAMAAAD04JH5AAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAMAUExURUxpcTQ/ffr89/PMmfn/9TRAgDI9f2Jmm29Tbsy4vPr9+zA8ffv9/ZSWs56jvTZBgfWTL+xNMnB4oWJqmvSqVOr6+u/382Rsm9jW2/SfQmtznsHEzYqPr5OYtF9mlfX49PSjS2Bnlr2/x2ZumnZ8pKKmv210oY+UslhhlPKsY+ybiFhgkY+VtKSovnyDqV9mlfjTu5aatvWxavS+qe2Fcy04evbUrPaVf+6dePTFmZ2mwvHVr/nGjvDCoPONfLe6y/CZiO26o4mWuvnMm0ZQi1phlFxklZ6jvPSvZYiPrvOrYMmipu7Gr+uroFBVj/C7fvS3dnF9qe+wZ5GYtH+Gqe58ZtPX2/BwW0xVjE1WjfOvZfPLoP29cbS0xL6UmMfY5KCtxu58aoaRsjZBgfBSOTZBgO9RODVAgPeVNDdBgTVAgTZBf+9SOTZAgDhCgO1SOO5QNzhCgi45fTE8fiw3e/aVNCo1eS44evBTOvaUMjU/f/BROO5JLu9GKvaPKC86fD5IjPeVMuxFKexGLO5MMjI7ezQ9fuxKMDpGivaSLl5ml/WVMik3fFBZj211oEJMjjJDhfJKLlVnnfBRNnqDqfZ5YO1ONvVYOi02eC88gPSRK0NNiDpFiFlhloKHqu2UgkpTjPONJkdblfBOMoWNsP6fOoySsvZjRvRPMvSWNCcxdcSor05hmWdvnKCftTtOi0BKiexqV7GlsueekWZ1pTdDhu+7lPCij/C9iPaMIlNcmPSIcjM/g++ukv2aMUZVk+p8aT1RkOtbRvC4gO1hS0VPiuiMe/dTNHh8ot+pouxVPfZtUd+im/S4dPiEbNWsq++QfvScOf6zYeyFcPOzaSk6gO11YXN6o1Rdkf+kQfuVLfqSKmp7qXmIsPloTv+pRfxVOfp0WzxGg8yTkrSouP6tVfdxVpaat11tob2vsOOVifZbP/xjSdKnpbudp7OTn5yZsc65qd+RhcC2wdayt5WHoOG/nXF/qe6yiO+igeuniO3W06qMnbuIjeuyppOUrUcVgIwAAAD6dFJOUwD8HgQB/P0EAQIM/BEliv78/7L+/gYY9CX9oTenR78p/Zoe26RO5pnj/f6lrTyRsSiE6GCm/G3+hXVeVIm92SfaP/36me7MdNVCwf74rfKrvM6u1rXF7OmZopmBl569prXq6/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////7///////////////////////////////////////////////4gEt2dAAAPYUlEQVR42rxYa1BTZxo+gQCJYk3kouxqZbZ1V1nber+M93Wm01ad7ezM7s7s7uzszuEk5HLOyUlycvMkEZIQQgBJAgGCiqAoCAhR5CJyp6AoXRGkBS+gVqui66121ba7+50Epx0b6I+e+A6EzPDjfb7nfd7nfb8Pgn5WsCM4ELRkzfYEb0dHV1fX4Y2xUCj0GiMC/C5Z4R1r275owOvWO8x5AAH3taVnQ1DM8hVdY2Y8pa2hf9Hztg6raWzz60MA8izfeHXM7JbiUlN3X2N/dGeRl9+zeRo3lM1+DflDoTkbmspS3LhMhstc1oTBYo8tvrN54N6KGPq/oeyg5w/ZeK+tu62XMgtwHHe5X0Qb7AZDXF9dwvJVkX4MHHYw+Q/Z3MOn8KHu7i4rJcOlMndzvKHcbitu2bBhNPqD9esiIT8RwQHB5URuHnPLpDLK7e0u8uKUu5ff29zv0dq1/c+b4xtHhuMX/eWvITE+EEEggsOBVoy7e3E6rNaOgaKhXr3b3PWs3FZuGGxr6ysu53kMpcM/IILh/Gzuih693uFwlJWV9eT19Nxrqu/uoFK8nTbS09nVUefRknYtj3dpZHh40cfrl0QyLkD2lvrux3cfP74B4rsHD759+vTJN98cv6b3DhbHvXCbm+NsZDKAQJbzbJ6W4bhlf1vO5jJrgHNrnc7Kktb26uq07IsXS05/V/8g03n0TO5AfHS32XR+0JBEklqSBAh4NkPp6IUN3FBGFRC7MEOttlgsqalqVVrVt23Xik5nG9UZJQe3bd/am8L3NhhI+vi8kZHS0fhla1ZFQhxGCWDP2z8dE2FYqrKgZN+NoZ5bDyozUALGLKo9tU/+m0tZ6/p5PM+lltG4ZWvWxsYEYQTN2hMmwtQZ1VkPus86yh7nZFsIFBaKRZgu3wIqcW119CVw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",
		Tag:         "RVN",
		Name:        "Ravencoin (RVN)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: false,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://rvn.polispay.com",
		Protocol:    "rvn",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     175,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18Dash Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: []int{60},
				ScriptHash: []int{122},
				Wif:        []int{128},
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
		CoinGeckoId:      "ravencoin",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "rvn.polispay.com",
		BlockTime:        1,
		MinConfirmations: 4,
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "",
		PubKeyHashAddrID:  []byte{60},
		ScriptHashAddrID:  []byte{122},
		PrivateKeyID:      []byte{128},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        175,
		Base58CksumHasher: base58.Sha256D,
	},
}
