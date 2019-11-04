package messages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldBeAbleToBuildValidRequestMessageBody(t *testing.T) {

	// given
	expectedBodyBytes := []byte{99, 0, 12}
	requestBody := RequestBody{
		Type:          99,
		TransactionId: 12,
	}

	// when
	actualBodyBytes, err := requestBody.toBytes()

	// then
	assert.Nil(t, err)
	assert.Equal(t, expectedBodyBytes, actualBodyBytes)

}

func Test_ShouldBeAbleToBuildValidRequestMessage(t *testing.T) {

	// given
	expectedMessageBytes := []byte{2, 0, 0, 3, 98, 0, 13}
	requestMessage := RequestMessage{
		Header: Tlv{
			MessageType: 2,
			Version:     0,
			Length:      3,
		},
		Body: RequestBody{
			Type:          98,
			TransactionId: 13,
		},
	}

	// when
	actualMessageBytes, err := requestMessage.ToBytes()

	// then
	assert.Nil(t, err)
	assert.Equal(t, expectedMessageBytes, actualMessageBytes)

}
