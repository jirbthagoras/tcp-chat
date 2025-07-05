package main

import (
	"log"
	"net"
)

type server struct {
	rooms   map[string]*room
	command chan command
}

func newServer() *server {
	return &server{
		rooms:   make(map[string]*room),
		command: make(chan command),
	}
}

func (s *server) newClient(conn net.Conn) *client {
	log.Printf("new client has connected: %s", conn.RemoteAddr().String())

	return &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.command,
	}
}
