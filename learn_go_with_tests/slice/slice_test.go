package slice

import "testing"

func TestSum(t *testing.T) {
	numa := [5]int{1, 2, 3, 4, 5}
	got := Sum(numa)
	want := 15
	if got != want {
		t.Errorf("want '%d' but got '%d' given %v", want, got, numa)
	}
}
