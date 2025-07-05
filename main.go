package main

import (
	"log"
	"net"
)

func main() {
	// creates a new server instance
	// s := newServer()

	// listens the tcp port
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start the server: %s", err.Error())
	}

	// dont forget to close the listener
	defer listener.Close()
}
