package messages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldBeAbleToDecodeValidTlvToBytes(t *testing.T) {

	// given
	tlv := Tlv{
		MessageType: 0,
		Version:     0,
		Length:      12,
	}

	// when
	tlvBytes, err := tlv.ToBytes()

	// then
	assert.Nil(t, err)
	assert.Equal(t, []byte{0, 0, 0, 12}, tlvBytes)

}
