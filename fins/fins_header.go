package fins

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type FinsHeader struct {
	icf byte
	rsv byte
	gct byte
	dna byte
	da1 byte
	da2 byte
	sna byte
	sa1 byte
	sa2 byte
	sid byte
}

func newHeader(srcAddr string, dstAddr string) *FinsHeader {
	h := new(FinsHeader)
	h.icf = icf()
	h.rsv = rsv()
	h.gct = gct()
	h.SetDstNetwork(0)
	h.SetDstNode(getAddrNode(dstAddr))
	h.SetDstUnit(0)
	h.SetSrcNetwork(0)
	h.SetSrcNode(getAddrNode(srcAddr))
	h.SetSrcUnit(0)
	h.SetIdentifier(0)
	return h
}

func getAddrNode(addr string) byte {
	addrPort := strings.Split(addr, ":")
	if len(addrPort) != 2 {
		panic(fmt.Sprintln("invalid PLC address:", addr))
	}

	parts := strings.Split(addrPort[0], ".")
	if len(parts) != 4 {
		panic(fmt.Sprintln("invalid PLC address:", addr))
	}

	node, err := strconv.ParseUint(parts[3], 10, 8)
	if err != nil {
		log.Fatal(err)
		panic(fmt.Sprintln("invalid PLC address:", addr))
	}

	return byte(node)
}

func (f *FinsHeader) Format() []byte {

	return []byte{
		f.icf, f.rsv, f.gct,
		f.dna, f.da1, f.da2,
		f.sna, f.sa1, f.sa2,
		f.sid}
}

func icf() byte {
	return 1 << 7
}

func rsv() byte {
	return 0x00
}

func gct() byte {
	return 0x02
}

func (f *FinsHeader) SetDstNetwork(network byte) {
	if network == 0 || (1 <= network && network <= 0x7F) {
		f.dna = network
		return
	}
	panic(fmt.Sprintf("invalid network at dna: %d", network))
}

func (f *FinsHeader) SetDstNode(node byte) {
	if node == 0 || node == 0xFF || (0 <= node && node <= 0x20) {
		f.da1 = node
		return
	}
	panic(fmt.Sprintf("invalid network at da1: %d", node))
}

func (f *FinsHeader) SetDstUnit(unitNo byte) {
	if unitNo == 0 || unitNo == 0xFE || unitNo == 0xE1 ||
		(1 <= unitNo && unitNo <= 0x7F) {
		f.da2 = unitNo
		return
	}
	panic(fmt.Sprintf("invalid unitNo at da2: %d", unitNo))
}

func (f *FinsHeader) SetSrcNetwork(srcNetwork byte) {
	if srcNetwork == 0 || (1 <= srcNetwork && srcNetwork <= 0x7F) {
		f.sna = srcNetwork
		return
	}
	panic(fmt.Sprintf("invalid srcNetwork at sna: %d", srcNetwork))
}

func (f *FinsHeader) SetSrcNode(node byte) {
	if node == 0 || (1 <= node && node <= 0x20) {
		f.sa1 = node
		return
	}
	panic(fmt.Sprintf("invalid node at sa1: %d", node))
}

func (f *FinsHeader) SetSrcUnit(unitNo byte) {
	if unitNo == 0 || (10 <= unitNo && unitNo <= 0x1F) {
		f.sa2 = unitNo
		return
	}
	panic(fmt.Sprintf("invalid unitNo at sa2: %d", unitNo))
}

func (f *FinsHeader) SetIdentifier(identifier byte) {
	if 0 <= identifier && identifier <= 0xFF {
		f.sid = identifier
		return
	}
	panic(fmt.Sprintf("invalid identifier at sid: %d", identifier))
}
