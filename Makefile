.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt: sqlc
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build ./cmd/...
