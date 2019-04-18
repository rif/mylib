package mylib

import "testing"

func TestGenerateMessage(t *testing.T) {
	message := GenerateMessage(1)
	expect := "hello world 1"
	if message != expect {
		t.Errorf("expected %s got %s", expect, message)
	}
}
