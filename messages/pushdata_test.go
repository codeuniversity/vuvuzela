package messages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldBeAbleToBuildValidMessage(t *testing.T) {

	// given
	expectedMessage := []byte{0, 0, 0, 26, 0, 0, 0, 0, 94, 92, 206, 202, 147, 88, 99, 0, 0, 12, 114, 97, 115, 112, 98, 101, 114, 114, 121, 112, 105, 10}

	message := PushDataMessage{
		Header: Tlv{},
		Body: PushDataBody{
			TransactionId:  0,
			SequenceNumber: 0,
			MacAddress:     MacAddress{Value: "5e:5c:ce:ca:93:58"},
			Type:           99,
			Version:        alfredVersion,
			Length:         0,
			Data:           "raspberrypi\n",
		},
	}

	// when

	messageBytes, err := message.ToBytes()

	// then

	assert.Nil(t, err)
	assert.Equal(t, expectedMessage, messageBytes)
}

func Test_ShouldBeAbleToGetValidBodyFromBytes(t *testing.T) {

	// given
	bodyBytes := []byte{0, 0, 0, 0, 94, 92, 206, 202, 147, 88, 99, 0, 0, 12, 114, 97, 115, 112, 98, 101, 114, 114, 121, 112, 105, 10}

	expectedBody := PushDataBody{
		TransactionId:  0,
		SequenceNumber: 0,
		MacAddress:     MacAddress{Value: "5e:5c:ce:ca:93:58"},
		Type:           99,
		Version:        alfredVersion,
		Length:         12,
		Data:           "raspberrypi\n",
	}

	// when

	actualBody, err := newPushDataBodyFromBytes(bodyBytes)

	// then

	assert.Nil(t, err)
	assert.Equal(t, expectedBody, actualBody)
}

func Test_ShouldBeAbleToGetValidMessageFromBytes(t *testing.T) {

	// given
	bytes := []byte{0, 0, 0, 26, 0, 0, 0, 0, 94, 92, 206, 202, 147, 88, 99, 0, 0, 12, 114, 97, 115, 112, 98, 101, 114, 114, 121, 112, 105, 10}

	expectedMessage := PushDataMessage{
		Tlv{
			MessageType: 0,
			Version:     0,
			Length:      26,
		},
		PushDataBody{
			TransactionId:  0,
			SequenceNumber: 0,
			MacAddress:     MacAddress{Value: "5e:5c:ce:ca:93:58"},
			Type:           99,
			Version:        alfredVersion,
			Length:         12,
			Data:           "raspberrypi\n",
		}}

	// when

	actualMessage, err := NewPushDataMessageFromBytes(bytes)

	// then

	assert.Nil(t, err)
	assert.Equal(t, expectedMessage, actualMessage)
}
