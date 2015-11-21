package main
import "fmt"

const MOD = 1000000007

func main() {
	n, a, b := 0, 0, 0
	fmt.Scan(&n, &a, &b)
	a -= 1
	b -= 1

	graph := make([][]int, n)
	for i := 0 ; i < n ; i++ {
		graph[i] = make([]int, n)
	}
	m := 0
	fmt.Scan(&m)
	for i := 0 ; i < m ; i++ {
		a := 0
		b := 0
		fmt.Scan(&a, &b)
		graph[a-1][b-1] = 1
		graph[b-1][a-1] = 1
	}

	dp := make([][]int, n)
	flg := make([][]bool, n)
	for i := 0 ; i < n ; i++ {
		dp[i] = make([]int, n)
		flg[i] = make([]bool, n)
	}
	dp[0][a] = 1
	flg[0][a] = true

	for i := 1 ; i < n ; i++ {
		for j := 0 ; j < n ; j++ {
			for k := 0 ; k < n ; k++ {
				if graph[j][k] == 1 && flg[i-1][k] {
					dp[i][j] += dp[i-1][k]
					flg[i][j] = true
				}
			}
			dp[i][j] %= MOD
		}
		if flg[i][b] {
			fmt.Println(dp[i][b])
			break
		}
	}
}
