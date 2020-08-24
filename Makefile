TAG = dev
IMG ?= quay.io/xenitab/ingress-healthz:$(TAG)

lint:
	golangci-lint run -E misspell

fmt:
	go fmt ./...

vet:
	go vet ./...

test: fmt vet
	go test -timeout 1m ./...

build:
	go build -a -o bin/ingress-healthz cmd/ingress-healthz/main.go

docker-build:
	docker build -t $(IMG) .
