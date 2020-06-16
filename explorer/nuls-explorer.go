package explorer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type NulsExplorer struct {
	Url    string
	client *http.Client
	rpcUrl string
}

// api models
type NulsInfo struct {
	Success bool            `json:"success"`
	Data    NulsBalanceData `json:"data"`
}

type NulsError struct {
	Success bool          `json:"success"`
	Data    NulsErrorData `json:"data"`
}

type NulsBalanceData struct {
	Total         string `json:"total"`
	Freeze        string `json:"freeze"`
	Available     string `json:"available"`
	TimeLock      string `json:"timeLock"`
	ConsensusLock string `json:"consensusLock"`
	Nonce         string `json:"nonce"`
	NonceType     int    `json:"nonceType"`
}

type NulsErrorData struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

type NulsInfoModel struct {
	AssetChanId int `json:"assetChainId"`
	AssetId     int `json:"assetId"`
}

func NewNulsWrapper(url string) *NulsExplorer {
	ne := &NulsExplorer{
		Url: url,
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       30 * time.Second,
		},
		rpcUrl: "https://public1.nuls.io",
	}
	return ne
}

// Methods for Bitcoin-like coins.

func (b *NulsExplorer) GetAddress(address string) (response Address, err error) {
	fmt.Println("como")

	body := &NulsInfoModel{
		AssetChanId: 1,
		AssetId:     1,
	}
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(body)
	var nulsInfo NulsInfo
	var nulsError NulsError
	nulsdata, err := b.callWrapper("POST", "/api/accountledger/balance/"+address, 1, buf)
	if err != nil {
		return response, err
	}
	fmt.Println("estas")

	err = json.Unmarshal(nulsdata, &nulsInfo)
	if err != nil {
		return response, err
	}
	if !nulsInfo.Success {
		err = json.Unmarshal(nulsdata, &nulsError)
		if err != nil {
			return response, err
		}
		return response, errors.New(nulsError.Data.Message)
	}
	s, _ := json.MarshalIndent(nulsInfo, "", "\t")
	fmt.Print(string(s))
	response = Address{
		Page:               1,
		TotalPages:         1,
		ItemsOnPage:        1,
		Address:            address,
		Balance:            nulsInfo.Data.Available,
		TotalReceived:      "",
		TotalSent:          "",
		UnconfirmedBalance: nulsInfo.Data.Freeze,
		UnconfirmedTxs:     1,
		Txs:                1,
		Txids:              []string{},
	}
	return
}

func (b *NulsExplorer) GetXpub(xpub string) (response Xpub, err error) {
	data, err := b.callWrapper("GET", "xpub/"+xpub+"?details=txs&gap=1000", 2, nil)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return
}

func (b *NulsExplorer) GetUtxo(xpub string, confirmed bool) (response []Utxo, err error) {
	var url string
	if confirmed {
		url = "utxo/" + xpub + "?confirmed=true&gap=1000"
	} else {
		url = "utxo/" + xpub + "?gap=1000"
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

func (b *NulsExplorer) GetTx(txid string) (response Tx, err error) {
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

func (b *NulsExplorer) GetFee(nBlocks string) (response Fee, err error) {
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

func (b *NulsExplorer) GetEthAddress(addr string) (response EthAddr, err error) {
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

func (b *NulsExplorer) GetTxEth(txid string) (response EthTx, err error) {
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

func (b *NulsExplorer) SendTx(rawTx string) (response string, err error) {
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

func (b *NulsExplorer) SendTxWithMessage(rawTx string) (string, error, string) {
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

func (b *NulsExplorer) callWrapper(method string, route string, version int, body io.Reader) (response []byte, err error) {
	fmt.Println(b.Url + route)
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var request *http.Request
	switch method {
	case "GET":
		request, err = http.NewRequest(method, b.Url+route, body)
	case "POST":
		request, err = http.NewRequest(method, b.Url+route, body)
	}
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
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

// utils
func (b *NulsExplorer) FindDepositTxId(address string, amount int64) (string, error) {
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
			if amount != voutValue {
				continue
			}
			for _, voutAddress := range vout.Addresses {
				if voutAddress == address {
					return txId, nil
				}
			}
		}
	}
	return "", errors.New("txid not found")
}
