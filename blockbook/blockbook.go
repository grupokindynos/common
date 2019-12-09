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
	data, err := b.callWrapper("xpub/" + xpub + "?details=txs")
	if err != nil {
		return response, nil
	}
	err = json.Unmarshal(data, &response)
	return
}

func (b *BlockBook) GetUtxo(xpub string) (response []Utxo, err error) {
	data, err := b.callWrapper("utxo/" + xpub + "?confirmed=true")
	if err != nil {
		return response, nil
	}
	err = json.Unmarshal(data, &response)
	return
}

func (b *BlockBook) GetTx(txid string) (response Tx, err error) {
	data, err := b.callWrapper("tx/" + txid)
	if err != nil {
		return response, nil
	}
	err = json.Unmarshal(data, &response)
	return
}

// Methods for Ethereum

func (b *BlockBook) GetEthAddress(addr string) (response EthAddr, err error) {
	data, err := b.callWrapper("address/" + addr + "?details=txs")
	if err != nil {
		return response, nil
	}
	err = json.Unmarshal(data, &response)
	return
}

func (b *BlockBook) GetTxEth(txid string) (response EthTx, err error) {
	data, err := b.callWrapper("tx/" + txid)
	if err != nil {
		return response, nil
	}
	err = json.Unmarshal(data, &response)
	return
}

// Methods for all coins

func (b *BlockBook) SendTx(rawTx string) (response string, err error) {
	data, err := b.callWrapper("sendtx/" + rawTx)
	if err != nil {
		return response, nil
	}
	var blockbookAnswer SendTx
	err = json.Unmarshal(data, &blockbookAnswer)
	if blockbookAnswer.Result != "" {
		return blockbookAnswer.Result, nil
	} else {
		return "", errors.New(blockbookAnswer.Error.Message)
	}
}

func (b *BlockBook) checkBlockStatus() bool {
	data, err := b.callWrapper("status")
	if err != nil {
		return false
	}
	var StatusInfo Status
	err = json.Unmarshal(data, &StatusInfo)
	if err != nil {
		return false
	}
	if !StatusInfo.Blockbook.InitialSync {
		return true
	}
	return false
}

func (b *BlockBook) callWrapper(route string) ([]byte, error) {
	resp, err := http.Get(b.Url + "/api/v2/" + route)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func NewBlockBookWrapper(url string) (*BlockBook, error) {
	bb := &BlockBook{
		Url: url,
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       30,
		},
	}
	if bb.checkBlockStatus() {
		return bb, nil
	} else {
		return nil, errors.New("blockbook not available or not synced")
	}
}
