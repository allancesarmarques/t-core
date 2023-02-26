export PATH := ${PATH}:$(go env GOPATH)/bin

all: clean swag build run

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ./cmd/main.go

run:
	go run ./cmd/main.go

clean:
	go mod tidy

lint:
	gofmt -w .

swag-install:
	cd /tmp && go install github.com/swaggo/swag/cmd/swag@latest

swag: swag-install
	swag init -g ./cmd/main.go
