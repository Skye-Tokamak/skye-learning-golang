package injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("test greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Blaze")

		got := buffer.String()
		want := "Hello,Blaze"
		assertString(t, got, want)

	})
}
func assertString(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
