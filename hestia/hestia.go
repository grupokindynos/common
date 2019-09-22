package hestia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)
// ProductionURL is the current hestia production url
var ProductionURL = "https://hestia.polispay.com"

// HttpClient a usable client with hardcoded timeout
var HttpClient = http.Client{
	Timeout: time.Second * 5,
}

// GetServiceProperties is a function to return hestia properties for multiple services
func GetServiceProperties(adminFbToken string) (string, error) {
	res, err := HttpClient.Get(ProductionURL + "/config" )
	if err != nil {
	}
	defer func() {
		_ = res.Body.Close()
	}()
	contents, _ := ioutil.ReadAll(res.Body)
	var response Response
	err = json.Unmarshal(contents, &response)
	if err != nil {
		return "", err
	}
	var jwe string
	jweBytes, err := json.Marshal(response.Data)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(jweBytes, &jwe)
	if err != nil {
		return "", err
	}
	return jwe, nil
}

// GetCoinsAvailability is a function to return hestia properties for multiple crypto used on the environment
func GetCoinsAvailability(adminFbToken string) (string, error) {
	res, err := HttpClient.Get(ProductionURL + "/coins" )
	if err != nil {
	}
	defer func() {
		_ = res.Body.Close()
	}()
	contents, _ := ioutil.ReadAll(res.Body)
	var response Response
	err = json.Unmarshal(contents, &response)
	if err != nil {
		return "", err
	}
	var jwe string
	jweBytes, err := json.Marshal(response.Data)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(jweBytes, &jwe)
	if err != nil {
		return "", err
	}
	return jwe, nil
}
