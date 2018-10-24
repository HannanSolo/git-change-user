all: dependencies build

dependencies:
	go get

build:
	go build -o bin/gcu gcu.go
