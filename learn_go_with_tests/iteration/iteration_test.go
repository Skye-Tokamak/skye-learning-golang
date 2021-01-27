package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("want '%q' but got '%q'", want, got)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
func ExampleRepeat() {
	repeat := Repeat("a2", 3)
	fmt.Println(repeat)
	//Output:a2a2a2
}
