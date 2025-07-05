package main

import "net"

type client struct {
	conn    net.Conn
	nick    string
	room    *room
	command chan<- command
}
