export GOPATH      ?= $(shell pwd)

GO         ?= go
GOFMT      ?= gofmt

editor:
	sh -c $(EDITOR) .

build:
	go install github.com/aktowns/passgo
	bin/passgo
