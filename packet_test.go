package server

import "testing"

func TestSetSize(t *testing.T) {
	p := Packet{[]byte{9, 0, 1}}
	p.setSize(32)

	if p.Buffer()[0] != 32 {
		t.Error("Expected byte 32, got", p.Buffer()[0])
	}

	if p.Buffer()[1] != 0 {
		t.Error("Expected byte 0, got", p.Buffer()[1])
	}

	p.setSize(829)

	if p.Buffer()[0] != 61 {
		t.Error("Expected byte 61, got", p.Buffer()[0])
	}

	if p.Buffer()[1] != 3 {
		t.Error("Expected byte 3, got", p.Buffer()[1])
	}
}

func TestAddSize(t *testing.T) {
	p := Packet{[]byte{9, 0, 1}}
	p.addSize(272)

	if p.Buffer()[0] != 25 {
		t.Error("Expected byte 25, got", p.Buffer()[0])
	}

	if p.Buffer()[1] != 1 {
		t.Error("Expected byte 1, got", p.Buffer()[1])
	}
}

func TestSize(t *testing.T) {
	p := Packet{[]byte{3, 0, 1}}

	if p.Size() != 3 {
		t.Error("Expected size 3, got", p.Size())
	}

	p = Packet{[]byte{8, 13, 2}}

	if p.Size() != 3336 {
		t.Error("Expected size 3336, got", p.Size())
	}
}

func TestType(t *testing.T) {
	p := Packet{[]byte{3, 0, 19}}

	if p.Type() != 19 {
		t.Error("Expected type 19, got", p.Type())
	}
}

func TestContainsString(t *testing.T) {
	p := Packet{[]byte{5, 0, 1, 4, 0}}

	if !p.containsString() {
		t.Error("Expected true")
	}

	p = Packet{[]byte{5, 0, 1, 4, 2}}

	if p.containsString() {
		t.Error("Expected false")
	}

	p = Packet{[]byte{3, 0, 1}}

	if p.containsString() {
		t.Error("Expected false")
	}
}

func TestExtractString(t *testing.T) {
	p := Packet{[]byte{6, 0, 1, 'G', 'o', 0}}

	if n, s := p.extractString(); s != "Go" || n != 3 {
		t.Error("Expected string Go & 3, got", s, n)
	}

	p = Packet{[]byte{9, 0, 1, 'k', 'a', 'l', 0, 'a', 0}}

	if _, s := p.extractString(); s != "kal" {
		t.Error("Expected string kal, got", s)
	}

	p = Packet{[]byte{3, 0, 1}}

	if n, s := p.extractString(); s != "" || n != 0 {
		t.Error("Expected empty string & 0, got", s, n)
	}

	p = Packet{[]byte{4, 0, 1, 0}}

	if n, s := p.extractString(); s != "" || n != 1 {
		t.Error("Expected empty string & 1, got", s, n)
	}
}