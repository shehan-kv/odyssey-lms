.DEFAULT_GOAL := build

.PHONY:sqlc fmt vet build

sqlc:
	sqlc generate

fmt: sqlc
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build ./cmd/...
