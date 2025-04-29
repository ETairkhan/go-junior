.PHONY: build
build:
	go build -o j ./cmd/server/main.go

.DEFAULT_GOAL := build