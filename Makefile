.PHONY: test

GO=go
GITHUB=github.com/hiroeorz/omron-fins-go
PACKAGE=github.com/hiroeorz/fins
GOOS=linux
GOARCH=arm
GOARM=5

all: get build

get:
	${GO} get ${GITHUB}/fins

build:
	${GO} install ${PACKAGE}

arm:
	GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} ${GO} install ${PACKAGE}

clean:
	rm -f ${GOPATH}/pkg/*/${PACKAGE}.a

fmt:
	go fmt ./...
