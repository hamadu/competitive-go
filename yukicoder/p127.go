package main
import "fmt"

var a []int
var memo [][][]int
var INF = 114514

func main()  {
	var n int
	fmt.Scan(&n)

	a = make([]int, n)
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&a[i])
	}
	memo = make([][][]int, 2)
	for i := 0 ; i <= 1 ; i++ {
		memo[i] = make([][]int, n+1)
		for j := 0 ; j <= n ; j++ {
			memo[i][j] = make([]int, n+1)
			for k := 0 ; k <= n ; k++ {
				memo[i][j][k] = -1
			}
		}
	}

	ans := 0
	for i := 0 ; i <= 1 ; i++ {
		for j := 0 ; j < n ; j++ {
			for k := j ; k < n ; k++ {
				res := dfs(i, j, k)
				ans = max(ans, res)
			}
		}
	}
	fmt.Println(ans)
}

func dfs(flg, left, right int) int {
	if (left > right) {
		return -INF
	} else if (left == right) {
		return 1
	} else if (memo[flg][left][right] != -1) {
		return memo[flg][left][right]
	}
	res := 0
	if flg == 0 {
		if a[left] < a[right] {
			res = max(res, dfs(1, left, right)+1)
		}
		res = max(res, dfs(0, left, right-1))
	} else {
		if a[left] > a[right] {
			res = max(res, dfs(0, left, right)+1)
		}
		res = max(res, dfs(1, left+1, right))
	}
	memo[flg][left][right] = res
	return res
}


func max(a, b int) int {
	if a < b  {
		return b
	} else {
		return a
	}
}
