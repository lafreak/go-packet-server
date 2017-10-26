package main

import (
	"fmt"
	"github.com/lafreak/go-event-server"
)

func main() {
	socket := server.New("localhost:3000")

	socket.OnConnected(func(s *server.Session) {
		fmt.Println("Client connected.")
	})

	socket.OnDisconnected(func(s *server.Session) {
		fmt.Println("Client disconnected.")
	})

	socket.Start()
}