package connection

import (
	"github.com/stretchr/testify/assert"
	"net"
	"os"
	"testing"
)

func Test_ShouldUseAlfredSocketByDefault(t *testing.T) {

	// given

	// when
	conn := New(&Configuration{})

	// then
	assert.Equal(t, "/var/run/alfred.sock", conn.configuration.SocketFile)

}

func Test_ShouldBeAbleToSendTestPayloadToSocket(t *testing.T) {

	// given

	expectedMessage := []byte{0, 0, 0, 25, 0, 0, 0, 0, 94, 92, 206, 202, 147, 88, 65, 0, 0, 11, 116, 101, 115, 116, 109, 101, 115, 115, 97, 103, 101}

	config := Configuration{"../tmp/test.sock"}
	conn := New(&config)

	// Setup local socket

	receivedMessages := make(chan []byte)

	ln, err1 := net.Listen("unix", config.SocketFile)
	defer os.Remove(config.SocketFile)
	if err1 != nil {
		panic(err1)
	}
	go func() {
		for {
			conn2, err := ln.Accept()
			if err != nil {
				panic(err)
			}
			reader := make([]byte, 29)
			_, err = conn2.Read(reader)
			if err != nil {
				panic(err)
			}
			receivedMessages <- reader

			conn2.Close()
			ln.Close()
		}
	}()

	// when
	err := conn.Send(65, "testmessage")

	// then
	receivedMessage := <-receivedMessages
	assert.Nil(t, err)
	assert.Equal(t, expectedMessage, receivedMessage)

}

func Test_ShouldThrowErrorWhenSocketIsNotAvailable(t *testing.T) {

	// given
	config := Configuration{"../tmp/unavailable.sock"}
	conn := New(&config)

	// when
	err := conn.Send(65, "testmessage")

	// then
	assert.NotNil(t, err)

}
