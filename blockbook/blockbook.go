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

func (b *BlockBook) GetXpub(xpub string) {

}

func (b *BlockBook) GetUtxo(addr string) {

}

func (b *BlockBook) GetTx(txid string) {

}

// Methods for Ethereum

func (b *BlockBook) GetEthAddress(addr string) {

}

func (b *BlockBook) GetTxEth(txid string) {

}

// Methods for all coins

func (b *BlockBook) SendTx(txid string) {

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
