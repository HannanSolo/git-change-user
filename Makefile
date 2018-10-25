all: dependencies build

dependencies:
	go get

build:
	go build -o bin/gcu gcu.go

install:
	go build -o bin/gcu gcu.go
	mv bin/gcu /bin
