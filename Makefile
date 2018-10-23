all: install build

install:
	go get

build:
	go build -o bin/gcu gcu.go
