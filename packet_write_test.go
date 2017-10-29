package server

import "testing"

func TestInt8(t *testing.T) {
	p := Packet{[]byte{3, 0, 1}}

	p.Write(uint8(10))

	if p.Size() != 4 {
		t.Fatal("Expected size 4, got", p.Size())
	}

	if p.Buffer()[3] != 10 {
		t.Fatal("Expected byte 10, got", p.Buffer()[3])
	}

	p.Write(byte(12))

	if p.Size() != 5 {
		t.Fatal("Expected size 5, got", p.Size())
	}

	if p.Buffer()[4] != 12 {
		t.Fatal("Expected byte 12, got", p.Buffer()[4])
	}

	p.Write(int8(-1))

	if p.Size() != 6 {
		t.Fatal("Expected size 6, got", p.Size())
	}

	if p.Buffer()[5] != 255 {
		t.Fatal("Expected byte 255, got", p.Buffer()[5])
	}
}

func TestInt16(t *testing.T) {
	p := Packet{[]byte{3, 0, 1}}

	p.Write(uint16(721))

	if p.Size() != 5 {
		t.Fatal("Expected size 5, got", p.Size())
	}

	if p.Buffer()[3] != 209 {
		t.Fatal("Expected byte 209, got", p.Buffer()[3])
	}

	if p.Buffer()[4] != 2 {
		t.Fatal("Expected byte 2, got", p.Buffer()[4])
	}

	p.Write(int16(-1))

	if p.Size() != 7 {
		t.Fatal("Expected size 7, got", p.Size())
	}

	for i := 5; i < 7; i++ {
		if p.Buffer()[i] != 255 {
			t.Fatal("Expected byte 255, got", p.Buffer()[i])
		}
	}
}

func TestInt32(t *testing.T) {
	p := Packet{[]byte{3, 0, 1}}

	p.Write(uint32(18))

	if p.Size() != 7 {
		t.Fatal("Expected size 7, got", p.Size())
	}

	if p.Buffer()[3] != 18 {
		t.Fatal("Expected byte 18, got", p.Buffer()[3])
	}

	for i := 4; i < 7; i++ {
		if p.Buffer()[i] != 0 {
			t.Fatal("Expected byte 0, got", p.Buffer()[i])
		}
	}

	p.Write(int32(-100))

	if p.Size() != 11 {
		t.Fatal("Expected size 11, got", p.Size())
	}

	if p.Buffer()[7] != 156 {
		t.Fatal("Expected byte 156, got", p.Buffer()[7])
	}

	for i := 8; i < 11; i++ {
		if p.Buffer()[i] != 255 {
			t.Fatal("Expected byte 255, got", p.Buffer()[i])
		}
	}
}

func TestInt64(t *testing.T) {
	p := Packet{[]byte{3, 0, 1}}

	p.Write(uint64(13))

	if p.Size() != 11 {
		t.Fatal("Expected size 11, got", p.Size())
	}

	if p.Buffer()[3] != 13 {
		t.Fatal("Expected byte 13, got", p.Buffer()[3])
	}

	for i := 4; i < 11; i++ {
		if p.Buffer()[i] != 0 {
			t.Fatal("Expected byte 0, got", p.Buffer()[i])
		}
	}

	p.Write(int64(-20))

	if p.Size() != 19 {
		t.Fatal("Expected size 19, got", p.Size())
	}

	if p.Buffer()[11] != 236 {
		t.Fatal("Expected byte 236, got", p.Buffer()[11])
	}

	for i := 12; i < 19; i++ {
		if p.Buffer()[i] != 255 {
			t.Fatal("Expected byte 255, got", p.Buffer()[i])
		}
	}
}

func TestString(t *testing.T) {
	p := Packet{[]byte{3, 0, 1}}

	p.Write("He llo")

	if p.Size() != 10 {
		t.Fatal("Expected size 10, got", p.Size())
	}

	if p.Buffer()[3] != 72 {
		t.Fatal("Expected byte 72('H'), got", p.Buffer()[3])
	}
	if p.Buffer()[4] != 101 {
		t.Fatal("Expected byte 101('e'), got", p.Buffer()[4])
	}
	if p.Buffer()[5] != 32 {
		t.Fatal("Expected byte 32(' '), got", p.Buffer()[5])
	}
	if p.Buffer()[6] != 108 {
		t.Fatal("Expected byte 108('l'), got", p.Buffer()[6])
	}
	if p.Buffer()[7] != 108 {
		t.Fatal("Expected byte 108('l'), got", p.Buffer()[7])
	}
	if p.Buffer()[8] != 111 {
		t.Fatal("Expected byte 111('o'), got", p.Buffer()[8])
	}
	if p.Buffer()[9] != 0 {
		t.Fatal("Expected byte 0, got", p.Buffer()[9])
	}
}

func TestFloat32(t *testing.T) {
	p := Packet{[]byte{3, 0, 1}}

	p.Write(float32(10.25))

	if p.Size() != 7 {
		t.Fatal("Expected size 7, got", p.Size())
	}

	if p.Buffer()[3] != 0 {
		t.Fatal("Expected byte 0, got", p.Buffer()[3])
	}
	if p.Buffer()[4] != 0 {
		t.Fatal("Expected byte 0, got", p.Buffer()[4])
	}
	if p.Buffer()[5] != 36 {
		t.Fatal("Expected byte 36, got", p.Buffer()[5])
	}
	if p.Buffer()[6] != 65 {
		t.Fatal("Expected byte 65, got", p.Buffer()[6])
	}
}

func TestFloat64(t *testing.T) {
	p := Packet{[]byte{3, 0, 1}}

	p.Write(10.3195123)

	if p.Size() != 11 {
		t.Fatal("Expected size 11, got", p.Size())
	}

	if p.Buffer()[3] != 232 {
		t.Fatal("Expected byte 232, got", p.Buffer()[3])
	}
	if p.Buffer()[4] != 86 {
		t.Fatal("Expected byte 86, got", p.Buffer()[4])
	}
	if p.Buffer()[5] != 190 {
		t.Fatal("Expected byte 190, got", p.Buffer()[5])
	}
	if p.Buffer()[6] != 29 {
		t.Fatal("Expected byte 29, got", p.Buffer()[6])
	}
	if p.Buffer()[7] != 151 {
		t.Fatal("Expected byte 151, got", p.Buffer()[7])
	}
	if p.Buffer()[8] != 163 {
		t.Fatal("Expected byte 163, got", p.Buffer()[8])
	}
	if p.Buffer()[9] != 36 {
		t.Fatal("Expected byte 36, got", p.Buffer()[9])
	}
	if p.Buffer()[10] != 64 {
		t.Fatal("Expected byte 64, got", p.Buffer()[10])
	}
}