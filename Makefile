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

build_all: vet build

topic: vet build
	@./bin/add_topic

consumer: vet build
	@./bin/consumer

run_multi_consumer:
	 @./bin/consumer >> out_consumer/data_consumer_1.txt
	 @./bin/consumer >> out_consumer/data_consumer_2.txt &
	 @./bin/consumer >> out_consumer/data_consumer_3.txt &
	 @./bin/consumer >> out_consumer/data_consumer_4.txt &
	 @./bin/consumer >> out_consumer/data_consumer_5.txt &