package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/grupokindynos/common/aes"
	"io/ioutil"
	"log"
)

func main() {
	file := flag.String("file", "", "private key location string")
	password := flag.String("key", "", "encryption key (max length 32 bytes)")
	flag.Parse()
	if *file == "" {
		log.Fatal(errors.New("please set a file path"))
	}
	if *password == "" {
		log.Fatal(errors.New("please set a password"))
	}
	fileBytes, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}
	encrypted, err := aes.Encrypt([]byte(*password), fileBytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted Key")
	fmt.Println(encrypted)
}
