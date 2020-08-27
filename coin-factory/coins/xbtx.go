package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var BitcoinSubsidium = Coin{
	Info: CoinInfo{
		Icon:        "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAASABIAAD/7QA4UGhvdG9zaG9wIDMuMAA4QklNBAQAAAAAAAA4QklNBCUAAAAAABDUHYzZjwCyBOmACZjs+EJ+/+EATEV4aWYAAE1NACoAAAAIAAGHaQAEAAAAAQAAABoAAAAAAAOgAQADAAAAAQABAACgAgAEAAAAAQAAAPqgAwAEAAAAAQAAAPoAAAAA/9sAQwACAgICAgECAgICAwICAwMGBAMDAwMHBQUEBggHCQgIBwgICQoNCwkKDAoICAsPCwwNDg4PDgkLEBEQDhENDg4O/9sAQwECAwMDAwMHBAQHDgkICQ4ODg4ODg4ODg4ODg4ODg4ODg4ODg4ODg4ODg4ODg4ODg4ODg4ODg4ODg4ODg4ODg4O/8AAEQgAMgAyAwEiAAIRAQMRAf/EABsAAQADAAMBAAAAAAAAAAAAAAAHCAkBBQoE/8QALhAAAQMEAQIEBQQDAAAAAAAAAQIDBAAFBhEHEiEIMUFRCRMiYXEUFRYyFyNy/8QAHAEAAgICAwAAAAAAAAAAAAAAAAYDBQECBAcI/8QALREAAgIBAQYEBQUAAAAAAAAAAQIAEQMEBRIhMVFhBhNBcRQiMqHwUoGR0eH/2gAMAwEAAhEDEQA/AN/KUrpchyC14viUy93iQY8KOnaulBWtZ9EISO6lE9gkdyajd0xoXc0BxJPICZAJNCd1SspOZfHNyNa3bq3h2OxMRt7HUhh+4wFXGYtY/qXEpWlpr7pHzCPc+VQJiHxKuRoGdhm+XK2ZLby6QuLdrEm3udG+wS8wshKvTakkfalsbf0DE7m8w6heH9/aWo2dquRAB7mbrUqEuFOeMN5vwZdxx5a7feIqU/uVnlLSX4pVsBQI7ONqIIS4nsdEEJUCkTbTDhzYtRiGTEbU+srXR8bFXFERSlKnkcVTvlLlG1yvFp/jBb3TLjQWEo6inpjrkJU46/o+awwlKUexWo+tXErDrx3LyLBfiLqyaE47DbutqgzokpvfTptBiuA/8qQgkeyx70r7fXK+zyqepF9K7/vUs9AmN89P0lv8uxzH3ON5uR2tlmHDZdMW3GPGS7IeX19CUtdQP1LXsDQ6j3JNUVz7gfN81daNxXGlQ32VKWcmeQ0qCdEFSXgn5oKD30nYOtEaNTNxzzvYrXxrh9qut2i3LKWG1PobQFOQ7c44lRbQtzyL2idgb6Sojtqpsav+PZ/hj+fXyJ/HcOjsgQLetofq7vLPZSygbKh17baaHdZ2sjWq85J8S5ZtKvluLFuvA9WU2KA40LqqIJuOG+2IURa/q9ewqZleHe93nhH4kOPO2bKHrpYbbJTBny3Ya47d2gulKH0JbUo6SlRLjZV3Py0HQ2BXpSHlWLF18PF5XkOH9cJKMnnym1z2kq2YypEkdDOx59AUR7f6z6VtOPKu9vC2py58ORXN7u6DdXvV810BxPA8h7Ra2oqh1Km7E5pSlP8AKGKrT4nvDtaPELwei0mQ1asuta1yLBc3UlSG1qAC2XQO5ac0netlJSlQB6dGy1KjyImVCjiwZurMjBl5ieYvPeMeTeI8kcs+cYzPskhtSktSVgpYfSnyU2+NtOg+mjv3A8hJHCfJGeRM0iCwWN3MbsFfLgtCK5NkME9ttoQDr8jX3Oq9EsiNHlxVsSmG5DCxpTbqApKvyD2r5oNqtlsbUi3W6NASr+wjR0tg/npApOzeG9Nmb6uHsL/n/JeDauULRX8/O8rjwfxzmLK2s05LQYt7cJdiWtxxK3WVqT0l58p+nrCfpS2kkIBPqQlFnqUpo0mkw6LD5WIUPuT1PUykyOcjFjFKUrmyOKUpRCKUpRCKUpRCKUpRCf/Z",
		Tag:         "XBTX",
		Name:        "Bitcoin Subsidium (XBTX)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://xbtx.polispay.com",
		Protocol:    "xbtx",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     0,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18IDX Signed Message:\n",
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
		Exchange:         "mock",
		FallBackExchange: "birake",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "xbtx.polispay.com",
		BlockTime:        1,
		MinConfirmations: 1,
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{122},
		PubKeyHashAddrID:  []byte{60},
		PrivateKeyID:      []byte{128},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        0,
		Base58CksumHasher: base58.Sha256D,
	},
}
