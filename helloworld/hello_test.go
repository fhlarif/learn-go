package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Fathul", "")
		want := "Assalam, Fathul"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Assalam! Apa Khabar' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Assalam!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Kelantanese", func(t *testing.T) {
		got := Hello("Fathul", "Kelantenese")
		want := "Guano, Fathul"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Fathul", "Japanese")
		want := "Konichiwa, Fathul"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
