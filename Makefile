PACKAGES=$(shell go list ./...)
VERSION ?= v1.0.0
IMAGE_VERSION ?= $(VERSION)

.PHONY: all binary test vet lint clean

SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')

default: binary

all: binary

fmt:
	gofmt -w $(SRCS)

vet:
	@go vet ${PACKAGES}

lint:
	@ go get -v github.com/golang/lint/golint
	$(foreach file,$(SRCS),golint $(file) || exit;)

binary: lint fmt vet
	mkdir -p bin/
	CGO_ENABLED=0 go build -o bin/authz-broker --ldflags "-X \"main.version=$(VERSION)\"" -a -installsuffix cgo ./broker/main.go

test: binary
	go test -v ${PACKAGES}

clean:
	rm -rf bin/
