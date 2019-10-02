package plutus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/grupokindynos/common/jwt"
	"github.com/grupokindynos/common/tokens/mrt"
)

// HttpClient a usable client with hardcoded timeout
var HttpClient = http.Client{
	Timeout: time.Second * 15,
}

//GetWalletAddress gets a deposit address from a given coin
func GetWalletAddress(productionURL string, coin string, signature string, service string, username string, password string, plutusKey string, masterPass string) (address string, err error) {
	requestURL := productionURL + "/address/" + coin
	signMessage, err := jwt.EncodeJWS(service, signature)

	res, header, err := getPlutusData(requestURL, username, password, signMessage)

	var resStr string
	err = json.Unmarshal(res, &resStr)
	if err != nil {
		return "", errors.New("unable to unmarshal response")
	}

	valid, payload := mrt.VerifyMRTToken(header, resStr, plutusKey, masterPass)
	if !valid {
		return "", errors.New("could not validate signature")
	}

	err = json.Unmarshal(payload, &address)
	if err != nil {
		return "", errors.New("unable to unmarshal response")
	}

	return address, err
}

//GetWalletInfo gets all the information from a given coin
func GetWalletInfo(productionURL string, coin string, signature string, service string, username string, password string, plutusKey string, masterPass string) (address string, err error) {
	requestURL := productionURL + "/status/" + coin
	signMessage, err := jwt.EncodeJWS(service, signature)

	res, header, err := getPlutusData(requestURL, username, password, signMessage)

	var resStr string
	err = json.Unmarshal(res, &resStr)
	if err != nil {
		return "", errors.New("unable to unmarshal response")
	}

	valid, payload := mrt.VerifyMRTToken(header, resStr, plutusKey, masterPass)
	if !valid {
		return "", errors.New("could not validate signature")
	}

	fmt.Println(string(payload))
	/*
		err = json.Unmarshal(payload, &address)
		if err != nil {
			return "", errors.New("unable to unmarshal response")
		}
	*/

	return address, err
}

//GetPlutusData makes a GET request to the plutus API and returns the data as a json array
func getPlutusData(requestURL string, username string, password string, signMessage string) (data []byte, header string, err error) {
	req, _ := http.NewRequest("GET", requestURL, nil)
	req.Header.Set("service", signMessage)
	req.SetBasicAuth(username, password)

	res, err := HttpClient.Do(req)

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
