package main

import (
	"fmt"
	"github.com/hiroeorz/omron-fins-go/fins"
	"log"
)

func main() {
	srcAddr := "192.168.10.30:9600"
	dstAddr := "192.168.10.6:9600"
	listenChan := fins.Listen(srcAddr)

	vals, err := fins.ReadDM(listenChan, srcAddr, dstAddr, 100, 10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(vals)
}
