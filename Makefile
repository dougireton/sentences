.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./bin

build:
	GOOS=linux GOARCH=amd64 go build -o bin/api ./api

local-api:
	sam local start-api