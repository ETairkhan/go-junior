.PHONY: build
build:
	go build -o j ./cmd/j/main.go

.DEFAULT_GOAL := build