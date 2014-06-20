# omron-fins-go

[![Build Status](https://travis-ci.org/hiroeorz/omron-fins-go.svg?branch=master)](https://travis-ci.org/hiroeorz/omron-fins-go)

## About

This is fins command client written by Go.

This library support communication to omron PLC from Go application.

* omron: <http://www.omron.co.jp/>
* omron PLC: <http://www.fa.omron.co.jp/products/category/automation-systems/programmable-controllers/>

##Install

```
$ go get github.com/hiroeorz/omron-fins-go/fins
```

## Usage

Read DM values sample.

```go
package main

import (
	"fmt"
	"github.com/hiroeorz/fins"
	"log"
)

func main() {
	srcAddr := "192.168.0.1:9600" // Host address
	dstAddr := "192.168.0.6:9600"  // PLC address

	// Start listener.
	listenChan := fins.Listen(srcAddr)

	// Send ReadDM request (startAddress:100 getCount:10).
	vals, err := fins.ReadDM(listenChan, srcAddr, dstAddr, 100, 10)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(vals)
	// => [12799 24752 12799 24768 12799 24784 12799 24800 12799 24816]
}
```
