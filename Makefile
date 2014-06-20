.PHONY: test

GO=go
PACKAGE=github.com/hiroeorz/omron-fins-go/fins
GOOS=linux
GOARCH=arm
GOARM=5

all: build

build:
	${GO} install ${PACKAGE}

arm:
	GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} ${GO} install ${PACKAGE}

clean:
	rm -f ${GOPATH}/pkg/*/${PACKAGE}.a

fmt:
	go fmt ./...
