package messages

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type RequestMessage struct {
	Header Tlv
	Body   RequestBody
}

type RequestBody struct {
	Type          uint8
	TransactionId uint16
}

func (message *RequestMessage) MessageId() uint8 {
	return 2
}

func (message *RequestMessage) ToBytes() ([]byte, error) {

	message.Header = Tlv{
		MessageType: 2,
		Version:     0,
		Length:      3,
	}

	headerBytes, err := message.Header.ToBytes()
	if err != nil {
		return []byte{}, fmt.Errorf("unable to convert request header to bytes: %s", err.Error())
	}

	bodyBytes, err := message.Body.toBytes()
	if err != nil {
		return []byte{}, fmt.Errorf("unable to convert request body to bytes: %s", err.Error())
	}

	return append(headerBytes, bodyBytes...), nil

}

func (message *RequestBody) toBytes() ([]byte, error) {

	buf := &bytes.Buffer{}

	buf.Write([]byte{message.Type})
	err := binary.Write(buf, binary.BigEndian, message.TransactionId)
	if err != nil {
		return []byte{}, fmt.Errorf(
			"could not decode transaction number %v to binary: %s",
			message.TransactionId,
			err.Error(),
		)
	}

	return ioutil.ReadAll(buf)

}
