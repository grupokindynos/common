package coins

import (
	"github.com/martinboehm/btcutil/base58"
	"github.com/martinboehm/btcutil/chaincfg"
)

var SnowGem = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAIAAAACACAMAAAD04JH5AAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAMAUExURUxpcf9VVf8/P/8Af+cnOr8+P/8AU/8AAOcoO/8AAOcnOucoOugnOqoLVegoO38ACecoOuYnOucmN+YlPQAAAOgnO+cnO+YoOukoO90mNeUjPegnOtopOucnOugoO+goO+cnOekoOtwgPMwzM+ooOusoO+woO+onOucpOeYnOesnO+YoO+goPOMoNegnO+koO+gnOuYkO+onO+woO+0oPNseOOcoOugmOuknO+gnOucnOucnOusnO+ooO/9/auooO+woO+woPOQqOfgqPugiOOcmOeooO+cnOucoPPApPOwoO+4oPOglOecnOuonOucoO+0oPOonO+YnOuknOucoOuUnOv8qKugnOvAoPO0oPOknOt8jNe0oPM8qKuwoO+snOuknOu0oPOImOe8pPOkoPO0oPOgoOugnOuknOuooO+knOugoOugnO+gnO+0oO+4pPOonO+ooO+goO+woO+0oO+knO+gnOu4oPOgnO+snO+YnOugnOuUmN+0oPOonO+gnOuknOugmOv8zM+ooOuwoO+gnOuwoO+0oO+knOu4oPOgnOqEREewoO+4oPOkoO+knOucnOu0oPOonOu4oPOwoO+knOu0oPOsmPO4oPOwoO+0pO+wnPOgoOukoOu0oPOknO+goO+woO+0oPOkoO+4oO+skOekoO+gnOukoO+soOusoO/UpPe4oPO0oPOooOu8oPO8oPOkoOu4oPOknO+woPOgnOuwoPOonO+cnOu0oPOwoO+ooO+soO+gnOugnOu0oO+4pO+woO+ooO/MpPewoO+YnOfcqPuwoO+cnO+ooO+gnOu0oPOwoO+knOucnOu0oO+gnO+QlOecnOu0pO+gnO+wqPPUpPfIqPeMmOvQqPuwpO/YqPu8qP+gnOvUqP/MqPugoO/grP/crP/orP/UqPusoPPYqPvMqPu0pPO8pPeooPOwpPPQqPvcqP/ApPfEqPfIqPukoO/EpPegoO/8sQuooO+4pPO4pPfIqPfsrQPgqP/YqP/0rQPEpPv8sQPUrPvJcUTIAAADgdFJOUwADBAL7BAMB/wL8/fwD/QL8/CAUAfZi/f4KEJUL+vJjLPwHBXH8gFMeFhz+GBL6Jvgi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",
		Tag:         "XSG",
		Name:        "SnowGem (XSG)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: true,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://xsg.polispay.com",
		Protocol:    "xsg",
		TxVersion:   1,
		TxBuilder:   "bitcore",
		HDIndex:     410,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x86SnowGem Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				// Pending
				PubKeyHash: []int{0x1C, 0x28},
				ScriptHash: []int{0x1C, 0x2D},
				Wif:        []int{0x80},
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "stex",
		FallBackExchange: "mercatox",
		CoinGeckoId:      "snowgem",
	},
	BlockchainInfo: BlockchainInfo{
		ExternalSource:   "xsg.polispay.com",
		BlockTime:        2,
		MinConfirmations: 3,
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   2,
		Bech32HRPSegwit:   "",
		PubKeyHashAddrID:  []byte{0x1C, 0x28},
		ScriptHashAddrID:  []byte{0x1C, 0x2D},
		PrivateKeyID:      []byte{0x80},
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        72,
		Base58CksumHasher: base58.Sha256D,
	},
}
