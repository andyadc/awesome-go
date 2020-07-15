package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello2("andy", "Spanish")
	want := "Hello, andy"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("World")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}
