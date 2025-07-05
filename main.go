package main

import (
	"log"
	"net"
)

func main() {
	// creates a new server instance
	s := newServer()

	// listens the tcp port
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start the server: %s", err.Error())
	}

	// dont forget to close the listener
	defer listener.Close()
	log.Printf("server listened to port: 8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
