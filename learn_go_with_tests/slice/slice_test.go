package slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of elements", func(t *testing.T) {
		numa := []int{1, 2, 3, 4, 5}
		got := Sum(numa)
		want := 15
		if got != want {
			t.Errorf("want '%d' but got '%d' given %v", want, got, numa)
		}
	})
	t.Run("collection of arrays", func(t *testing.T) {
		numaa := []int{1, 2, 3}
		got := SumAll([]int{1, 2, 3}, []int{1, 2})
		want := []int{6, 3}
		//reflect.DeepEqual判断数组内容是否相等
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want '%v' but got '%v' given %v", want, got, numaa)
		}
	})
}
