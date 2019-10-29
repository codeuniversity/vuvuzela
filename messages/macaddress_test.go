package messages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldBeAbleToDecodeValidMacAddressToBytes(t *testing.T) {

	// given
	expectedBytes := []byte{0x5e, 0x5c, 0xce, 0xca, 0x93, 0x58}
	validMacAddress := MacAddress{"5e:5c:ce:ca:93:58"}

	// when
	actualBytes, err := validMacAddress.ToBytes()

	// then

	assert.Nil(t, err)
	assert.Equal(t, expectedBytes, actualBytes)

}

func Test_ShouldThrowErrorWhenInvalidMacIsProvided(t *testing.T) {

	// given
	invalidMacAddress := MacAddress{"5e:5c:ce:ca:93:58:invalid"}

	// when
	macBytes, err := invalidMacAddress.ToBytes()

	// then

	assert.NotNil(t, err)
	assert.Equal(t, []byte{}, macBytes)

}

func Test_ShouldBeAbleToGetValidMacAddressFromBytes(t *testing.T) {

	// given
	expectedMacAddress := MacAddress{"5e:5c:ce:ca:93:58"}
	macAddressBytes := []byte{0x5e, 0x5c, 0xce, 0xca, 0x93, 0x58}

	// when
	actualMacAddress, err := newMacAddressFromBytes(macAddressBytes)

	// then

	assert.Nil(t, err)
	assert.Equal(t, expectedMacAddress, actualMacAddress)

}

func Test_ShouldThrowErrorWhenInvalidMacBytesAreProvided(t *testing.T) {

	// given
	invalidMacBytes := []byte{0x5e, 0x5c, 0xce, 0xca, 0x93, 0x58, 0x00}

	// when
	macAddress, err := newMacAddressFromBytes(invalidMacBytes)

	// then

	assert.NotNil(t, err)
	assert.Equal(t, MacAddress{}, macAddress)

}
