package explorer

import (
	"errors"
	coinfactory "github.com/grupokindynos/common/coin-factory"
	"github.com/grupokindynos/common/coin-factory/coins"
	"strings"
)

type ExplorerFactory struct {
}

func NewExplorerFactory() *ExplorerFactory {
	exFactory := new(ExplorerFactory)
	return exFactory
}

func (e *ExplorerFactory) createInstance(name string, url string) (Explorer, error) {
	var exName = strings.ToLower(name)
	if exName == "blockbook" {
		return NewBlockBookWrapper(url), nil
	} else if exName == "nuls" {
		return NewNulsWrapper(url), nil
	} else {
		return nil, errors.New("Explorer not found")
	}
}

func (e *ExplorerFactory) GetExplorerByCoin(coin coins.Coin) (Explorer, error) {
	coinInfo, _ := coinfactory.GetCoin(coin.Info.Tag)
	if coinInfo.Info.Protocol == "nuls" {
		return e.createInstance("nuls", coinInfo.Info.Blockbook)
	}
	return e.createInstance("blockbook", coinInfo.Info.Blockbook)
}
