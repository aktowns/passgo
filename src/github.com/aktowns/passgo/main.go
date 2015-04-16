package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	passwordsPath := path.Join(homeDir(), ".passwords.yml")
	privateKeyPath := path.Join(homeDir(), ".passgo")

	store := createYAMLStore(passwordsPath)
	args := os.Args[1:]

	if !fileExist(privateKeyPath) {
		writeRSAPrivateKey(generateRSAPrivateKey(2048), privateKeyPath)
	}

	privKey := readRSAPrivateKey(privateKeyPath)

	if len(args) == 0 {
		fmt.Println("Usage: " + os.Args[0] + " key")
		fmt.Println("       " + os.Args[0] + " key value")
	} else if len(args) == 1 {
		contents := store.ReadBinary(args[0])
		fmt.Printf(string(decryptRSAString(privKey, contents)))
	} else if len(args) == 2 && args[1] == "-" {
		contents, err := ioutil.ReadAll(os.Stdin)
		check(err)

		store.WriteBinary(args[0], encryptRSAString(&privKey.PublicKey, contents))
	} else {
		val := strings.Join(args[1:], " ")

		store.WriteBinary(args[0], encryptRSAString(&privKey.PublicKey, []byte(val)))
	}
}
