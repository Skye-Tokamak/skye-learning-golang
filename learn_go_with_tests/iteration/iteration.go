package iteration

func Repeat(s string, n int) string {
	repeated := ""
	for i := 0; i < n; i++ {
		repeated += s
	}
	return repeated
}
