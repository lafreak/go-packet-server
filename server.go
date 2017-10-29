package server

import (
	"net"
	"io"
	"bufio"
	"encoding/binary"
	"github.com/satori/go.uuid"
)

type server 	struct {
	address	string
	events map[byte]func(s *Session, p* Packet)
	sessions map[string]*Session
	onConnected func(s *Session)
	onDisconnected func(s *Session)
	onUnknownPacket func(s *Session, p* Packet)
}

type Session struct {
	conn net.Conn
	reader *bufio.Reader
	id uuid.UUID
}

func New(address string) *server {
	return &server{
		address,
		make(map[byte]func(s *Session, p *Packet)),
		make(map[string]*Session),
		func(s *Session) {},
		func(s *Session) {},
		func(s *Session, p *Packet) {}}
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
		go s.listen(&Session{
			conn,
			bufio.NewReader(conn),
			uuid.NewV4()})
	}
}

func (s *server) listen(session *Session) {
	s.sessions[session.Id()] = session
	s.onConnected(session)

	defer func() {
		s.onDisconnected(session)
		delete(s.sessions, session.Id())
		session.conn.Close()
	}()

	for {
		packets, err := session.receive()
		if err != nil {
			return
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

func (s *Session) receive() ([]*Packet, error) {
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
		if m > 1024 {
			m = 1024
		}

		p := ToPacket(buffer[:m])

		n -= int(m)
		buffer = append(buffer[m:], make([]byte, m)...)

		if m < 3 {
			continue
		}

		packets = append(packets, p)
	}

	return packets, nil
}

func (s *Session) Send(type_ byte, data ...interface{}) int {
	p := NewPacket(type_)
	p.Write(data...)
	return s.SendPacket(p)
}

func (s *Session) SendPacket(p *Packet) int {
	n, _ := s.conn.Write(p.Buffer())
	return n
}

func (s *server) SendToAll(type_ byte, data ...interface{}) {
	p := NewPacket(type_)
	p.Write(data...)
	s.SendPacketToAll(p)
}

func (s *server) SendPacketToAll(p *Packet) {
	for _, session := range s.sessions {
		session.SendPacket(p)
	}
}

func (s *Session) Id() string {
	return s.id.String()
}