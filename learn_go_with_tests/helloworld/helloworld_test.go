package main

import "testing"

func TestHelloworld(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		} else {
			t.Log("subtest passed")
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Helloworld("Saria", "")
		want := "Hello,Saria"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello if got an empty string", func(t *testing.T) {
		got := Helloworld("", "")
		want := "Hello,Doc"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Helloworld("Blaze", "Spanish")
		want := "Hola,Blaze"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Helloworld("Ros", "French")
		want := "Bonjour,Ros"
		assertCorrectMessage(t, got, want)
	})
}
