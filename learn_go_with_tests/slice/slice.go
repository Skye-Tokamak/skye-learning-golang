package slice

func Sum(numa [5]int) int {
	var addall int
	//range用法
	for _, num := range numa {
		addall += num
	}
	return addall
}
