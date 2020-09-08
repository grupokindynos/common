package hestia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/grupokindynos/common/tokens/mrt"
	"github.com/grupokindynos/common/tokens/mvt"
)

// ProductionURL is the current hestia production url
var ProductionURL = "https://hestia.polispay.com"

// HttpClient a usable client with hardcoded timeout
var HttpClient = http.Client{
	Timeout: time.Second * 5,
}

// GetServiceProperties is a function to return hestia properties for multiple services
func GetServiceProperties(productionURL string, adminFbToken string) (string, error) {
	res, err := HttpClient.Get(productionURL + "/config")
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
func GetCoinsAvailability(productionURL string, adminFbToken string) (string, error) {
	res, err := HttpClient.Get(productionURL + "/coins")
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

// VerifyToken ask Hestia if a firebase token is valid
func VerifyToken(productionURL string, service string, masterPassword string, fbToken string, hestiaAuthUser string, hestiaAuthPass string, servicePrivKey string, hestiaPubKey string) (valid bool, uid string, err error) {
	req, err := mvt.CreateMVTToken("POST", productionURL+"validate/token", service, masterPassword, fbToken, hestiaAuthUser, hestiaAuthPass, servicePrivKey)
	client := http.Client{
		Timeout: time.Second * 5,
	}
	res, err := client.Do(req)
	if err != nil {
		return false, "", err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	headerSignature := res.Header.Get("service")
	if headerSignature == "" {
		return false, "", err
	}
	contents, _ := ioutil.ReadAll(res.Body)
	var contentStr string
	err = json.Unmarshal(contents, &contentStr)
	if err != nil {
		return false, "", err
	}
	valid, data := mrt.VerifyMRTToken(headerSignature, contentStr, hestiaPubKey, masterPassword)
	var hestiaTokenRes TokenVerification
	err = json.Unmarshal(data, &hestiaTokenRes)
	if err != nil {
		return false, "", err
	}
	return hestiaTokenRes.Valid, hestiaTokenRes.UID, err
}
