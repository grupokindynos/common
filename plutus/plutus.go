package plutus

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/grupokindynos/common/jwt"
	"github.com/grupokindynos/common/tokens/mrt"
)

// HTTPClient a usable client with hardcoded timeout
var HTTPClient = http.Client{
	Timeout: time.Second * 15,
}

//GetWalletTX gets a transaction info from the wallets
func GetWalletTX(productionURL string, coin string, tx string, signature string, service string, username string, password string, plutusKey string, masterPass string) (transaction Transaction, err error) {
	requestURL := productionURL + "/tx/" + coin + "/" + tx
	signMessage, _ := jwt.EncodeJWS(service, signature)

	res, header, err := getPlutusData(requestURL, username, password, signMessage)
	if err != nil {
		return transaction, err
	}

	payload, err := parsePlutusData(res, header, plutusKey, masterPass)
	if err != nil {
		return transaction, err
	}

	err = json.Unmarshal(payload, &transaction)

	if err != nil {
		return transaction, errors.New("unable to unmarshal response")
	}

	return transaction, err
}

//GetWalletBalance gets an address from a given coin
func GetWalletBalance(productionURL string, coin string, signature string, service string, username string, password string, plutusKey string, masterPass string) (balance Balance, err error) {
	requestURL := productionURL + "/balance/" + coin
	signMessage, _ := jwt.EncodeJWS(service, signature)

	res, header, err := getPlutusData(requestURL, username, password, signMessage)
	if err != nil {
		return balance, err
	}

	payload, err := parsePlutusData(res, header, plutusKey, masterPass)
	if err != nil {
		return balance, err
	}

	err = json.Unmarshal(payload, &balance)
	if err != nil {
		return balance, errors.New("unable to unmarshal response")
	}

	return balance, err
}

//GetWalletInfo gets an address from a given coin
func GetWalletInfo(productionURL string, coin string, signature string, service string, username string, password string, plutusKey string, masterPass string) (info Info, err error) {
	requestURL := productionURL + "/info/" + coin
	signMessage, _ := jwt.EncodeJWS(service, signature)

	res, header, err := getPlutusData(requestURL, username, password, signMessage)
	if err != nil {
		return info, err
	}

	payload, err := parsePlutusData(res, header, plutusKey, masterPass)
	if err != nil {
		return info, err
	}

	err = json.Unmarshal(payload, &info)
	if err != nil {
		return info, errors.New("unable to unmarshal response")
	}

	return info, err
}

//GetWalletAddress gets an address from a given coin
func GetWalletAddress(productionURL string, coin string, signature string, service string, username string, password string, plutusKey string, masterPass string) (address string, err error) {
	requestURL := productionURL + "/address/" + coin
	signMessage, _ := jwt.EncodeJWS(service, signature)

	res, header, err := getPlutusData(requestURL, username, password, signMessage)
	if err != nil {
		return address, err
	}

	payload, err := parsePlutusData(res, header, plutusKey, masterPass)
	if err != nil {
		return address, err
	}

	err = json.Unmarshal(payload, &address)
	if err != nil {
		return address, errors.New("unable to unmarshal response")
	}

	return address, err
}

//GetWalletStatus gets all the information from a given coin
func GetWalletStatus(productionURL string, coin string, signature string, service string, username string, password string, plutusKey string, masterPass string) (status Status, err error) {
	requestURL := productionURL + "/status/" + coin
	signMessage, _ := jwt.EncodeJWS(service, signature)

	res, header, err := getPlutusData(requestURL, username, password, signMessage)
	if err != nil {
		return status, err
	}

	payload, err := parsePlutusData(res, header, plutusKey, masterPass)
	if err != nil {
		return status, err
	}

	err = json.Unmarshal(payload, &status)
	if err != nil {
		return status, errors.New("unable to unmarshal response")
	}

	return status, err
}

func parsePlutusData(res []byte, header string, plutusKey string, masterPass string) (payload []byte, err error) {
	var resStr string
	err = json.Unmarshal(res, &resStr)
	if err != nil {
		return payload, errors.New("unable to unmarshal response")
	}

	valid, payload := mrt.VerifyMRTToken(header, resStr, plutusKey, masterPass)
	if !valid {
		return payload, errors.New("could not validate signature")
	}

	return payload, nil
}

func getPlutusData(requestURL string, username string, password string, signMessage string) (data []byte, header string, err error) {
	req, _ := http.NewRequest("GET", requestURL, nil)
	req.Header.Set("service", signMessage)
	req.SetBasicAuth(username, password)

	res, err := HTTPClient.Do(req)

	if err != nil {
		return data, header, errors.New("request timed out")
	}
	defer func() {
		_ = res.Body.Close()
	}()
	header = res.Header.Get("service")

	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return data, header, err
	}

	return data, header, err

}
