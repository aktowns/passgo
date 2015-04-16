package main

// Store represents a generic storage interface..
type Store interface {
	Write(key string, value string)
	WriteBinary(key string, value []byte)
	Read(key string) string
	ReadBinary(key string) []byte
}
