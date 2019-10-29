package messages

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type MacAddress struct {
	Value string
}

func (macAddress *MacAddress) ToBytes() ([]byte, error) {

	strippedMacAddress := strings.ReplaceAll(macAddress.Value, ":", "")
	result, err := hex.DecodeString(strippedMacAddress)
	if err != nil {
		return []byte{}, err
	} else if len(result) != 6 {
		return []byte{}, fmt.Errorf("invalid MAC address format: %s", macAddress.Value)
	}
	return result, nil
}

func newMacAddressFromBytes(bytes []byte) (MacAddress, error) {

	if len(bytes) != 6 {
		return MacAddress{}, fmt.Errorf("cannot encode MAC address because of invalid length %v", len(bytes))
	}

	var result string

	for i, b := range bytes {
		result += hex.EncodeToString([]byte{b})
		if i != len(bytes)-1 {
			result += ":"
		}
	}

	return MacAddress{result}, nil

}
