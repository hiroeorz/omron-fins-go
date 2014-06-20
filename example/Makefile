.PHONY: test

GO=go
GOOS=linux
GOARCH=arm
GOARM=5

all: build

build:
	${GO} build

arm:
	GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} ${GO} build

clean:
	rm -f example

fmt:
	go fmt ./...
