package blockbook

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type BlockBook struct {
	Url    string
	client *http.Client
}

// Methods for Bitcoin-like coins.

func (b *BlockBook) GetXpub(xpub string) (response Xpub, err error) {
	data, err := b.callWrapper("xpub/"+xpub+"?details=txs", 2)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return
}

func (b *BlockBook) GetUtxo(xpub string, confirmed bool) (response []Utxo, err error) {
	var url string
	if confirmed {
		url = "utxo/" + xpub + "?confirmed=true"
	} else {
		url = "utxo/" + xpub
	}
	data, err := b.callWrapper(url, 2)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return
}

func (b *BlockBook) GetTx(txid string) (response Tx, err error) {
	data, err := b.callWrapper("tx/"+txid, 2)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return
}

func (b *BlockBook) GetFee(nBlocks string) (response Fee, err error) {
	data, err := b.callWrapper("estimatefee/"+nBlocks, 1)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return
}

// Methods for Ethereum

func (b *BlockBook) GetEthAddress(addr string) (response EthAddr, err error) {
	data, err := b.callWrapper("address/"+addr+"?details=txs", 2)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return
}

func (b *BlockBook) GetTxEth(txid string) (response EthTx, err error) {
	data, err := b.callWrapper("tx/"+txid, 2)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return
}

// Methods for all coins

func (b *BlockBook) SendTx(rawTx string) (response string, err error) {
	data, err := b.callWrapper("sendtx/"+rawTx, 2)
	if err != nil {
		return response, err
	}
	var blockbookAnswer SendTx
	err = json.Unmarshal(data, &blockbookAnswer)
	if err != nil {
		return response, err
	}
	if blockbookAnswer.Result != "" {
		return blockbookAnswer.Result, nil
	} else {
		return "", errors.New(blockbookAnswer.Error.Message)
	}
}

func (b *BlockBook) callWrapper(route string, version int) ([]byte, error) {
	var versionStr string
	switch version {
	case 1:
		versionStr = "/api/v1/"
	case 2:
		versionStr = "/api/v2/"
	default:
		versionStr = "/api/"
	}
	resp, err := http.Get(b.Url + versionStr + route)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func NewBlockBookWrapper(url string) *BlockBook {
	bb := &BlockBook{
		Url: url,
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       30,
		},
	}
	return bb
}
