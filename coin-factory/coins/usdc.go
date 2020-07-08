package coins

// USDCoin coinfactory information
var USDCoin = Coin{
	Info: CoinInfo{
		Icon:         "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAIAAAACACAYAAADDPmHLAAAAAXNSR0IArs4c6QAAAKhlWElmTU0AKgAAAAgABQESAAMAAAABAAEAAAEaAAUAAAABAAAASgEbAAUAAAABAAAAUgExAAIAAAAkAAAAWodpAAQAAAABAAAAfgAAAAAAAABIAAAAAQAAAEgAAAABQWRvYmUgUGhvdG9zaG9wIENDIDIwMTkgKE1hY2ludG9zaCkAAAOgAQADAAAAAQABAACgAgAEAAAAAQAAAICgAwAEAAAAAQAAAIAAAAAArScmrgAAAAlwSFlzAAALEwAACxMBAJqcGAAABJhpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4KICAgPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICAgICAgPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIKICAgICAgICAgICAgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iCiAgICAgICAgICAgIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIgogICAgICAgICAgICB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iCiAgICAgICAgICAgIHhtbG5zOnRpZmY9Imh0dHA6Ly9ucy5hZG9iZS5jb20vdGlmZi8xLjAvIj4KICAgICAgICAgPHhtcE1NOkluc3RhbmNlSUQ+eG1wLmlpZDpEMUZGRDE2N0YwQjAxMUU4QjBCOUNGMjkwNTQyMTUxQzwveG1wTU06SW5zdGFuY2VJRD4KICAgICAgICAgPHhtcE1NOkRvY3VtZW50SUQ+eG1wLmRpZDpEMUZGRDE2OEYwQjAxMUU4QjBCOUNGMjkwNTQyMTUxQzwveG1wTU06RG9jdW1lbnRJRD4KICAgICAgICAgPHhtcE1NOk9yaWdpbmFsRG9jdW1lbnRJRD54bXAuZGlk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",
		Tag:          "USDC",
		Name:         "USDCoin (USDC)",
		Trezor:       true,
		Ledger:       true,
		Segwit:       false,
		Masternodes:  false,
		Token:        true,
		StableCoin:   true,
		TokenNetwork: "ethereum",
		Contract:     "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		Decimals:     6,
		Blockbook:    "https://eth2.trezor.io",
	},
	Rates: RatesSource{
		Exchange:         "binance",
		FallBackExchange: "crex24",
	},
}
