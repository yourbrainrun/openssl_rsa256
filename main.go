package main

import (
	"encoding/base64"
	"fmt"
	"openssl_rsa256/rsa_crypt"
)

func main() {
	data := "Spring is here again, but not all of it exists"
	cryptAlgo := rsa_crypt.GetRsaCrypt()

	encrypt, err := Encrypt(data, &cryptAlgo)
	if err != nil {
		return
	}
	fmt.Println(base64.StdEncoding.EncodeToString(encrypt), "encrypt")

	decrypt, err := Decrypt(encrypt, &cryptAlgo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(decrypt), "decrypt")
}

func Encrypt(data string, algo rsa_crypt.ICrypt) ([]byte, error) {
	return algo.Encrypt(data)
}

func Decrypt(data []byte, algo rsa_crypt.ICrypt) ([]byte, error) {
	return algo.Decrypt(data)
}
