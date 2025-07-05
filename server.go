package main

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
