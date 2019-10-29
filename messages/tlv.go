package messages

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type Tlv struct {
	MessageType uint8
	Version     uint8
	Length      uint16
}

func (tlv *Tlv) ToBytes() ([]byte, error) {

	buf := &bytes.Buffer{}

	buf.Write([]byte{tlv.MessageType})

	buf.Write([]byte{tlv.Version})

	err := binary.Write(buf, binary.BigEndian, tlv.Length)
	if err != nil {
		return []byte{}, fmt.Errorf(
			"could not decode lenth property of tlv (%v) to binary: %s",
			tlv.Length,
			err.Error(),
		)
	}

	return ioutil.ReadAll(buf)

}

func newTlvFromBytes(bytes []byte) (Tlv, error) {

	if len(bytes) != 4 {
		return Tlv{}, fmt.Errorf("tlv byte array does not meet Length of 4: %v", bytes)
	}

	return Tlv{
		MessageType: bytes[0],
		Version:     bytes[1],
		Length:      binary.BigEndian.Uint16(bytes[2:]),
	}, nil

}
