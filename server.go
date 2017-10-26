package server

import (
	"net"
	"io"
)

type server 	struct {
	address	string
	onConnected func(s *Session)
	onDisconnected func(s *Session)
}

type Session struct {
	conn net.Conn
}

func New(address string) *server {
	return &server{
		address,
		func(s *Session) {},
		func(s *Session) {}}
}

func (s *server) OnConnected(callback func(s *Session)) {
	s.onConnected = callback
}

func (s *server) OnDisconnected(callback func(s *Session)) {
	s.onDisconnected = callback
}

func (s *server) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.listen(&Session{conn})
	}
}

func (s *server) listen(session *Session) {
	s.onConnected(session)

	defer func() {
		s.onDisconnected(session)
		session.conn.Close()
	}()

	var buffer = make([]byte, 1024)

	for {
		_, err := session.conn.Read(buffer)

		if err == io.EOF || err != nil {
			return
		}
	}
}