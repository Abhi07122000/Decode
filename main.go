package main

import (
	"encoding/binary"
	"fmt"
)

const (
	packetSize  = 44
	shortSize   = 2
	string1Size = 12
	byteSize    = 1
	string2Size = 8
	short2Size  = 2
	string3Size = 15
	longSize    = 4
)

type DecodedStruct struct {
	Short1  uint16
	String1 string
	Byte1   byte
	String2 string
	Short2  uint16
	String3 string
	Long1   uint32
}

func decodePacket(packet []byte) (DecodedStruct, error) {
	if len(packet) != packetSize {
		return DecodedStruct{}, fmt.Errorf("Invalid packet size")
	}

	var decoded DecodedStruct
	reader := binary.BigEndian

	decodeString := func(offset, length int) string {
		return string(packet[offset : offset+length])
	}

	offset := 0

	decoded.Short1 = reader.Uint16(packet[offset : offset+shortSize])
	offset += shortSize

	decoded.String1 = decodeString(offset, string1Size)
	offset += string1Size

	decoded.Byte1 = packet[offset]
	offset += byteSize

	decoded.String2 = decodeString(offset, string2Size)
	offset += string2Size

	decoded.Short2 = reader.Uint16(packet[offset : offset+short2Size])
	offset += short2Size

	decoded.String3 = decodeString(offset, string3Size)
	offset += string3Size

	decoded.Long1 = reader.Uint32(packet[offset : offset+longSize])

	return decoded, nil
}

func main() {
	packet := []byte{0x04, 0xD2, 0x6B, 0x65, 0x65, 0x70, 0x64, 0x65, 0x63, 0x6F, 0x64, 0x69, 0x6E, 0x67, 0x38, 0x64, 0x6F, 0x6E, 0x74, 0x73, 0x74, 0x6F, 0x70, 0x03, 0x15, 0x63, 0x6F, 0x6E, 0x67, 0x72, 0x61, 0x74, 0x75, 0x6C, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x73, 0x07, 0x5B, 0xCD, 0x15}
	decoded, err := decodePacket(packet)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Decoded struct: {%d, \"%s\", %d, \"%s\", %d, \"%s\", %d}\n", decoded.Short1, decoded.String1, decoded.Byte1, decoded.String2, decoded.Short2, decoded.String3, decoded.Long1)
}
