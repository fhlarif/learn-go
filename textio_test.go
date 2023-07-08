package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Fathul")
		want := "Assalam! Apa Khabar, Fathul"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Assalam! Apa Khabar' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Assalam! Apa Khabar!"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
