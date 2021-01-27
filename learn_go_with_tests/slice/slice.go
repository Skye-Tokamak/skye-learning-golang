package slice

func Sum(numa [5]int) int {
	var addall int
	for i := 0; i < 5; i++ {
		addall += numa[i]
	}
	return addall
}
