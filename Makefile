#!/bin/bash

export PKGS=$(shell go list ./...)
export GOPRIVATE=github.com/ujunglangit-id
export GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct
export GO111MODULE=on

vet:
	@echo "---- VERIFY ----"
	@go vet ./... ${PKGS}

build:
	@echo "---- BUILD -----"
	@go mod tidy
	@go build -o ./bin/add_topic ./cmd/add_topic/
	@go build -o ./bin/consumer ./cmd/consumer/
	@go build -o ./bin/data_loader ./cmd/data_loader/

topic: vet build
	@./bin/add_topic

consumer: vet build
	@./bin/consumer