package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("it should return a greeting with name when passed a name arg", func(t *testing.T) {
		got := Hello("Asa", "")
		want := "Hello, Asa"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})

	t.Run("it should return 'Hello, world' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})

	t.Run("it should return 'Bonjour, %s' when passed French as an arg", func(t *testing.T) {
		got := Hello("Asa", "french")
		want := "Bonjour, Asa"

		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
