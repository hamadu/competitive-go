package main
import (
	"fmt"
)

var age int
var P [3]int
var dp [][]float64

func main() {
	fmt.Scan(&age, &P[0], &P[1], &P[2])
	age = 80 - age

	dp = make([][]float64, 2)
	dp[0] = make([]float64, 1<<14)
	dp[1] = make([]float64, 1<<14)
	dp[0][0] = 1.0

	for i := 0 ; i < age ; i++ {
		fi := i % 2
		ti := 1 - fi
		for p := 0 ; p < (1<<14) ; p++ {
			dp[ti][p] = 0.0
		}
		for p := 0 ; p < (1<<14) ; p++ {
			if dp[fi][p] == 0 {
				continue
			}
			dfs(ti, dp[fi][p], uint(p), 0, 0)
		}
	}

	exp := 0.0
	for p := 0 ; p < (1<<14) ; p++ {
		cnt := 0
		for i := uint(0) ; i < 14 ; i++ {
			if p & (1<<i) == 0 {
				cnt++
			}
		}
		exp += float64(cnt) * dp[age%2][p]
	}
	exp *= 2

	fmt.Printf("%.9f\n", exp)
}

func dfs(ti int, rate float64, ptn, tptn, idx uint) {
	if idx == 14 {
		dp[ti][tptn] += rate
		return
	}
	if ptn & (1<<idx) >= 1 {
		dfs(ti, rate, ptn, tptn | (1<<idx), idx+1)
	}  else {
		pw := 0
		if idx >= 1 && ptn & (1<<(idx-1)) == 0 {
			pw++
		}
		if idx <= 12 && ptn & (1<<(idx+1)) == 0 {
			pw++
		}
		dfs(ti, rate * float64(P[pw]) / 100, ptn, tptn | (1<<idx), idx+1)
		dfs(ti, rate * float64(100 - P[pw]) / 100, ptn, tptn, idx+1)
	}
}
