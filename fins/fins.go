package fins

import (
	"errors"
	"fmt"
	"log"
	"net"
)

type ReqCommand struct {
	ch      chan []byte
	command []byte
	dstAddr string
}

// ReadIO get uint16 values from PLC DM Areas.
func ReadDM(listenChan chan *ReqCommand,
	srcAddr string, dstAddr string, startAddr uint16, readCount uint8) ([]uint16, error) {

	header := newHeader(srcAddr, dstAddr)
	command := readDMCom(header, 0x82, startAddr, readCount)

	bytes, err := syncSend(listenChan, srcAddr, dstAddr, command)
	if err != nil {
		log.Fatal(err)
		return []uint16{}, errors.New("fins: send command error")
	}

	return parseReadDM(bytes)
}

// Listen start listenLoop and return channel that used in listenLoop.
func Listen(srcAddr string) chan *ReqCommand {
	udpAddress, err := net.ResolveUDPAddr("udp4", srcAddr)
	if err != nil {
		log.Fatal(err)
		panic(fmt.Sprintf("error resolving UDP port: %s\n", srcAddr))
	}

	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		log.Fatal(err)
		panic(fmt.Sprintf("error listening UDP port: %s\n", srcAddr))
	}

	listenChan := make(chan *ReqCommand)
	go listenLoop(listenChan, conn)
	return listenChan
}

// listenLoop wait request through a channel and send request and receive response.
func listenLoop(listenChan chan *ReqCommand, conn *net.UDPConn) {
	defer conn.Close()
	for {
		reqCom := <-listenChan

		n, err := send(reqCom.command, reqCom.dstAddr)
		if err != nil || n == 0 {
			log.Println("cannot write reqest: ", err, " bytes:", n)
			continue
		}

		var buf []byte = make([]byte, 1500)
		n, address, err := conn.ReadFromUDP(buf)

		if err != nil {
			log.Fatal(err)
		}

		if address != nil && n > 0 {
			reqCom.ch <- buf[0:n]
		} else {
			log.Println("cannot read reqest: ", err, " bytes:", n)
			reqCom.ch <- nil
		}
	}
}

// syncSend send binary command to PLC and wait response.
func syncSend(listenChan chan *ReqCommand, srcAddr string, dstAddr string, command []byte) ([]byte, error) {
	reqCom := &ReqCommand{make(chan []byte), command, dstAddr}
	listenChan <- reqCom
	bytes := <-reqCom.ch
	return bytes, nil
}

// send is send command data to PLC.
func send(command []byte, dstAddr string) (int, error) {
	conn, err := net.Dial("udp", dstAddr)
	if err != nil {
		return 0, err
	}

	defer conn.Close()
	return conn.Write(command)
}
