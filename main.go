package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("server started on :8888")

	// creates a lobby first
	r := &room{
		name:    "lobby",
		members: make(map[net.Addr]*client),
	}
	s.rooms["lobby"] = r

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}

		c := s.newClient(conn)
		go c.readInput()
	}
}
