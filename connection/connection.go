package connection

import (
	"fmt"
	"github.com/codeuniversity/vuvuzela/messages"
	"net"
)

type Connection struct {
	configuration *Configuration
}

type Configuration struct {
	SocketFile string
}

func New(configuration *Configuration) Connection {
	if configuration.SocketFile == "" {
		configuration.SocketFile = "/var/run/alfred.sock"
	}
	return Connection{configuration}
}

func (connection *Connection) Send(typeId uint8, message string) error {

	messageToSend := messages.PushDataMessage{
		Header: messages.Tlv{},
		Body: messages.PushDataBody{
			TransactionId:  0,
			SequenceNumber: 0,
			MacAddress:     messages.MacAddress{Value: "5e:5c:ce:ca:93:58"},
			Type:           typeId,
			Version:        0,
			Length:         0,
			Data:           message,
		},
	}

	conn, err := net.Dial("unix", connection.configuration.SocketFile)
	if err != nil {
		return fmt.Errorf("cannot connect to Alfred server, are you sure it is running? %s", err.Error())
	}

	defer conn.Close()

	messageBytes, err := messageToSend.ToBytes()
	if err != nil {
		return fmt.Errorf("cannot assemble message to send. please make sure you're providing a valid message: %s", err.Error())
	}

	_, err = conn.Write(messageBytes)
	if err != nil {
		return fmt.Errorf("cannot send message to socket: %s", err.Error())
	}

	return nil
}
