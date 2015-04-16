package main

import (
	"log"
	"os"
	"os/user"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func homeDir() string {
	usr, err := user.Current()
	check(err)

	return usr.HomeDir
}
