package rsa_crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"openssl_rsa256/config"
)

type RsaCrypt struct {
	PrivateKeyPath string
	PublicKeyPath  string
}

func GetRsaCrypt() RsaCrypt {
	var rsaCrypt RsaCrypt
	rsaCrypt.PrivateKeyPath = config.GetRsaConf().PrivatePath
	rsaCrypt.PublicKeyPath = config.GetRsaConf().PublicPath
	return rsaCrypt
}
func (r *RsaCrypt) Encrypt(data string) ([]byte, error) {
	publicKey, err := GetContent(r.PublicKeyPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode([]byte(publicKey))

	if block == nil {
		return nil, errors.New("public key error")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, key.(*rsa.PublicKey), []byte(data))
}
func (r *RsaCrypt) Decrypt(data []byte) ([]byte, error) {
	privateKey, err := GetContent(r.PrivateKeyPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("private key error")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, key.(*rsa.PrivateKey), []byte(data))
}

func GetContent(pathStr string) (string, error) {
	bytes, err := ioutil.ReadFile(pathStr)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
