package explorer

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
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

type RpcTxResponse struct {
	JsonRpc string         `json:"jsonrpc"`
	Id      string         `json:"id"`
	Result  TxListResponse `json:"result"`
}

type TxListResponse struct {
	PageNumber int      `json:"pageNumber"`
	PageSize   int      `json:"pageSize"`
	TotalCount int      `json:"totalCount"`
	List       []NulsTx `json:"list"`
}

type NulsTx struct {
	TxHash       string  `json:"txHash"`
	Address      string  `json:"address"`
	Type         int     `json:"type"`
	CreateTime   int     `json:"createTime"`
	Height       int     `json:"height"`
	ChainId      int     `json:"chainId"`
	AssetId      int     `json:"assetId"`
	Symbol       string  `json:"symbol"`
	Values       int     `json:"values"`
	Fee          NulsFee `json:"fee"`
	Balance      int     `json:"balance"`
	TransferType int     `json:"transferType"`
	Status       int     `json:"status"`
}

type NulsFee struct {
	ChainId int    `json:"chainId"`
	AssetId int    `json:"assetId"`
	Symbol  string `json:"symbol"`
	Value   int    `json:"value"`
}

type RpcRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type NulsTxResponse struct {
	Success bool         `json:"success"`
	Data    NulsTxDetail `json:"data"`
}

type NulsFrom struct {
	Address       string `json:"address"`
	AssetsChainId int    `json:"assetsChainId"`
	AssetsId      int    `json:"assetsId"`
	Amount        string `json:"amount"`
	Nonce         string `json:"nonce"`
	Locked        int    `json:"locked"`
}

type NulsTo struct {
	Address       string `json:"address"`
	AssetsChainId int    `json:"assetsChainId"`
	AssetsId      int    `json:"assetsId"`
	Amount        string `json:"amount"`
	LockTime      int    `json:"lockTime"`
}

type NulsTxDetail struct {
	Hash                 string     `json:"hash"`
	Type                 int        `json:"type"`
	Time                 string     `json:"time"`
	Timestamp            int        `json:"timestamp"`
	BlockHeight          int        `json:"blockHeight"`
	BlockHash            string     `json:"blockHash"`
	Remark               string     `json:"remark"`
	TransactionSignature string     `json:"transactionSignature"`
	TxDataHex            string     `json:"txDataHex"`
	Status               int        `json:"status"`
	Size                 int        `json:"size"`
	InBlockIndex         int        `json:"inBlockIndex"`
	From                 []NulsFrom `json:"from"`
	To                   []NulsTo   `json:"to"`
}

type NulsCalculatedFee struct {
	Success bool `json:"success"`
	Data    struct {
		Value string `json:"value"`
	} `json:"data"`
}

type NulsFeeRequest struct {
	AddressCount int         `json:"addressCount"`
	FromLength   int         `json:"fromLength"`
	ToLength     int         `json:"toLength"`
	Remark       interface{} `json:"remark"`
	Price        interface{} `json:"price"`
}

type BroadcastRequest struct {
	TxHex string `json:"txHex"`
}

type NulsBroadcast struct {
	Success bool `json:"success"`
	Data    struct {
		Value bool   `json:"value"`
		Hash  string `json:"hash"`
	} `json:"data"`
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

func (b *NulsExplorer) GetAddress(address string) (response Address, err error) {
	body := &NulsInfoModel{
		AssetChanId: 1,
		AssetId:     1,
	}
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(body)
	var nulsInfo NulsInfo
	nulsdata, err := b.callWrapper("POST", "/api/accountledger/balance/"+address, 1, buf)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(nulsdata, &nulsInfo)
	if err != nil {
		return response, err
	}
	if !nulsInfo.Success {
		return response, handleNulsError(nulsdata)
	}
	// get the tx history
	var rpcresponse RpcTxResponse
	body2 := &RpcRequest{
		Jsonrpc: "2.0",
		Method:  "getAcctTxs",
		Params:  []interface{}{1, address, 0, 1, 1, -1, -1, 1, 1},
		Id:      1234,
	}
	buf = new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(body2)
	nulsdata, err = b.callWrapper("POST", address, 2, buf)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(nulsdata, &rpcresponse)
	if err != nil {
		return response, err
	}

	cont := rpcresponse.Result.TotalCount
	var txList []NulsTx
	var txIds []string
	pNumber := 1
	for cont > 0 {
		pSize := 100
		if cont < 100 {
			pSize = cont
		}
		body2 = &RpcRequest{
			Jsonrpc: "2.0",
			Method:  "getAcctTxs",
			Params:  []interface{}{1, address, 0, 1, 1, -1, -1, pNumber, pSize},
			Id:      1234,
		}
		buf = new(bytes.Buffer)
		_ = json.NewEncoder(buf).Encode(body2)
		nulsdata, err = b.callWrapper("POST", address, 2, buf)
		if err != nil {
			return response, err
		}
		err = json.Unmarshal(nulsdata, &rpcresponse)
		if err != nil {
			return response, err
		}
		txList = append(txList, rpcresponse.Result.List...)
		cont -= pSize
		pNumber += 1
	}
	totalReceived := 0
	totalSent := 0
	unconfirmedTxs := 0
	for _, element := range txList {
		txIds = append(txIds, element.TxHash)
		if element.Status == 0 {
			unconfirmedTxs += 1
			continue
		}
		if element.TransferType == 1 {
			totalReceived += element.Values
		} else if element.TransferType == -1 {
			totalSent += element.Values
		}
	}
	response = Address{
		Page:               1,
		TotalPages:         1,
		ItemsOnPage:        cont,
		Address:            address,
		Balance:            nulsInfo.Data.Available,
		TotalReceived:      strconv.Itoa(totalReceived),
		TotalSent:          strconv.Itoa(totalSent),
		UnconfirmedBalance: nulsInfo.Data.Freeze,
		UnconfirmedTxs:     unconfirmedTxs,
		Txs:                len(txIds),
		Txids:              txIds,
	}
	return
}

func handleNulsError(data []byte) (err error) {
	var nulsError NulsError
	err = json.Unmarshal(data, &nulsError)
	if err != nil {
		return err
	}
	return errors.New(nulsError.Data.Message)
}

func (b *NulsExplorer) GetXpub(xpub string) (response Xpub, err error) {
	return response, errors.New("method not available for this coin type")
}

func (b *NulsExplorer) GetUtxo(xpub string, confirmed bool) (response []Utxo, err error) {
	return response, errors.New("method not available for this coin type")

}

func (b *NulsExplorer) GetTx(txid string) (response Tx, err error) {
	var tx NulsTxResponse
	data, err := b.callWrapper("GET", "/api/tx/"+txid, 1, nil)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &tx)
	if err != nil {
		return response, err
	}
	if !tx.Success {
		return response, handleNulsError(data)
	}
	var Vins []TxVin
	var Vouts []TxVout

	for _, fromTx := range tx.Data.From {
		vin := TxVin{
			Addresses: []string{fromTx.Address},
			Hex:       fromTx.Nonce,
			N:         fromTx.AssetsId,
			Sequence:  fromTx.AssetsChainId,
			Txid:      tx.Data.Hash,
			Value:     fromTx.Amount,
			Vout:      len(tx.Data.To),
		}
		Vins = append(Vins, vin)
	}
	for _, toTx := range tx.Data.To {
		vout := TxVout{
			Value:     toTx.Amount,
			N:         toTx.AssetsId,
			Spent:     true,
			Hex:       tx.Data.Hash,
			Addresses: []string{toTx.Address},
		}
		Vouts = append(Vouts, vout)
	}
	value := 0
	inputValue, _ := strconv.Atoi(tx.Data.From[0].Amount)
	outputValue, _ := strconv.Atoi(tx.Data.To[0].Amount)
	if len(tx.Data.To) > 0 {
		value = outputValue
	} else {
		value = inputValue
	}
	response = Tx{
		BlockHash:     tx.Data.BlockHash,
		BlockHeight:   tx.Data.BlockHeight,
		BlockTime:     tx.Data.Timestamp,
		Confirmations: 1,
		Fees:          strconv.Itoa(inputValue - outputValue),
		Hex:           tx.Data.TxDataHex,
		LockTime:      tx.Data.To[0].LockTime,
		Txid:          tx.Data.Hash,
		Value:         strconv.Itoa(value),
		ValueIn:       tx.Data.From[0].Amount,
		Version:       tx.Data.Type,
		Vin:           Vins,
		Vout:          Vouts,
	}
	return
}

func (b *NulsExplorer) GetFee(nBlocks string) (response Fee, err error) {
	body := &NulsFeeRequest{
		AddressCount: 1,
		FromLength:   1,
		ToLength:     1,
		Remark:       nil,
		Price:        nil,
	}
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(body)
	var nulsFee NulsCalculatedFee
	nulsdata, err := b.callWrapper("POST", "/api/accountledger/calcTransferTxFee", 1, buf)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(nulsdata, &nulsFee)
	if err != nil {
		return response, err
	}
	if !nulsFee.Success {
		return response, handleNulsError(nulsdata)
	}
	response = Fee{
		Result: nulsFee.Data.Value,
	}
	return
}

// Methods for Ethereum

func (b *NulsExplorer) GetEthAddress(addr string) (response EthAddr, err error) {
	return response, errors.New("method not available for this coin type")

}

func (b *NulsExplorer) GetTxEth(txid string) (response EthTx, err error) {
	return response, errors.New("method not available for this coin type")
}

// Methods for all coins

func (b *NulsExplorer) SendTx(rawTx string) (response string, err error) {
	body := &BroadcastRequest{
		TxHex: rawTx,
	}
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(body)
	var nulsRes NulsBroadcast
	nulsdata, err := b.callWrapper("POST", "/api/accountledger/transaction/broadcast", 1, buf)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(nulsdata, &nulsRes)
	if err != nil {
		return response, err
	}
	if !nulsRes.Success {
		return response, handleNulsError(nulsdata)
	}
	return nulsRes.Data.Hash, nil

}

func (b *NulsExplorer) SendTxWithMessage(rawTx string) (string, error, string) {
	txid, err := b.SendTx(rawTx)
	return txid, err, ""
}

func (b *NulsExplorer) callWrapper(method string, route string, version int, body io.Reader) (response []byte, err error) {
	var finalRoute string
	switch version {
	case 1:
		finalRoute = b.Url + route
	case 2:
		finalRoute = b.rpcUrl
	default:
		finalRoute = b.Url + route
	}
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var request *http.Request
	switch method {
	case "GET":
		request, err = http.NewRequest(method, finalRoute, body)
	case "POST":
		request, err = http.NewRequest(method, finalRoute, body)
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
