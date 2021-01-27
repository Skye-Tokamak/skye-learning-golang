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

//取切片用[i,j]
func SumAllTails(numaa ...[]int) []int {
	var sums []int
	for _, array := range numaa {
		if len(array) == 0 {
			sums = append(sums, 0)
		} else {
			tail := array[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return sums
}
