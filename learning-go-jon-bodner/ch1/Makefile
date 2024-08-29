.DEFAULT_GOAL := run 

fmt:
	go fmt ./...

lint: fmt
	golint ./...

vet: fmt
	go vet ./...

build: vet
	go build .

run: vet
	go run .

.PHONY:fmt lint
