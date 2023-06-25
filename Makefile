
build:	unit-testing build-backgroundworker build-mux build-infra
run: start-services  start-backgroundworker

lint:
	go vet ./...
	golint ./...

build-backgroundworker:
	mkdir -p bin && cd cmd/rabbit && GOOS=windows GOARCH=amd64 go build -o ../../bin

build-mux:
	cd cmd/mux && GOOS=windows GOARCH=amd64 go build -o ../../bin

build-infra:
	cd infrastructure && docker-compose build

start-services:
	cd infrastructure && docker-compose up -d

start-backgroundworker:
	cd bin && ./rabbit.exe

start-api:
	cd bin && ./mux.exe

unit-testing:
	go test -coverprofile=c.out ./...
	go tool cover -html=c.out

