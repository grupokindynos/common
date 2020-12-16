package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var Firo = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAIAAAACACAYAAADDPmHLAAAH+UlEQVR4Xu2dT4wURRTGXw27ILuRaAIiplliQoyo9KIbEDhwIV446IkDxkgEhjl40hOJicrBGGKi0QtZZoGEADGcDIk3Y+IFQR0DM9FFTYzCEFAwisLC7uqUadhZZme6Z+pfd73qenucrXpV9b7f+6p6ZrqHAf15nQHm9epp8ZBbAA4E4QQDWGhI44livTpoKBaqMLkAoByE3EZWi/Wq8/lzcgEHguG9DPibNkRPGpMxtnPXxXOHMM1JZC7OAFBeHm4DDsdFFmW7TWMejJR+rX5rex4i46MHwJa9iyRPpA32bQItAK4L3w4HVhDQAZA34bGDgAaAvAuPFQTrAPgmPDYQrAFQDsJTALBB5CCV9zaM8dd3Xax9YGOdVgDwveqThLZxUMwUgP3B8Lo+4GdskO7KmIXbk4t2Xvvhn6zmmxkAVPVykmblBpkAQOLLid9snQUEqQNA4quJnxUEqQEwFjy5ksO8n/SWT72jDDSW9s8vVSrTaWQjFQDKQVgDgKfSmLCvMTnwfbvrtT2m128cgHIQ3gKA+0xPlOIBQIGdL144t8pkLowCQOKblCYhlmEIjAHgm+0/f/IoLFljZ5djjD3IGPvLBG5GAPDtwLfzwlkTudeKUSgUjGhnJIhPl3oYxG+SYwICbQB8Ej9KPCYAovnoQqAFAImv5eLGOutAoAyAb+JjrP5WglQhUALAx0/1sFl/u31kCgBVvzH3NhpIBQJpByDxjWpmPJgsBFIA+Po1Luz230oRY2yQMTYhSpYsAFbuwRNdTBrtXBJf5f0BYQB8tH7sJ/9uwItuBQRAlyy6WP2yLiAEgI/VzwoF2PGLE/d3JiIs4gIEQEL6XK5+GRfoCYCP1b9s/QhsOXEwjTNl5jF7uQABECOJqeo/tOJp4NzuhVOvbxZ3BcDH6t/6xUlY9OiQdqUeXrkWGlOpfI9Tem7dICAA2tJpqvoPDq2RFiqtDkoA+Fj9eRS/CVUSBIkOQACo1yOm6lcCoLx89QvA2SfqKXCvZ56rP1KjMdC/pPRj5Vq7MrEOQNWvDjDG6u/mAgSAwe/5YRY/giDuHNABAMaHMKrXo1jPvNv/bBY4bC5eqn7empUOAHyzf2/En1G93QUIAEM3eWC3/6RzgNcA+Fb9ceeAOQCQ/YudGdpbuVL90bwZwNSuenVBcw3eAuBj9cdtAwSAWtHf6XXr6h9wfGSzRgQ7XVsPgl4C4HP1t58DvAOgb2AhbD//pXbpffXO+1AbPaIdx0aAWAcoB+FNABiwMaEsx/S9+mdy/XuxXl06cyi8+5IPVwDr3ngNVpe2a/N2ZNVGmL4pfO+F9nhpBGi6wOwWgAEAU9WZRsJMxsRw2YgOAF/Eb4JkGwJUAPgmPgYICACTnq4Yy6YLEACKopnsRgDMZJO2AJNYicUiBxDLU6qtyAHIAVIFrFtwcgBrqb83MDkAOYA1DMkBrKWeHKAj9XQVkD2N5ADZ57xjRDoD0BnAGoaoHCDKgm/bgM3qj/KNDgCfILAtPloAdPxw2ca1sOXjsk6IO30/3boDrpxx++lgIknocICxILzBAQZFOmNsY2oLwVCdGeT3SrFeXRaNk4svhT53+CMY2rxJO2+eiD/nLuFcAEDVL8d+ru4LeOXnb6DQ1yeXgZjWvlR/6wEwF1sAVb88+7lxABJfXnwAmCzWq7M/7ev07eEEgDwAuXlABIkvL377/t9xBohewHCDiMjSCACRLHW26ekAY8uHd3DOUT8qm8RXE78BsL5Ur55p7e3kY+IIADUAhB4Th30bIPHVxI/b/2PPANGLoyvCZwr/QUV9qPR6EgBquZV6VCxWFyDx1cRPqv5EByAA1BONtaf04+KxQUDVr46W0g9GYAJgYMli2Fb5TD0DMz2r+w/D1+9+qB3HtQDKAGCAoH9wAF4eP2Uk5z594tdMmNaPRkVBGo2G3Z+9MiI9wNHVm2Dy+t+GorkTRhuAvEBA1R8Pbc/fDcwDAD6K3+3SrxUFIQBch8BHAHpZfxOC3APgo/ii1d/1jaC4HcPFA6GPAIhWvzQAnPMHOOd/unIG9lF84LxUvFQ7IKqR8BbQDOiSC/gIgEz1SzuASxCQ+GIeIO0ArkDgGwBTk//e/+rV72+IyX6vVS4B8E18mVN/OyDKAGB+b8A3AGT3faU3gpKsBduhkMSX2wS0HADjecAnAHQqX/qdwF5cYXACn8RvLO2fX6pUpnvp0uv/RhwgGoRz/hDn/LdeA6bx/+8OHoPTe99LIzTKmBz4vt312h4TkzMGQDSZ8tDwODT44yYmRjESMsD5ePFS7QlT+TEKwB0IguFxAILAlEBz4hgWX/mdwF6LKwfh2wDwVq929H/xDJi0faOXgd2W4MqNpuIy2Glp6sAXN3vjW0D7IASBHjQmLvW6zSB1AO6eC8JcfLFUT0r53mmLn9oZIG6pBIEcAFmInykA0WAngg0Lr4Pjv7kqp6N0a9VP9aQHmumQyRZA5wIxebKq+syuAroteyxY8yKHxjGx1OS7FYfCS7vrZ63kwooDtMrp+9nARtWjcADftwXbwjfzb90BfAMBi/BoAWhOLG9bAzbh0QOQFxCwCu8MAM2Jjj42srgwMX3VheuBpAcyYZw7ujOASJJGg/DZAsBpkbZZtYl7CGNWY+uM4yQAMQfH2wCwQCcRCn3nPHVboT+KLrkAIC6T5SC8DAAPm8kyu1ysn3vETCxcUXILAK40450NAYBXm0xm9j9yjzqueCnjFAAAAABJRU5ErkJggg==",
		Tag:         "FIRO",
		Name:        "Firo (FIRO)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://xzc.polispay.com",
		Protocol:    "xzc",
		TxVersion:   1,
		TxBuilder:   "bitcoinjs",
		HDIndex:     136,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18ZCoin Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: []int{82},
				ScriptHash: []int{7},
				Wif:        []int{210},
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "bittrex",
		CoinGeckoId:      "zcoin",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "xzc.polispay.com",
		BlockTime:        5,
		MinConfirmations: 2,
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "",
		ScriptHashAddrID:  []byte{7},
		PubKeyHashAddrID:  []byte{82},
		PrivateKeyID:      []byte{210},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        136,
		Base58CksumHasher: base58.Sha256D,
	},
}
