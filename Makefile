.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./bin

test:
	go test -v ./api
	go test -v ./sentences

build:
	GOOS=linux GOARCH=amd64 go build -o bin/api ./api


local-api:
	sam local start-api