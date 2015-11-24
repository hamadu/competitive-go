package main
import (
	"fmt"
	"sort"
)

func main() {
	l1, l2, l3 := 0, 0, 0
	r, b, y := 0, 0, 0
	fmt.Scan(&l1, &l2, &l3, &r, &b, &y)
	size := []int{l1+l2, l2+l3, l3+l1}
	cnt := []int{r, b, y}

	sort.Ints(size)
	sort.Ints(cnt)

	total := 0
	for i := 0; i < 3; i++ {
		total += size[i] * cnt[2-i] * 2
	}
	fmt.Println(total)
}