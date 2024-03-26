# Makefile

IMAGE_NAME = go-games-api
IMAGE_TAG = latest

.PHONY: run test build

run:
	go run main.go

test:
	go test -v .

build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
