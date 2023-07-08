package main

import "testing"

func TestHello(t *testing.T) {
	gott := Hello("Fathul")
	wantt := "Hello, Fathul"

	if gott != wantt {
		t.Errorf("got %q want %q", gott, wantt)
	}
}
