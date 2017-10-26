package main

import (
	"fmt"
	"github.com/lafreak/go-event-server"
)

/*
	This example takes you to character selection screen when using game client of game KalOnline.
 */

func main() {
	game := server.New("localhost:3000")

	game.OnConnected(func(s *server.Session) {
		fmt.Println("Client connected.")
	})

	game.OnDisconnected(func(s *server.Session) {
		fmt.Println("Client disconnected.")
	})

	game.OnUnknownPacket(func(s *server.Session, p *server.Packet) {
		fmt.Println("Unknown packet:", p.Type())
	})

	// C2S_CONNECT
	game.On(9, func(s *server.Session, p *server.Packet) {
		// S2C_CODE
		s.Send(125, 0, byte(0), 604800, 0, 0, uint64(0), byte(0), byte(0), byte(2))
	})

	// C2S_ANS_CODE
	game.On(4, func(s *server.Session, p *server.Packet) {})

	// C2S_LOGIN
	game.On(8, func(s *server.Session, p *server.Packet) {
		var login, password, mac string
		p.Read(&login, &password, &mac)

		fmt.Println(login, password, mac)

		// S2C_ANS_LOGIN
		s.Send(149, byte(1))
	})

	// C2S_SECOND_LOGIN
	game.On(10, func(s *server.Session, p *server.Packet) {
		// S2C_PLAYERINFO
		s.Send(114, byte(0), byte(0), 0, byte(1), 1, "Liplay", byte(4), byte(11), byte(60), 0, uint16(5), uint16(5), uint16(5), uint16(5), uint16(5), byte(0), byte(0), byte(0))
	})

	err := game.Start()
	if err != nil {
		fmt.Println(err)
	}
}