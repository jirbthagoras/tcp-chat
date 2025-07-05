package main

import (
	"bufio"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	for {
		// reads all the message sent by user via the peer connection
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		// processes the message
		msg = strings.Trim(msg, "\r\n")

		// args := strings.Split(msg, " ")
		// cmd := strings.TrimSpace(args[0])
	}
}
