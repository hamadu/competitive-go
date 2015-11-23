package main

import "fmt"

const INF = 114514

func main()  {
	n := 0
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	best := -INF
	for i := 0 ; i < n ; i++ {
		aoki_best := -INF
		taka_best := -INF
		for j := 0 ; j < n ; j++ {
			if i == j {
				continue
			}
			aoki := 0
			taka := 0
			flg := 0
			for k := min(i, j) ; k <= max(i, j) ; k++ {
				if flg == 0 {
					taka += a[k]
				} else {
					aoki += a[k]
				}
				flg = flg^1
			}
			if aoki_best < aoki {
				aoki_best = aoki
				taka_best = taka
			}
		}
		best = max(best, taka_best)
	}
	fmt.Println(best)
}


func max(a, b int) int {
	if a < b  {
		return b
	} else {
		return a
	}
}

func min(a, b int) int {
	if a > b  {
		return b
	} else {
		return a
	}
}