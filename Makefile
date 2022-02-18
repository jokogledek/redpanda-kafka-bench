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
	@go build -o ./cmd/add_topic/app ./cmd/add_topic/
	@go build -o ./cmd/consumer/app ./cmd/consumer/
	@go build -o ./cmd/data_loader/app ./cmd/data_loader/

topic: vet build
	@./cmd/add_topic/app

consumer: vet build
	@./cmd/consumer/app