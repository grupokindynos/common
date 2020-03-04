package blockbook

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BlockBook struct {
	Url    string
	client *http.Client
}

// Methods for Bitcoin-like coins.

func (b *BlockBook) GetAddress(address string) (response Address, err error) {
	data, err := b.callWrapper("GET", "address/"+address, 2, nil)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return
}

func (b *BlockBook) GetXpub(xpub string) (response Xpub, err error) {
	data, err := b.callWrapper("GET", "xpub/"+xpub+"?details=txs", 2, nil)
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
	data, err := b.callWrapper("GET", url, 2, nil)
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
	data, err := b.callWrapper("GET", "tx/"+txid, 2, nil)
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
	data, err := b.callWrapper("GET", "estimatefee/"+nBlocks, 1, nil)
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
	data, err := b.callWrapper("GET", "address/"+addr+"?details=txs", 2, nil)
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
	data, err := b.callWrapper("GET", "tx/"+txid, 2, nil)
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
	data, err := b.callWrapper("POST", "sendtx/", 2, strings.NewReader(rawTx))
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
		return "", errors.New(blockbookAnswer.Error)
	}
}

func (b *BlockBook) SendTxWithMessage(rawTx string) (string, error, string) {
	data, err := b.callWrapper("POST", "sendtx/", 2, strings.NewReader(rawTx))
	if err != nil {
		return "", err, string(data)
	}
	var blockbookAnswer SendTx
	err = json.Unmarshal(data, &blockbookAnswer)
	if err != nil {
		return "", err, string(data)
	}
	if blockbookAnswer.Result != "" {
		return blockbookAnswer.Result, nil, string(data)
	} else {
		return "", errors.New(blockbookAnswer.Error), string(data)
	}
}

func (b *BlockBook) callWrapper(method string, route string, version int, body io.Reader) (response []byte, err error) {
	var versionStr string
	switch version {
	case 1:
		versionStr = "/api/v1/"
	case 2:
		versionStr = "/api/v2/"
	default:
		versionStr = "/api/"
	}
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var request *http.Request
	switch method {
	case "GET":
		request, err = http.NewRequest(method, b.Url+versionStr+route, nil)
	case "POST":
		request, err = http.NewRequest(method, b.Url+versionStr+route, body)
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		resp.Body.Close()
	}()
	response, err = ioutil.ReadAll(resp.Body)
	return response, err
}

func NewBlockBookWrapper(url string) *BlockBook {
	bb := &BlockBook{
		Url: url,
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       30 * time.Second,
		},
	}
	return bb
}

// utils
func (b *BlockBook) FindDepositTxId(address string, amount int64) (string, error) {
	addressData, err := b.GetAddress(address)
	if err != nil {
		return "", err
	}
	for _, txId := range addressData.Txids {
		tx, err := b.GetTx(txId)
		if err != nil {
			return "", err
		}
		for _, vout := range tx.Vout {
			voutValue, _ := strconv.ParseInt(vout.Value, 10, 64)
			if amount != voutValue {continue}
			for _, voutAddress := range vout.Addresses {
				if voutAddress == address {
					return txId, nil
				}
			}
		}
	}
	return "", errors.New("txid not found")
}
