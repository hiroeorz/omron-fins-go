package fins

import (
	"errors"
	"fmt"
)

func readDMCom(header *FinsHeader, ioFacility uint8, startAddress uint16,
	readCount uint8) []byte {

	if ioFacility != 0x82 && ioFacility != 0x02 {
		panic(fmt.Sprintf("invalid ioFacility: %d\n", ioFacility))
	}

	var addressBit byte = 0

	headerBytes := header.Format()
	addressLower := byte(startAddress)
	addressUpper := byte(startAddress >> 8)
	countLower := byte(readCount)
	countUpper := byte(readCount >> 8)

	code := []byte{1, 1}
	paramsBytes := []byte{
		ioFacility,
		addressUpper, addressLower,
		addressBit,
		countUpper, countLower}

	bytes1 := append(headerBytes, code...)
	bytes2 := append(bytes1, paramsBytes...)
	return bytes2
}

func parseReadDM(bytes []byte) ([]uint16, error) {
	err := validate(bytes)
	if err != nil {
		return []uint16{}, err
	}

	body := bytes[14:]
	var result []uint16

	for i := 0; i < (len(body) / 2); i++ {
		var upper uint16 = (uint16(body[i*2]) << 8)
		var lower uint16 = uint16(body[i*2+1])
		result = append(result, (upper | lower))
	}

	return result, nil
}

func validate(bytes []byte) error {
	finishCode1 := bytes[12]
	finishCode2 := bytes[13]

	if finishCode1 != 0 || finishCode1 != 0 {
		msg := fmt.Sprintln("failure code:", finishCode1, ":", finishCode2)
		return errors.New(msg)
	}

	return nil
}
