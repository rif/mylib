package mylib

import "testing"

func TestGenerateMessage(t *testing.T) {
	message := GenerateMessage("world", 1)
	expect := "hello world 1"
	if message != expect {
		t.Errorf("expected %s got %s", expect, message)
	}
}

func TestGenerateMessageName(t *testing.T) {
	message := GenerateMessage("john", 2)
	expect := "hello john 2"
	if message != expect {
		t.Errorf("expected %s got %s", expect, message)
	}
}
