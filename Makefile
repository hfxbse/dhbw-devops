GO_CMD = /usr/bin/env go
all: build

build: build-auth build-checkout build-products

build-auth:
	cd auth && CGO_ENABLED=0 $(GO_CMD) build -o ../out/auth ./cmd/main.go

build-checkout:
	cd checkout && CGO_ENABLED=0 $(GO_CMD) build -o ../out/checkout ./cmd/main.go

build-products:
	cd products && CGO_ENABLED=0 $(GO_CMD) build -o ../out/products ./cmd/main.go

clean:
	rm -f out/*
