package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to GO", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountDownSpyer{})

		got := buffer.String()
		want := `3
2
1
GO`
		assertString(t, got, want)
	})
	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountDownSpyer{}
		Countdown(spySleepPrinter, spySleepPrinter)
		got := spySleepPrinter.Calls
		want := `write
sleep
write
sleep
write
sleep
write`
		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
func assertString(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
