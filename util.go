package competitive_go

func resize(a []int, size int) {
	b := make([]int, size)
	for i := range b {
		b[i] = a[i]
	}
	return b
}