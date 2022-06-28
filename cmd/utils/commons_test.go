package utils

import "testing"

func TestInterfaceAsInteger(t *testing.T) {
	if 2 != InterfaceAsInteger("2") {
		t.Errorf("Unable to convert interface to integer")
	}
}

func TestContains(t *testing.T) {
	if !Contains([]string{"1", "2", "3"}, "1") {
		t.Error("Contains logic failure")
	}
}
