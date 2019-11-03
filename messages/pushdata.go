package messages

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type PushDataMessage struct {
	Header Tlv
	Body   PushDataBody
}

type PushDataBody struct {
	TransactionId  uint16
	SequenceNumber uint16
	MacAddress     MacAddress
	Type           uint8
	Version        uint8
	Length         uint16
	Data           string
}

func (message *PushDataMessage) MessageId() uint8 {
	return 0
}

func (message *PushDataMessage) ToBytes() ([]byte, error) {

	bodyBytes, err := message.Body.toBytes()
	if err != nil {
		return []byte{}, fmt.Errorf("unable to convert push data body to bytes: %s", err.Error())
	}

	message.Header.Length = uint16(len(bodyBytes))
	headerBytes, err := message.Header.ToBytes()

	return append(headerBytes, bodyBytes...), nil

}

func (message *PushDataBody) toBytes() ([]byte, error) {

	buf := &bytes.Buffer{}

	err := binary.Write(buf, binary.BigEndian, message.TransactionId)
	if err != nil {
		return []byte{}, fmt.Errorf(
			"could not decode transaction id %v to binary: %s",
			message.TransactionId,
			err.Error(),
		)
	}

	err = binary.Write(buf, binary.BigEndian, message.SequenceNumber)
	if err != nil {
		return []byte{}, fmt.Errorf(
			"could not decode sequence number %v to binary: %s",
			message.TransactionId,
			err.Error(),
		)
	}

	macAddressBytes, err := message.MacAddress.ToBytes()
	if err != nil {
		return []byte{}, fmt.Errorf(
			"could not decode mac address %s to binary: %s",
			message.MacAddress.Value,
			err.Error(),
		)
	}
	buf.Write(macAddressBytes)
	buf.Write([]byte{message.Type})
	buf.Write([]byte{message.Version})
	binaryData := []byte(message.Data)

	err = binary.Write(buf, binary.BigEndian, uint16(len(binaryData)))
	if err != nil {
		return []byte{}, fmt.Errorf("could not encode data Length to binary: %s", err.Error())
	}
	buf.Write(binaryData)

	return ioutil.ReadAll(buf)

}

func newPushDataBodyFromBytes(bytes []byte) (PushDataBody, error) {

	macAddress, err := newMacAddressFromBytes(bytes[4:10])
	if err != nil {
		return PushDataBody{}, err
	}

	b := PushDataBody{
		TransactionId:  binary.BigEndian.Uint16(bytes[0:2]),
		SequenceNumber: binary.BigEndian.Uint16(bytes[2:4]),
		MacAddress:     macAddress,
		Type:           bytes[10],
		Version:        bytes[11],
		Length:         binary.BigEndian.Uint16(bytes[12:14]),
		Data:           string(bytes[14:]),
	}

	return b, nil

}

func NewPushDataMessageFromBytes(bytes []byte) (PushDataMessage, error) {

	tlv, err := newTlvFromBytes(bytes[:4])
	if err != nil {
		return PushDataMessage{}, err
	}

	b, err := newPushDataBodyFromBytes(bytes[4:])
	if err != nil {
		return PushDataMessage{}, err
	}

	return PushDataMessage{
		Header: tlv,
		Body:   b,
	}, nil

}
