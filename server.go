package server

import (
	"net"
	"io"
	"bufio"
	"encoding/binary"
)

type server 	struct {
	address	string
	events map[byte]func(s *Session, p* Packet)
	onConnected func(s *Session)
	onDisconnected func(s *Session)
	onUnknownPacket func(s *Session, p* Packet)
	reader *bufio.Reader
}

type Session struct {
	conn net.Conn
}

func New(address string) *server {
	return &server{
		address,
		make(map[byte]func(s *Session, p *Packet)),
		func(s *Session) {},
		func(s *Session) {},
		func(s *Session, p *Packet) {},
		nil}
}

func (s *server) OnConnected(callback func(s *Session)) {
	s.onConnected = callback
}

func (s *server) OnDisconnected(callback func(s *Session)) {
	s.onDisconnected = callback
}

func (s *server) OnUnknownPacket(callback func(s *Session, p *Packet)) {
	s.onUnknownPacket = callback
}

func (s *server) On(type_ byte, callback func(s *Session, p *Packet)) {
	s.events[type_] = callback
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

	s.reader = bufio.NewReader(session.conn)

	for {
		packets, err := s.receive()
		if err != nil {
			return;
		}

		for _, p := range packets {
			if event, ok := s.events[p.Type()]; ok {
				event(session, p)
			} else {
				s.onUnknownPacket(session, p)
			}
		}
	}
}

func (s *server) receive() ([]*Packet, error) {
	buffer := make([]byte, 1024)
	packets := make([]*Packet, 0)

	n, err := s.reader.Read(buffer)
	if err == io.EOF || err != nil {
		return nil, err
	}

	for n > 0 {
		m := binary.LittleEndian.Uint16(buffer[:2])
		if m == 0 {
			m = 1
		}

		p := ToPacket(buffer[:m])

		n -= int(m)
		buffer = append(buffer[m:], make([]byte, m)...)

		if m < 3 {
			continue;
		}

		packets = append(packets, p)
	}

	return packets, nil
}

func (s *Session) Send(type_ byte, data ...interface{}) int {
	p := NewPacket(type_)
	p.Write(data...)
	n, _ := s.conn.Write(p.Buffer())
	return n
}
