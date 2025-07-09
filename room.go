package main

import (
	"fmt"
	"net"
)

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) whisper(sender *client, target string, msg string) {
	for _, member := range r.members {
		if member.nick == target {
			member.msg(fmt.Sprintf("%s (whisper): %s", sender.nick, msg))
			sender.msg(fmt.Sprintf("you whispered %s: %s", sender.nick, msg))
		}
	}
}

func (r *room) broadcast(sender *client, msg string) {
	for addr, m := range r.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}
