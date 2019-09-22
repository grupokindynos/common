package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/grupokindynos/common/aes"
	coinfactory "github.com/grupokindynos/common/coin-factory"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type EnvironmentVars struct {
	HerokuUsername  string
	HerokuPassword  string
	AuthUsername    string
	AuthPassword    string
	GinMode         string
	KeyPassword     string
	TychePubKey     string
	TychePrivKey    string
	AdrestiaPubKey  string
	AdrestiaPrivKey string
	HestiaPubKey    string
	HestiaPrivKey   string
	PlutusPubKey    string
	PlutusPrivKey   string
	CoinsVars       []CoinVar
}

func (ev *EnvironmentVars) CheckVars() error {
	if ev.HerokuPassword == "" {
		return errors.New("missing heroku password")
	}
	if ev.HerokuUsername == "" {
		return errors.New("missing heroku username")
	}
	if ev.GinMode == "" {
		return errors.New("missing gin mode")
	}
	if ev.KeyPassword == "" {
		return errors.New("missing key password")
	}
	if ev.AuthUsername == "" {
		return errors.New("missing auth username")
	}
	if ev.AuthPassword == "" {
		return errors.New("missing auth password")
	}
	for _, coinVar := range ev.CoinsVars {
		err := coinVar.CheckVars()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ev *EnvironmentVars) ToString() string {
	str := "" +
		"AUTH_USERNAME=" + ev.AuthUsername + "\n" +
		"AUTH_PASSWORD=" + ev.AuthPassword + "\n" +
		"KEY_PASSWORD=" + ev.KeyPassword + "\n" +
		"GIN_MODE=" + ev.GinMode + "\n" +
		"HEROKU_USERNAME=" + ev.HerokuUsername + "\n" +
		"HEROKU_PASSWORD=" + ev.HerokuPassword + "\n" +
		"TYCHE_PUBLIC_KEY=" + ev.TychePubKey + "\n" +
		"ADRESTIA_PUBLIC_KEY=" + ev.AdrestiaPubKey + "\n" +
		"HESTIA_PUBLIC_KEY=" + ev.HestiaPubKey + "\n" +
		"PLUTUS_PRIVATE_KEY=" + ev.AdrestiaPrivKey + "\n" +
		"PLUTUS_PUBLIC_KEY=" + ev.PlutusPubKey + "\n"

	for _, coinVar := range ev.CoinsVars {
		str += coinVar.ToString()
	}
	return str
}

type CoinVar struct {
	Coin       string
	RpcUser    string
	RpcPass    string
	RpcPort    string
	SshUser    string
	SshHost    string
	SshPrivKey string
	SshPubKey  string
	SshPort    string
	ColdAddrs  string
}

func (cv *CoinVar) CheckVars() error {
	if cv.Coin == "" {
		return errors.New("missing coin tag")
	}
	if cv.RpcUser == "" {
		return errors.New("missing rpc user for " + cv.Coin)
	}
	if cv.RpcPass == "" {
		return errors.New("missing rpc pass for " + cv.Coin)
	}
	if cv.RpcPort == "" {
		return errors.New("missing rpc port for " + cv.Coin)
	}
	if cv.SshUser == "" {
		return errors.New("missing ssh user for " + cv.Coin)
	}
	if cv.SshHost == "" {
		return errors.New("missing ssh ip for " + cv.Coin)
	}
	if cv.SshPrivKey == "" {
		return errors.New("missing ssh private key for " + cv.Coin)
	}
	if cv.SshPort == "" {
		return errors.New("missing ssh port for " + cv.Coin)
	}
	if cv.ColdAddrs == "" {
		return errors.New("missing cold address for " + cv.Coin)
	}
	return nil
}

func (cv *CoinVar) ToString() string {
	str := "" +
		strings.ToUpper(cv.Coin) + "_RPC_USER=" + cv.RpcUser + "\n" +
		strings.ToUpper(cv.Coin) + "_RPC_PASS=" + cv.RpcPass + "\n" +
		strings.ToUpper(cv.Coin) + "_RPC_PORT=" + cv.RpcPort + "\n" +
		strings.ToUpper(cv.Coin) + "_SSH_USER=" + cv.SshUser + "\n" +
		strings.ToUpper(cv.Coin) + "_IP=" + cv.SshHost + "\n" +
		strings.ToUpper(cv.Coin) + "_SSH_PRIVKEY=" + cv.SshPrivKey + "\n" +
		strings.ToUpper(cv.Coin) + "_SSH_PORT=" + cv.SshPort + "\n" +
		strings.ToUpper(cv.Coin) + "_COLD_ADDRESS=" + cv.ColdAddrs + "\n"
	return str
}

type KeyPair struct {
	Private []byte
	Public  []byte
}

var (
	NewVars    EnvironmentVars
	OldVars    EnvironmentVars
	HerokuUser string
	HerokuPass string
)

// This script will only work with a full set of environment variables.
// Should only be used to recreate ssh keys and passwords
func main() {
	// First load the current .env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	HerokuUser = os.Getenv("HEROKU_USERNAME")
	HerokuPass = os.Getenv("HEROKU_PASSWORD")
	if HerokuUser == "" || HerokuPass == "" {
		panic(errors.New("no heroku login details, we can't continue"))
	}
	heroku.DefaultTransport.Username = HerokuUser
	heroku.DefaultTransport.Password = HerokuPass
	h := heroku.NewService(heroku.DefaultClient)
	_, err = h.AppInfo(context.Background(), "plutus-wallets")
	if err != nil {
		panic(err)
	}
	log.Println("Creating environment file...")
	// Move current .env file to a backup file
	date := time.Now().Format("2006-01-02")
	err = os.Rename("../.env", "../old-env-backup-"+date)
	if err != nil {
		log.Fatal("Error moving .env file")
	}
	// Get and object with old variables
	OldVars, err = getOldVars()
	if err != nil {
		panic(err)
	}
	// Generate new keys
	NewVars, err = genNewVars()
	if err != nil {
		panic(err)
	}
	log.Println("Updating remote keys...")
	// Update new ssh keys with a ssh client using old keys
	for _, server := range OldVars.CoinsVars {
		var newCoinPubKey string
		for _, newCoinVar := range NewVars.CoinsVars {
			if newCoinVar.Coin == server.Coin {
				newCoinPubKey = newCoinVar.SshPubKey
			}
		}
		log.Println("Updating remote for " + server.Coin)
		err := updateRemoteKey(server, newCoinPubKey)
		if err != nil {
			panic(err)
		}
	}
	// Update heroku environment variables.
	log.Println("Updating heroku deployment variables...")
	// Create environment map
	envMap := map[string]*string{
		"AUTH_PASSWORD":       &NewVars.AuthPassword,
		"AUTH_USERNAME":       &NewVars.AuthUsername,
		"KEY_PASSWORD":        &NewVars.KeyPassword,
		"TYCHE_PUBLIC_KEY":    &NewVars.TychePubKey,
		"ADRESTIA_PUBLIC_KEY": &NewVars.AdrestiaPubKey,
		"HESTIA_PUBLIC_KEY":   &NewVars.HestiaPubKey,
		"PLUTUS_PUBLIC_KEY":   &NewVars.PlutusPubKey,
		"PLUTUS_PRIVATE_KEY":  &NewVars.PlutusPrivKey,
		"GIN_MODE":            &NewVars.GinMode,
	}
	// First update main variables
	log.Println("Updating main heroku deployment variables...")
	_, err = h.ConfigVarUpdate(context.Background(), "plutus-wallets", envMap)
	if err != nil {
		panic("critical error, unable to update heroku variables")
	}
	for _, env := range NewVars.CoinsVars {
		log.Println("Updating heroku deployment variables for " + strings.ToUpper(env.Coin))
		coinVars := make(map[string]*string)
		coinVars[strings.ToUpper(env.Coin)+"_IP"] = &env.SshHost
		coinVars[strings.ToUpper(env.Coin)+"_RPC_USER"] = &env.RpcUser
		coinVars[strings.ToUpper(env.Coin)+"_RPC_PASS"] = &env.RpcPass
		coinVars[strings.ToUpper(env.Coin)+"_RPC_PORT"] = &env.RpcPort
		coinVars[strings.ToUpper(env.Coin)+"_SSH_USER"] = &env.SshUser
		coinVars[strings.ToUpper(env.Coin)+"_SSH_PORT"] = &env.SshPort
		coinVars[strings.ToUpper(env.Coin)+"_SSH_PRIVKEY"] = &env.SshPrivKey
		coinVars[strings.ToUpper(env.Coin)+"_COLD_ADDRESS"] = &env.ColdAddrs
		_, err := h.ConfigVarUpdate(context.Background(), "plutus-wallets", coinVars)
		if err != nil {
			panic("critical error, unable to update heroku variables")
		}
	}
	log.Println("Updating Plutus access to other microservices...")

	// TYCHE
	log.Println("Updating Plutus access to Tyche")
	tycheAccess := map[string]*string{
		"PLUTUS_AUTH_USERNAME": &NewVars.AuthUsername,
		"PLUTUS_AUTH_PASSWORD": &NewVars.AuthPassword,
		"TYCHE_PRIV_KEY":       &NewVars.TychePrivKey,
		"TYCHE_PUBLIC_KEY":     &NewVars.TychePubKey,
		"ADRESTIA_PUBLIC_KEY":  &NewVars.AdrestiaPubKey,
		"HESTIA_PUBLIC_KEY":    &NewVars.HestiaPubKey,
		"PLUTUS_PUBLIC_KEY":    &NewVars.PlutusPubKey,
	}
	_, err = h.ConfigVarUpdate(context.Background(), "tyche-shift", tycheAccess)
	if err != nil {
		panic("critical error, unable to update heroku variables")
	}

	// ADRESTIA
	log.Println("Updating Plutus access to Adrestia")
	addrestiaAccess := map[string]*string{
		"PLUTUS_AUTH_USERNAME": &NewVars.AuthUsername,
		"PLUTUS_AUTH_PASSWORD": &NewVars.AuthPassword,
		"ADRESTIA_PRIV_KEY":    &NewVars.AdrestiaPrivKey,
		"ADRESTIA_PUBLIC_KEY":  &NewVars.AdrestiaPubKey,
		"TYCHE_PUBLIC_KEY":     &NewVars.TychePubKey,
		"HESTIA_PUBLIC_KEY":    &NewVars.HestiaPubKey,
		"PLUTUS_PUBLIC_KEY":    &NewVars.PlutusPubKey,
	}
	_, err = h.ConfigVarUpdate(context.Background(), "adrestia-exchanges", addrestiaAccess)
	if err != nil {
		panic("critical error, unable to update heroku variables")
	}

	// HESTIA
	log.Println("Updating Hestia keys")
	hestiaAccess := map[string]*string{
		"ADRESTIA_PUBLIC_KEY": &NewVars.AdrestiaPubKey,
		"TYCHE_PUBLIC_KEY":    &NewVars.TychePubKey,
		"HESTIA_PUBLIC_KEY":   &NewVars.HestiaPubKey,
		"HESTIA_PRIVATE_KEY":  &NewVars.HestiaPrivKey,
		"PLUTUS_PUBLIC_KEY":   &NewVars.PlutusPubKey,
	}
	_, err = h.ConfigVarUpdate(context.Background(), "hestia-database", hestiaAccess)
	if err != nil {
		panic("critical error, unable to update heroku variables")
	}

	// Dump new keys to .env file
	err = saveNewVars()
	if err != nil {
		panic(err)
	}
}

func updateRemoteKey(coinVars CoinVar, newCoinPubKey string) error {
	privKey, err := parsePrivKey(coinVars.SshPrivKey, OldVars.KeyPassword)
	if err != nil {
		return err
	}
	sshConfig := &ssh.ClientConfig{
		User: coinVars.SshUser,
		Auth: []ssh.AuthMethod{privKey},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	connection, err := ssh.Dial("tcp", coinVars.SshHost+":"+coinVars.SshPort, sshConfig)
	if err != nil {
		fmt.Println(err)
		return err
	}
	session, err := connection.NewSession()
	if err != nil {
		fmt.Println(err)
		return err
	}
	// First cmd remove the second line (the first line is the main key and must be changed manually to prevent removing all access to server)
	// Second cmd append the newCoinPubKey to the authorized_keys file
	cmd := "sed -i '2d' .ssh/authorized_keys && sed -i '$a" + newCoinPubKey + "' .ssh/authorized_keys"
	err = session.Run(cmd)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getOldVars() (EnvironmentVars, error) {
	Vars := EnvironmentVars{
		HerokuUsername: os.Getenv("HEROKU_USERNAME"),
		HerokuPassword: os.Getenv("HEROKU_PASSWORD"),
		AuthUsername:   os.Getenv("AUTH_USERNAME"),
		AuthPassword:   os.Getenv("AUTH_PASSWORD"),
		GinMode:        os.Getenv("GIN_MODE"),
		KeyPassword:    os.Getenv("KEY_PASSWORD"),
		CoinsVars:      nil,
	}
	for key := range coinfactory.Coins {
		coinVars := CoinVar{
			Coin:       strings.ToUpper(key),
			RpcUser:    os.Getenv(strings.ToUpper(key) + "_RPC_USER"),
			RpcPass:    os.Getenv(strings.ToUpper(key) + "_RPC_PASS"),
			RpcPort:    os.Getenv(strings.ToUpper(key) + "_RPC_PORT"),
			SshUser:    os.Getenv(strings.ToUpper(key) + "_SSH_USER"),
			SshPrivKey: os.Getenv(strings.ToUpper(key) + "_SSH_PRIVKEY"),
			SshPubKey:  "",
			SshPort:    os.Getenv(strings.ToUpper(key) + "_SSH_PORT"),
			SshHost:    os.Getenv(strings.ToUpper(key) + "_IP"),
			ColdAddrs:  os.Getenv(strings.ToUpper(key) + "_COLD_ADDRESS"),
		}
		Vars.CoinsVars = append(Vars.CoinsVars, coinVars)
	}
	err := Vars.CheckVars()
	return Vars, err
}

func genNewVars() (EnvironmentVars, error) {
	newAuthUsername := generateRandomPassword(128)
	newAuthPassword := generateRandomPassword(128)
	newDecryptionKey := generateRandomPassword(32)
	// Tyche KeyPair
	newTychePair, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	tychePubKeyBytes, _ := asn1.Marshal(newTychePair.PublicKey)
	tychePrivKeyBytes := x509.MarshalPKCS1PrivateKey(newTychePair)

	// Adrestia KeyPair
	newAdrestiaKeyPair, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	addrestiaPubKeyBytes, _ := asn1.Marshal(newAdrestiaKeyPair.PublicKey)
	addrestiaPrivKeyBytes := x509.MarshalPKCS1PrivateKey(newAdrestiaKeyPair)

	// Hestia KeyPair
	newHestiaKeyPair, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	hestiaPubKeyBytes, _ := asn1.Marshal(newHestiaKeyPair.PublicKey)
	hestiaPrivKeyBytes := x509.MarshalPKCS1PrivateKey(newHestiaKeyPair)

	// Plutus KeyPair
	newPlutusKeyPair, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	plutusPubKeyBytes, _ := asn1.Marshal(newPlutusKeyPair.PublicKey)
	plutusPrivKeyBytes := x509.MarshalPKCS1PrivateKey(newPlutusKeyPair)
	Vars := EnvironmentVars{
		HerokuUsername:  os.Getenv("HEROKU_USERNAME"),
		HerokuPassword:  os.Getenv("HEROKU_PASSWORD"),
		AuthUsername:    newAuthUsername,
		AuthPassword:    newAuthPassword,
		GinMode:         os.Getenv("GIN_MODE"),
		KeyPassword:     newDecryptionKey,
		TychePrivKey:    base64.StdEncoding.EncodeToString(tychePrivKeyBytes),
		TychePubKey:     base64.StdEncoding.EncodeToString(tychePubKeyBytes),
		AdrestiaPrivKey: base64.StdEncoding.EncodeToString(addrestiaPrivKeyBytes),
		AdrestiaPubKey:  base64.StdEncoding.EncodeToString(addrestiaPubKeyBytes),
		PlutusPrivKey:   base64.StdEncoding.EncodeToString(plutusPrivKeyBytes),
		PlutusPubKey:    base64.StdEncoding.EncodeToString(plutusPubKeyBytes),
		HestiaPubKey:    base64.StdEncoding.EncodeToString(hestiaPubKeyBytes),
		HestiaPrivKey:   base64.StdEncoding.EncodeToString(hestiaPrivKeyBytes),
		CoinsVars:       nil,
	}
	for key := range coinfactory.Coins {
		log.Println("Creating vars for " + strings.ToUpper(key))
		keyPair := genPrivKeyPair()
		encryptedPrivKey, err := aes.Encrypt([]byte(newDecryptionKey), keyPair.Private)
		if err != nil {
			panic(err)
		}
		coinVars := CoinVar{
			Coin:       strings.ToUpper(key),
			RpcUser:    os.Getenv(strings.ToUpper(key) + "_RPC_USER"),
			RpcPass:    os.Getenv(strings.ToUpper(key) + "_RPC_PASS"),
			RpcPort:    os.Getenv(strings.ToUpper(key) + "_RPC_PORT"),
			SshUser:    os.Getenv(strings.ToUpper(key) + "_SSH_USER"),
			SshPrivKey: encryptedPrivKey,
			SshPubKey:  string(keyPair.Public),
			SshPort:    os.Getenv(strings.ToUpper(key) + "_SSH_PORT"),
			SshHost:    os.Getenv(strings.ToUpper(key) + "_IP"),
			ColdAddrs:  os.Getenv(strings.ToUpper(key) + "_COLD_ADDRESS"),
		}
		Vars.CoinsVars = append(Vars.CoinsVars, coinVars)
	}
	err = Vars.CheckVars()
	return Vars, err
}

func saveNewVars() error {
	err := ioutil.WriteFile("../../.env", []byte(NewVars.ToString()), 777)
	if err != nil {
		return err
	}
	return nil
}

func genPrivKeyPair() KeyPair {
	privateKey, err := generatePrivateKey()
	if err != nil {
		log.Fatal(err.Error())
	}
	privDER, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	privBlock := pem.Block{
		Type:    "EC PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}
	privateBytes := pem.EncodeToMemory(&privBlock)
	publicKeyBytes, err := generatePublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal(err.Error())
	}
	return KeyPair{Private: privateBytes, Public: publicKeyBytes}
}

func generatePrivateKey() (*ecdsa.PrivateKey, error) {
	pubkeyCurve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func generatePublicKey(privatekey *ecdsa.PublicKey) ([]byte, error) {
	publicKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}
	pubKeyBytes := ssh.MarshalAuthorizedKey(publicKey)
	return pubKeyBytes, nil
}

func generateRandomPassword(size int) string {
	res, err := password.Generate(size, 10, 0, false, true)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func parsePrivKey(privKey string, encryptionKey string) (ssh.AuthMethod, error) {
	decrypted, err := aes.Decrypt([]byte(encryptionKey), privKey)
	if err != nil {
		return nil, err
	}
	key, err := ssh.ParsePrivateKey([]byte(decrypted))
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}
