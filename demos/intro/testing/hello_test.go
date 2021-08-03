package main

import "testing"

// go test -cover
func TestName(t *testing.T) {
	name := getName()

	if name != "World!" {
		t.Error("Response from getName is unexpected value")
	}
}
