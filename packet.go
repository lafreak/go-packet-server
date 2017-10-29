package server

import (
	"fmt"
	"encoding/binary"
	"bytes"
	"math"
)

type Packet struct {
	data []byte
}

func ToPacket(data []byte) *Packet {
	if len(data) < 3 {
		return nil
	}

	return &Packet{data}
}

func NewPacket(type_ byte) *Packet {
	return &Packet{[]byte{3, 0, type_}}
}

func (p *Packet) Size() uint16 {
	return binary.LittleEndian.Uint16(p.data[:2])
}

func (p *Packet) setSize(size uint16) {
	binary.LittleEndian.PutUint16(p.data, size)
}

func (p *Packet) addSize(size uint16) {
	p.setSize(p.Size() + size)
}

func (p *Packet) removeSize(size uint16) {
	p.setSize(p.Size() - size)
	p.data = append(p.data[:3], p.data[3+size:]...)
}

func (p *Packet) containsString() bool {
	if p.Size() <= 3 {
		return false
	}

	for _, v := range p.data[3:] {
		if v == 0 {
			return true
		}
	}
	return false
}

func (p *Packet) extractString() (uint16, string) {
	if !p.containsString() {
		return 0, ""
	}

	for i, v := range p.data[3:] {
		if v == 0 {
			return uint16(i+1), string(p.data[3:i+3])
		}
	}
	return 0, ""
}

func (p *Packet) Type() byte {
	return p.data[2]
}

func (p *Packet) Buffer() []byte {
	return p.data
}

func (p *Packet) Stringify() string {
	return string(p.data[3:])
}

func (p *Packet) String() string {
	return fmt.Sprintf("%d ", p.data)
}

func (p *Packet) Write(data ...interface{}) {
	for _, value := range data {
		switch v := value.(type) {
		case uint8:
			p.data = append(p.data, v)
			p.addSize(1)
		case int8:
			p.data = append(p.data, byte(v))
			p.addSize(1)
		case uint16:
			b := make([]byte, 2)
			binary.LittleEndian.PutUint16(b, uint16(v))
			p.data = append(p.data, b...)
			p.addSize(2)
		case int16:
			b := make([]byte, 2)
			binary.LittleEndian.PutUint16(b, uint16(v))
			p.data = append(p.data, b...)
			p.addSize(2)
		case uint32:
			b := make([]byte, 4)
			binary.LittleEndian.PutUint32(b, uint32(v))
			p.data = append(p.data, b...)
			p.addSize(4)
		case int32:
			b := make([]byte, 4)
			binary.LittleEndian.PutUint32(b, uint32(v))
			p.data = append(p.data, b...)
			p.addSize(4)
		case int:
			b := make([]byte, 4)
			binary.LittleEndian.PutUint32(b, uint32(v))
			p.data = append(p.data, b...)
			p.addSize(4)
		case uint64:
			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b, uint64(v))
			p.data = append(p.data, b...)
			p.addSize(8)
		case int64:
			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b, uint64(v))
			p.data = append(p.data, b...)
			p.addSize(8)
		case string:
			b := []byte(v)
			p.data = append(p.data, b...)
			p.data = append(p.data, 0)
			p.addSize(uint16(len(v)+1))
		case float32:
			var b bytes.Buffer
			err := binary.Write(&b, binary.LittleEndian, float32(v))
			if err == nil {
				p.data = append(p.data, b.Bytes()...)
				p.addSize(4)
			}
		case float64:
			var b bytes.Buffer
			err := binary.Write(&b, binary.LittleEndian, float64(v))
			if err == nil {
				p.data = append(p.data, b.Bytes()...)
				p.addSize(8)
			}
		}
	}
}

func (p *Packet) Read(variables ...interface{}) {
	for _, variable := range variables {
		if p.Size() <= 3 {
			return
		}

		switch v := variable.(type) {
		case *byte:
			if p.Size() >= 3 + 1 {
				*v = p.data[3]
				p.removeSize(1)
			}
		case *int8:
			if p.Size() >= 3 + 1 {
				*v = int8(p.data[3])
				p.removeSize(1)
			}
		case *uint16:
			if p.Size() >= 3 + 2 {
				*v = binary.LittleEndian.Uint16(p.data[3:3 + 2])
				p.removeSize(2)
			}
		case *int16:
			if p.Size() >= 3 + 2 {
				*v = int16(binary.LittleEndian.Uint16(p.data[3:3 + 2]))
				p.removeSize(2)
			}
		case *int:
			if p.Size() >= 3 + 4 {
				*v = int(binary.LittleEndian.Uint32(p.data[3:3 + 4]))
				p.removeSize(4)
			}
		case *uint32:
			if p.Size() >= 3 + 4 {
				*v = binary.LittleEndian.Uint32(p.data[3:3 + 4])
				p.removeSize(4)
			}
		case *int32:
			if p.Size() >= 3 + 4 {
				*v = int32(binary.LittleEndian.Uint32(p.data[3:3 + 4]))
				p.removeSize(4)
			}
		case *uint64:
			if p.Size() >= 3 + 8 {
				*v = binary.LittleEndian.Uint64(p.data[3:3 + 8])
				p.removeSize(8)
			}
		case *int64:
			if p.Size() >= 3 + 8 {
				*v = int64(binary.LittleEndian.Uint64(p.data[3:3 + 8]))
				p.removeSize(8)
			}
		case *string:
			if n, s := p.extractString(); n > 0 {
				*v = s
				p.removeSize(n)
			}
		case *float32:
			if p.Size() >= 3 + 4 {
				*v = math.Float32frombits(binary.LittleEndian.Uint32(p.data[3:3 + 4]))
				p.removeSize(4)
			}
		case *float64:
			if p.Size() >= 3 + 8 {
				*v = math.Float64frombits(binary.LittleEndian.Uint64(p.data[3:3 + 8]))
				p.removeSize(8)
			}
		}
	}
}