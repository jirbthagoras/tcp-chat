package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		case CMD_MEMBERS:
			s.listMembers(cmd.client)
		case CMD_WHISPER:
			s.whisper(cmd.client, cmd.args)
		}
	}
}

func (s *server) newClient(conn net.Conn) *client {
	log.Printf("new client has joined: %s", conn.RemoteAddr().String())

	client := &client{
		conn:     conn,
		nick:     randomString(6),
		commands: s.commands,
		room:     s.rooms["lobby"],
	}

	client.room.members[conn.RemoteAddr()] = client
	return client
}

func (s *server) nick(c *client, args []string) {
	if len(args) < 2 {
		c.systemMsg("nick is required. usage: /nick [nickname]")
		return
	}

	res, _ := s.checkNickname(args[1])

	if res {
		c.err(fmt.Errorf("username %s already used", args[1]))
		return
	}

	c.nick = args[1]
	c.systemMsg(fmt.Sprintf("all right, I will call you %s", c.nick))
}

func (s *server) join(c *client, args []string) {
	if len(args) < 2 {
		c.systemMsg("room name is required. usage: /join [roomName]")
		return
	}

	roomName := args[1]

	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}

	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c)
	c.room = r

	r.broadcast(c, fmt.Sprintf("%s joined the room", c.nick))

	c.msg(fmt.Sprintf("welcome to %s", roomName))
}

func (s *server) listRooms(c *client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.systemMsg(fmt.Sprintf("available rooms: %s", strings.Join(rooms, ", ")))
}

func (s *server) msg(c *client, args []string) {
	if len(args) < 2 {
		c.systemMsg("message is required, usage: /msg MSG")
		return
	}

	msg := strings.Join(args[1:], " ")
	c.room.broadcast(c, c.nick+": "+msg)
}

func (s *server) quit(c *client) {
	log.Printf("client has left the chat: %s", c.conn.RemoteAddr().String())

	s.quitCurrentRoom(c)

	c.systemMsg("bye~")
	c.conn.Close()
}

func (s *server) listMembers(c *client) {
	var members []string
	for _, m := range c.room.members {
		members = append(members, m.nick)
	}

	c.systemMsg(fmt.Sprintf("available members in room %s: %s", c.room.name, strings.Join(members, ", ")))
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		oldRoom := s.rooms[c.room.name]
		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr())
		oldRoom.broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
	}
}

// func (s *server) checkRoom(c *client) {
// 	if c.room == nil {
// 		c.err(errors.New("please join a room first with: /join [roomname]"))
// 		return
// 	}
// }

func (s *server) checkNickname(nickname string) (bool, *client) {
	for _, room := range s.rooms {
		for _, member := range room.members {
			if member.nick == nickname {
				return true, member
			}
		}
	}
	return false, &client{}
}

func (s *server) whisper(c *client, args []string) {
	// grab the message and target
	msg := strings.Join(args[2:], " ")
	target := args[1]

	res, _ := s.checkNickname(target)
	if !res {
		c.err(fmt.Errorf("target %s does not exists in server", target))
		return
	}

	c.room.whisper(c, target, msg)
}

