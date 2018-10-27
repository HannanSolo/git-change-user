all: dependencies build

dependencies:
	go get

build:
	go build -o bin/gcu gcu.go

install: build
	mv bin/gcu /usr/local/bin
