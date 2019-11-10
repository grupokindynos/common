package obol

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// ProductionURL is the current obol production url
var ProductionURL = "https://obol.polispay.com"

// HttpClient a usable client with hardcoded timeout
var HttpClient = http.Client{
	Timeout: time.Second * 5,
}

// GetCoinRates is a function to return obol simple rates for a coin
func GetCoinRates(obolURL string, coin string) ([]Rate, error) {
	res, err := HttpClient.Get(obolURL + "/simple/" + coin)
	if err != nil {
		return []Rate{}, err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	contents, _ := ioutil.ReadAll(res.Body)
	var response Response
	err = json.Unmarshal(contents, &response)
	if err != nil {
		return []Rate{}, err
	}
	var rates []Rate
	rateBytes, err := json.Marshal(response.Data)
	if err != nil {
		return []Rate{}, err
	}
	err = json.Unmarshal(rateBytes, &rates)
	if err != nil {
		return []Rate{}, err
	}
	return rates, nil
}

// GetCoin2CoinRates is a function to return obol complex rates between 2 coins
func GetCoin2CoinRates(obolURL string, fromcoin string, tocoin string) (float64, error) {
	res, err := HttpClient.Get(obolURL + "/complex/" + fromcoin + "/" + tocoin)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	contents, _ := ioutil.ReadAll(res.Body)
	var response Response
	err = json.Unmarshal(contents, &response)
	if err != nil {
		return 0, err
	}
	var rate float64
	rateByte, err := json.Marshal(response.Data)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(rateByte, &rate)
	if err != nil {
		return 0, err
	}
	return rate, nil
}

// GetCoin2CoinRatesWithAmount is a function to return obol complex rates between 2 coins with a hardcoded amount
func GetCoin2CoinRatesWithAmount(obolURL string, fromcoin string, tocoin string, amount string) (CoinToCoinWithAmountResponse, error) {
	res, err := HttpClient.Get(obolURL + "/complex/" + fromcoin + "/" + tocoin + "?amount=" + amount)
	if err != nil {
		return CoinToCoinWithAmountResponse{}, err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	contents, _ := ioutil.ReadAll(res.Body)
	var response Response
	err = json.Unmarshal(contents, &response)
	if err != nil {
		return CoinToCoinWithAmountResponse{}, err
	}
	var rate CoinToCoinWithAmountResponse
	rateByte, err := json.Marshal(response.Data)
	if err != nil {
		return CoinToCoinWithAmountResponse{}, err
	}
	err = json.Unmarshal(rateByte, &rate)
	if err != nil {
		return CoinToCoinWithAmountResponse{}, err
	}
	return rate, nil
}
