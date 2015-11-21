package main

import(
	"fmt"
)

func readIntN(n int) []int {
	ret := make([]int, n)
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&ret[i])
	}
	return ret
}

func sum(a []int) int {
	sum := 0
	for _,v := range a {
		sum += v
	}
	return sum
}

func solve(a []int) int {
	n := len(a)
	total := sum(a)
	if total % n != 0 {
		return -1
	}
	average := total / n

	need := 0
	for i := 1 ; i < n ; i++ {
		left := sum(a[:i])
		right := sum(a[i:])
		if left != average * i || right != average * (n-i) {
			need += 1
		}
	}
	return need
}

func main() {
	var n int
	fmt.Scan(&n)
	a := readIntN(n)
	fmt.Println(solve(a))
}