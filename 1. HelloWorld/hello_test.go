package main

import "testing"

func TestHello(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assert(t, got, want)
	})

	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assert(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Eloa", "Spanish")
		want := "Hola, Eloa"
		assert(t, got, want)
	})

	t.Run("say hello in french", func(t *testing.T) {
		got := Hello("Marcus", "French")
		want := "Bonjour, Marcus"
		assert(t, got, want)
	})

	t.Run("say hello in german", func(t *testing.T) {
		got := Hello("Jonas", "German")
		want := "Hallo, Jonas"
		assert(t, got, want)
	})
}
