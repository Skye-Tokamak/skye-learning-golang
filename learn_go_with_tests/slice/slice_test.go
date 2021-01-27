package slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checkSum := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want '%v' but got '%v'", want, got)
		}
	}
	t.Run("collection of elements", func(t *testing.T) {
		numa := []int{1, 2, 3, 4, 5}
		got := Sum(numa)
		want := 15
		if got != want {
			t.Errorf("want '%d' but got '%d' given %v", want, got, numa)
		}
	})
	t.Run("collection of arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 2})
		want := []int{6, 3}
		//reflect.DeepEqual判断数组内容是否相等
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want '%v' but got '%v' ", want, got)
		}
	})
	t.Run("collection of elements tail", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 2})
		want := []int{5, 2}
		checkSum(t, got, want)

	})
	t.Run("collection of empty exist!", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2})
		want := []int{0, 2}
		checkSum(t, got, want)
	})
}
