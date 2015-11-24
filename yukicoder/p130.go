package main
import (
	"fmt"
	"sort"
)

var a []int

func main() {
	var n int
	fmt.Scan(&n)

	a = make([]int, n)
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	fmt.Println(solve(1<<30, 0, n))
}

func solve(bit, low, high int) int {
	if bit == 0 {
		return 0
	}

	ln := 0
	hn := 0
	med := -1
	for i := low ; i < high ; i++ {
		if a[i] & bit >= 1 {
			if med == -1 {
				med = i
			}
			hn++
		} else {
			ln++
		}
	}
	if hn == 0 || ln == 0 {
		return solve(bit>>1, low, high)
	}
	left := solve(bit>>1, low, med)
	right := solve(bit>>1, med, high)

	if left < right {
		return left | bit
	} else {
		return right | bit
	}
}
