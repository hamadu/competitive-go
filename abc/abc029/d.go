package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	fmt.Scan(&n)

	total := 0
	idx := 0
	for k := 1 ; ; k *= 10 {
		if k >= n * 10 {
			break
		}
		dp := make([][]int, 12)
		for i := range dp {
			dp[i] = make([]int, 2)
		}

		dp[11][0] = 1
		for i := 11 ; i >= 1 ; i-- {
			to := i-1
			for of := 0 ; of <= 1 ; of++ {
				base := dp[i][of]
				for d := 0 ; d <= 9 ; d++ {
					if to == idx && d != 1 {
						continue
					}
					ov := over(n, d, to)
					if of == 0 && ov == 1 {
						continue
					}
					tf := of
					if tf == 0 && ov == -1 {
						tf = 1
					}
					dp[to][tf] += base
				}
			}
		}
		total += dp[0][0] + dp[0][1]
		idx++
	}
	fmt.Println(total)
}


func over(n, k, d int) int {
	x := k * int(math.Pow10(d))
	y := n % int(math.Pow10(d+1))
	y /= int(math.Pow10(d))
	y *= int(math.Pow10(d))
	if x > y {
		return 1
	} else if x < y {
		return -1
	}
	return 0
}