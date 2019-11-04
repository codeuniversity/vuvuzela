package main

import (
	"github.com/lennartschoch/vuvuzela/connection"
	"os"
)

func main() {
	// TODO: Remove test code and start GRPC server instead

	if len(os.Args) > 1 {

		arg := os.Args[1]
		if arg == "send" {
			config := connection.Configuration{}
			conn := connection.New(&config)
			err := conn.Send(65, "testmessage")
			if err != nil {
				panic(err)
			}
		}

	}
}
