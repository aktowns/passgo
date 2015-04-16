package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func decryptRSAString(private *rsa.PrivateKey, input []byte) []byte {
	val, err := rsa.DecryptPKCS1v15(rand.Reader, private, input)
	check(err)

	return val
}

func encryptRSAString(private *rsa.PublicKey, input []byte) []byte {
	val, err := rsa.EncryptPKCS1v15(rand.Reader, private, input)
	check(err)

	return val
}

func readRSAPrivateKey(privateKey string) *rsa.PrivateKey {
	keyBytes, err := ioutil.ReadFile(privateKey)
	check(err)

	block, rest := pem.Decode(keyBytes)
	if len(rest) > 0 {
		fmt.Println("Extra data?")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	check(err)

	return key
}

func generateRSAPrivateKey(bits int) *rsa.PrivateKey {
	priv, err := rsa.GenerateKey(rand.Reader, bits)
	check(err)

	return priv
}

func writeRSAPrivateKey(key *rsa.PrivateKey, privateKeyLocation string) {
	privateKey := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})

	err := ioutil.WriteFile(privateKeyLocation, privateKey, 0600)
	check(err)
}

func writeRSAPublicKey(key *rsa.PrivateKey, publicKeyLocation string) {
	pub, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	check(err)

	publicKey := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: pub})

	err = ioutil.WriteFile(publicKeyLocation, publicKey, 0644)
	check(err)
}
