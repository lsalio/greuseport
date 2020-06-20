package greuseport

import (
	"testing"
)

func TestListen(t *testing.T) {
	if _, err := Listen("tcp", ":18098"); err != nil {
		t.Fatal(err)
	}

	if _, err := Listen("tcp", ":18098"); err != nil {
		t.Fatal(err)
	}
}

func TestListenPacket(t *testing.T) {
	if _, err := ListenPacket("udp", ":18099"); err != nil {
		t.Fatal(err)
	}

	if _, err := ListenPacket("udp", ":18099"); err != nil {
		t.Fatal(err)
	}
}
