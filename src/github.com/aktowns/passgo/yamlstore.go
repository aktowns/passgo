package main

import (
	"encoding/base64"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// YAMLStore uses yaml to fulfil the storage interface in Store.go
type YAMLStore struct {
	filename  string
	structure map[string]string
}

func createYAMLStore(filename string) Store {
	store := make(map[string]string)
	if fileExist(filename) {
		bytes, err := ioutil.ReadFile(filename)
		check(err)

		err = yaml.Unmarshal(bytes, &store)
		check(err)
	} else {
		d, err := yaml.Marshal(&store)
		check(err)

		err = ioutil.WriteFile(filename, d, 0600)
	}
	return YAMLStore{filename: filename, structure: store}
}

func (store YAMLStore) Write(key string, value string) {
	store.structure[key] = value

	d, err := yaml.Marshal(&store.structure)
	check(err)

	err = ioutil.WriteFile(store.filename, d, 0600)
	check(err)
}

func (store YAMLStore) WriteBinary(key string, value []byte) {
	store.Write(key, base64.StdEncoding.EncodeToString(value))
}

func (store YAMLStore) Read(key string) string {
	return store.structure[key]
}

func (store YAMLStore) ReadBinary(key string) []byte {
	val, err := base64.StdEncoding.DecodeString(store.Read(key))
	check(err)

	return val
}
