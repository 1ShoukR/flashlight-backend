# Simple Makefile for a Go project

all: build

build:
	@echo "Building..."
	
	@go build -o main.exe cmd/main.go

run:
	@go run cmd/main.go
