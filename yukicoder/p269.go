package main
import (
	"fmt"
)

const MOD = 1000000007

func split(n, m int) [][]int {
	dp := make([][]int, m+1)
	for i := 0 ; i <= m ; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1 ; i <= m ; i++ {
		for j := 0 ; j <= n ; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= i {
				dp[i][j] += dp[i][j-i]
			}
			dp[i][j] %= MOD
		}
	}
	return dp
}

func main() {
	n, s, k := 0, 0, 0
	fmt.Scan(&n, &s, &k)

	over := n*(n-1)*k/2
	left := s - over
	if left >= 0 {
		dp := split(left, n)
		fmt.Println(dp[n][left])
	} else {
		fmt.Println(0)
	}
}

