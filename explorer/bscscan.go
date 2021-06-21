package explorer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type BscScan struct {
	Url        string
	client     *http.Client
	bscScanUrl string
	bscApiUrl  string
}

// Methods for Bitcoin-like coins.

func (b *BscScan) GetAddress(address string) (response Address, err error) {
	panic("not implemented")
	//data, err := b.callWrapper("GET", "address/"+address, 2, nil)
	//if err != nil {
	//	return response, err
	//}
	//err = json.Unmarshal(data, &response)
	//if err != nil {
	//	return response, err
	//}
	//return
}

func (b *BscScan) GetXpub(xpub string) (response Xpub, err error) {
	panic("not implemented")
	//data, err := b.callWrapper("GET", "xpub/"+xpub+"?details=txs&gap=1000", 2, nil)
	//if err != nil {
	//	return response, err
	//}
	//err = json.Unmarshal(data, &response)
	//if err != nil {
	//	return response, err
	//}
	//return
}

func (b *BscScan) GetUtxo(xpub string, confirmed bool) (response []Utxo, err error) {
	panic("not implemented")
	//var url string
	//if confirmed {
	//	url = "utxo/" + xpub + "?confirmed=true&gap=1000"
	//} else {
	//	url = "utxo/" + xpub + "?gap=1000"
	//}
	//data, err := b.callWrapper("GET", url, 2, nil)
	//if err != nil {
	//	return response, err
	//}
	//err = json.Unmarshal(data, &response)
	//if err != nil {
	//	return response, err
	//}
	//return
}

func (b *BscScan) GetTx(txid string) (response Tx, err error) {
	data, err := b.callBscAPIWrapper("GET", "api/v1/tx/"+txid, nil)
	if err != nil {
		return response, err
	}
	var txInfo BSCTxInfoResponse
	err = json.Unmarshal(data, &txInfo)
	if err != nil {
		return response, err
	}
	response = Tx{
		BlockHash:      "",
		BlockHeight:    int(txInfo.Data.TxInfo.Block),
		BlockTime:      0,
		Confirmations:  int(txInfo.Data.TxInfo.Confirmations),
		Fees:           "",
		Hex:            "",
		LockTime:       0,
		Txid:           txInfo.Data.TxInfo.TransactionHash,
		Value:          "",
		ValueIn:        "",
		Version:        0,
		Vin:            nil,
		Vout:           nil,
		TokenTransfers: nil,
	}
	return
}

func (b *BscScan) GetFee(nBlocks string) (response Fee, err error) {
	panic("not implemented")
	//data, err := b.callWrapper("GET", "estimatefee/"+nBlocks, 1, nil)
	//if err != nil {
	//	return response, err
	//}
	//err = json.Unmarshal(data, &response)
	//if err != nil {
	//	return response, err
	//}
	//return
}

// Methods for Ethereum

func (b *BscScan) GetEthAddress(addr string) (response EthAddr, err error) {
	panic("not implemented")
	//data, err := b.callWrapper("GET", "address/"+addr+"?details=txs", 2, nil)
	//if err != nil {
	//	return response, err
	//}
	//err = json.Unmarshal(data, &response)
	//if err != nil {
	//	return response, err
	//}
	//return
}

func (b *BscScan) GetTxEth(txid string) (response EthTx, err error) {
	panic("not implemented")
	//data, err := b.callWrapper("GET", "tx/"+txid, 2, nil)
	//if err != nil {
	//	return response, err
	//}
	//err = json.Unmarshal(data, &response)
	//if err != nil {
	//	return response, err
	//}
	//return
}

// Methods for all coins

func (b *BscScan) SendTx(rawTx string) (response string, err error) {
	apikey := os.Getenv("BSC_SCAN_API_KEY") // Requires API Key to be set on target project
	// https://api.bscscan.com/api?module=proxy&action=eth_sendRawTransaction&hex=0xf8ab2c85012a05f20083030d4094b5bea8a26d587cf665f2d78f077cca3c7f6341bd80b844a9059cbb000000000000000000000000ea703e63ba6c9b5224969d6483327b8e65af76cc000000000000000000000000000000000000000000000000117a4bf9521c48008193a01409d51b4476ddad15c2a793b273104a6da0405e9abd09a4dd697de02318d801a03d80c428d5835632e06c327f5e524bc54cecc8248bd4229388cdaf0918097c4d&apikey=VGW5CG1EQQ2MK4C7DKAHNEH98SA9T7RNNA
	url := fmt.Sprintf("/api?module=proxy&action=eth_sendRawTransaction&hex=%s&apikey=%s", rawTx, apikey)
	data, err := b.callWrapper("POST", url, strings.NewReader(rawTx))
	if err != nil {
		return response, err
	}
	var bscScanAnswer BSCResponse
	err = json.Unmarshal(data, &bscScanAnswer)
	if err != nil || bscScanAnswer.Error.Message != "" {
		if err == nil {
			return response, errors.New(bscScanAnswer.Error.Message)
		}
		return response, err
	}
	if bscScanAnswer.Result != "" {
		return bscScanAnswer.Result, nil
	} else {
		return "", errors.New(bscScanAnswer.Error.Message)
	}
}

func (b *BscScan) SendTxWithMessage(rawTx string) (string, error, string) {
	log.Println("Sending tx through BSC Scan")
	apikey := os.Getenv("BSC_SCAN_API_KEY") // Requires API Key to be set on target project
	log.Println("API: ", apikey)
	// https://api.bscscan.com/api?module=proxy&action=eth_sendRawTransaction&hex=0xf8ab2c85012a05f20083030d4094b5bea8a26d587cf665f2d78f077cca3c7f6341bd80b844a9059cbb000000000000000000000000ea703e63ba6c9b5224969d6483327b8e65af76cc000000000000000000000000000000000000000000000000117a4bf9521c48008193a01409d51b4476ddad15c2a793b273104a6da0405e9abd09a4dd697de02318d801a03d80c428d5835632e06c327f5e524bc54cecc8248bd4229388cdaf0918097c4d&apikey=VGW5CG1EQQ2MK4C7DKAHNEH98SA9T7RNNA
	url := fmt.Sprintf("api?module=proxy&action=eth_sendRawTransaction&hex=%s&apikey=%s", rawTx, apikey)
	data, err := b.callWrapper("POST", url, strings.NewReader(rawTx))
	if err != nil {
		return "", err, string(data)
	}
	log.Println("data: ", data)
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

func (b *BscScan) callWrapper(method string, route string, body io.Reader) (response []byte, err error) {
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var request *http.Request
	url := fmt.Sprintf("%s/%s", b.bscScanUrl, route)
	log.Println("url: ", url)
	switch method {
	case "GET":
		request, err = http.NewRequest(method, url, nil)
	case "POST":
		request, err = http.NewRequest(method, url, body)
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

func (b *BscScan) callBscAPIWrapper(method string, route string, body io.Reader) (response []byte, err error) {
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var request *http.Request
	url := fmt.Sprintf("%s/%s", b.bscApiUrl, route)
	switch method {
	case "GET":
		request, err = http.NewRequest(method, url, nil)
	case "POST":
		request, err = http.NewRequest(method, url, body)
	}
	request.Header.Add("api-key", os.Getenv("PANCAKE_PASSWORD"))
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

func NewBscScanWrapper(url string) *BscScan {
	bb := &BscScan{
		Url: url,
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       30 * time.Second,
		},
		bscApiUrl: "https://pp-bsc-api.herokuapp.com",
		// bscScanUrl: "https://pp-bsc-api.herokuapp.com",
		bscScanUrl: "https://api.bscscan.com",
	}
	return bb
}

// utils
func (b *BscScan) FindDepositTxId(address string, amount int64) (string, error) {
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
