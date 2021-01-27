package slice

func Sum(numa []int) int {
	var addall int
	//range用法
	for _, num := range numa {
		addall += num
	}
	return addall
}

//可变参数这么用
func SumAll(numaa ...[]int) []int {
	var sums []int
	for _, array := range numaa {
		sums = append(sums, Sum(array))
	}

	return sums
}
