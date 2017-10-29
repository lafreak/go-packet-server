package server

import "testing"

func TestReadByte(t *testing.T) {
	p := Packet{[]byte{5, 0, 1, 8, 98}}

	var val1, val2 byte
	p.Read(&val1, &val2)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val1 != 8 {
		t.Fatal("Expected value 8, got", val1)
	}

	if val2 != 98 {
		t.Fatal("Expected value 98, got", val2)
	}
}

func TestReadInt8(t *testing.T) {
	p := Packet{[]byte{4, 0, 1, 255}}

	var val int8
	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val != -1 {
		t.Fatal("Expected value -1, got", val)
	}
}

func TestReadUint16(t *testing.T) {
	p := Packet{[]byte{5, 0, 1, 3, 7}}

	var val uint16
	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val != 1795 {
		t.Fatal("Expected value 1795, got", val)
	}
}

func TestReadInt16(t *testing.T) {
	p := Packet{[]byte{5, 0, 1, 254, 255}}

	var val int16
	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val != -2 {
		t.Fatal("Expected value -2, got", val)
	}
}

func TestReadUint32(t *testing.T){
	p := Packet{[]byte{7, 0, 1, 14, 0, 0, 0}}

	var val uint32
	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val != 14 {
		t.Fatal("Expected value 14, got", val)
	}
}

func TestReadInt32(t *testing.T) {
	p := Packet{[]byte{7, 0, 1, 4, 0, 0, 0}}

	var val int32
	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if len(p.Buffer()) != 3 {
		t.Fatal("Expected len 3, got", len(p.Buffer()))
	}

	if p.Buffer()[0] != 3 {
		t.Fatal("Expected byte 3, got", p.Buffer()[0])
	}

	if p.Buffer()[1] != 0 {
		t.Fatal("Expected byte 0, got", p.Buffer()[1])
	}

	if p.Buffer()[2] != 1 {
		t.Fatal("Expected byte 1, got", p.Buffer()[2])
	}

	if val != 4 {
		t.Fatal("Expected value 4, got", val)
	}
}

func TestReadInt(t *testing.T) {
	p := Packet{[]byte{7, 0, 1, 4, 0, 0, 0}}

	var val int
	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if len(p.Buffer()) != 3 {
		t.Fatal("Expected len 3, got", len(p.Buffer()))
	}

	if p.Buffer()[0] != 3 {
		t.Fatal("Expected byte 3, got", p.Buffer()[0])
	}

	if p.Buffer()[1] != 0 {
		t.Fatal("Expected byte 0, got", p.Buffer()[1])
	}

	if p.Buffer()[2] != 1 {
		t.Fatal("Expected byte 1, got", p.Buffer()[2])
	}

	if val != 4 {
		t.Fatal("Expected value 4, got", val)
	}
}

func TestReadUint64(t *testing.T) {
	p := Packet{[]byte{11, 0, 1, 255, 255, 255, 255, 255, 255, 255, 255}}

	var val uint64
	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val != 18446744073709551615 {
		t.Fatal("Expected value 18446744073709551615, got:", val)
	}
}

func TestReadInt64(t *testing.T) {
	p := Packet{[]byte{11, 0, 1, 254, 255, 255, 255, 255, 255, 255, 255}}

	var val int64
	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val != -2 {
		t.Fatal("Expected value -2, got:", val)
	}
}

func TestReadString(t *testing.T) {
	p := Packet{[]byte{6, 0, 1, 'G', 'o', 0}}

	var val string
	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val != "Go" {
		t.Fatal("Expected value Go, got:", val)
	}

	p = Packet{[]byte{9, 0, 1, 'k', 'a', 'l', 0, 'a', 0}}

	p.Read(&val)

	if p.Size() != 5 {
		t.Fatal("Expected size 5, got", p.Size())
	}

	if val != "kal" {
		t.Fatal("Expected value kal, got:", val)
	}

	p = Packet{[]byte{4, 0, 1, 0}}

	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val != "" {
		t.Fatal("Expected empty string, got:", val)
	}

	p = Packet{[]byte{3, 0, 1}}

	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val != "" {
		t.Fatal("Expected empty string, got:", val)
	}
}

func TestReadFloat32(t *testing.T) {
	p := Packet{[]byte{7, 0, 1, 205, 204, 12, 64}}

	var val float32
	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val != 2.2 {
		t.Fatal("Expected value 2.2, got", val)
	}
}

func TestReadFloat64(t *testing.T) {
	p := Packet{[]byte{11, 0, 1, 0x71, 0x3d, 0x0a, 0xd7, 0xa3, 0x70, 0xcd, 0x3f}}

	var val float64
	p.Read(&val)

	if p.Size() != 3 {
		t.Fatal("Expected size 3, got", p.Size())
	}

	if val != 0.23 {
		t.Fatal("Expected val 0.23, got", val)
	}
}