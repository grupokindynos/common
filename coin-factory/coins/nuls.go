package coins

import (
	"github.com/eabz/btcutil/base58"
	"github.com/eabz/btcutil/chaincfg"
)

// Nuls coinfactory information
var Nuls = Coin{
	Info: CoinInfo{
		Icon:        "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAMAAACdt4HsAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAL3UExURUxpcf////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////7//v3+/IG5Pf7+/f3+/f///n+4Op/Ka3+3OazRf322Np3IaIW7Q6rPfPf783y2NYm9SY7AUMHcn/b68Ya8RYO6QIC4O4e8RoS6QX63OH23N/7+/o2/T4a7RIi8R4K5Poi9SMLdoff68oC4PKnPenu2NMLdorrYlOnz3vv9+X63OYS7Qoi9R4O6QZHBVXm0MbfXkIq9St/tzrLUiPb68rXWjfX57/r8947AUfD36KPMcYe7RarQfJ7Jac3jsqLLcIm9SIa8RHy2NKDKbH+4OdXovpzHZqjPeZnGYo/AUvz9+oK5P/n79dnqxcTepI6/UHezLprGY3q1M9bowI/AUYu+S7/cnfr8+HazLOXw1vz++6vQfeny3Y/BUs/ktPX68N7tzY3AUPX58PH36uXx2M/ktdHluIy/TbDThnq1MprHY7HUiK7SgpbFXqPMcurz3sDbneHu0Xi0L4y/TsPeo6DLbX22N/P47cnhq323OIO5P+rz3/b68LPUicHcoNfpweTw1aHKboq+S5rHZNfowYW7QsfgqYG5PIi8SOfy2svirsbfp9vryIq+Svv9+sXfppvHZf39+7bWjqbNduXw17/bnJPDWN7tzO/254C5PPP47JXEW6nQe6/Sg63RgJLCVsHdoP3++7nYlNXov5zIZ3+4O5DBVMvir3WyKu/25pHCVazRgIu/TeDuz4G5Pt3sypnGY5bEXYO6P8bfpny2Nv7+/NzsyoW7RObx2JXEXIW7RYa8RrHThtjqw4K6P7LUifj79eQHI2gAAAA7dFJOUwAC/Pr7AwT+Af3sj6r42YddBboos/ZgaizBxm3zYhQX+fAV7Rl2d9jUPSeO0ZGQQI0pzvfn6LteiOu5+wORBQAABNlJREFUWMOtl2VwIzcUgOXEtnadxIFjZmYqPa03dq+uG7s9cy6cNHzMzFhmZmaGuzIzMzMzc/ujb9fOxbsrOTNNNZOMxqP37XvSQ0IsS3Jq/6YVTS4Z4yqgtMA15qDJRdMk/NEpka6X5MCDU6ZOnwGGNWP61CkIdnSJQPGRo8dqInKOTKm2oRS32mbs6JH6gSzLJpEhhQeijF2mRg2obMdfXIVDiGQTyzuJNGAmgF0G7pLtADMHSMQpVr9vPxSnIFwUEf36CsxA1QYNzCqeRgwcpJlqlZdIHkAOdLnwSJ522iJfnM/9fKmXo0R+sZkg2Zw9IZf/SYWZf8mFnsVGK9D58nnyXjjqxOWohWoh5Bvd0on2877PYE54Xs1K3JSaCXmZr+kgvfn6M6g9O+oPLb7HYkku9O58TRvp0SuHCgDnxysDid2r3t9lRNCcXj2IrSN8bP0F78eg7PDIak97ZTTx7fevmV6zvy0dWg7SB+wgBMT3PPP6pq/avvvhi6V7WYYOduiTMkKSBrtkWQwIBJ9e/vG26JZY4JM3IOMuZdk1WPcGByk0K6AyYJ0A/3W4+2lNOOb7PBOAQoWaCjYyfBilpu92HEwBjvA2436bZ/YhBgClw4ajuJOMMinA4KJVV18I5Z0AUOBH9VAzAMVGaRnA6aaGG1ChKdwSvMarGABesAJk6tbccRwYb1CBy4Kh6uBNuOkCgILj8BInULMFB0fDP1d4rkddugDY6QR8hIkmDRDgC/8W87xdW47PXnZSVg0mOsj4AqAWwO+R9mr/W5rvlp1cIQZQKBhPikwK6Cb8tWPn7rbgOlBU9UxfixCAovuTSeYwQIAn+fefDbHWQBNKLps3v0YIyIFJpISjgce3Hf4Ihnw7mlUFzjutJosG+xG36QpSgHr2z691oeACUBRYvEAIoOAmI7iAwwB+iQRiibvxImuPRq8UAUaQodYI1AG74LZgdeS+laCmsyMPADCUUAEAra9PhhLrmaInARGAigGMLdtZtWHTO5pLZwOITFDx8MvR8Ma6L3WC2ATRJapaWL3YsMHz0RKNIL5EtxjgVdTNLaFkvYL2iJ+R70g6AP8+DbVXf/jN1wDvekWOxHXlNIApsLWhKlQXffMVgA8ErlyURQNNh62fBcOV/icfL938Kj+YeOGcAjA4vZmpUPZeZEuotWHtS/EEN5x5CSUFUKDxFAxogBfu9UdCdZHK269K+YQxofBSWgeg5oSzQNWy67q5/tgj1fMvNgH0lGZJqhmAY447/lQv0+rMrXd6kpevOYOBNala0noG4MjZiaVM1TM1LLryEi8Y251UWrcUlhRAeVYD+CsWLtGrHFM5lTNVWCylrSMfqE8hIHDsnHSZVBXFHEep0mYprgxu9mxc+NyjuL3LH5jbAeApcIBe383lncEVN4QrkvHnH4CHglkA+8q7ucFg3kXbA/4HY77k2vVVcTFgX4PBa3GubVzhr2qrenh1RAjIaHEsTZaKV37HY/f7fBWV8RsFgMwmi9PmlWuV+Yn6wC2eFQKAoc3jNZpMe7Smxktbz+UCTI0mv9XVLLngnDIewNLqCprtcoX/ANhsO80jg6DdZwp35rC2+90fOLo/8nR/6EqPfbOyj32zso196cHT9d8Hz+6Pvv/D8M0d/0uyjP//AqaoJ1l31SbdAAAAAElFTkSuQmCC",
		Tag:         "NULS",
		Name:        "Nuls (NULS)",
		Trezor:      false,
		Ledger:      false,
		Segwit:      false,
		Masternodes: false,
		Token:       false,
		StableCoin:  false,
		Blockbook:   "https://api.nuls.io",
		Protocol:    "nuls",
		TxVersion:   1,
		TxBuilder:   "nuls",
		HDIndex:     8964,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18NULS Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x488b21e,
					Private: 0x488ade4,
				},
				PubKeyHash: []int{0x00},
				ScriptHash: []int{0x05},
				Wif:        []int{0x80},
			},
		},
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "",
	},
	BlockchainInfo: BlockchainInfo{
		BlockTime:        1,
		MinConfirmations: 1,
		ExternalSource:   "nulscan.io",
	},
	NetParams: &chaincfg.Params{
		AddressMagicLen:   1,
		Bech32HRPSegwit:   "",
		HDPrivateKeyID:    [4]byte{0x04, 0x88, 0xAD, 0xE4},
		HDPublicKeyID:     [4]byte{0x04, 0x88, 0xB2, 0x1E},
		HDCoinType:        8964,
		Base58CksumHasher: base58.Sha256D,
	},
}
