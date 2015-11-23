package main

import (
	"fmt"
)

const INF = 114514

func main()  {
	k, n := 0, 0
	fmt.Scan(&k, &n)

	pair := make([][]string, n)
	for i := 0 ; i < n ; i++ {
		pair[i] = make([]string, 2)
		for j := 0 ; j < 2 ; j++ {
			fmt.Scan(&pair[i][j])
		}
	}

	dfs(0, make([]int, 9), pair)

	for i := 0 ; i < k ; i++ {
		fmt.Println(answer[i])
	}
}

var answer []string


func isOK(length []int, pair [][]string) bool {
	n := len(pair)
	answer = make([]string, 9)
	for i := 0 ; i < n ; i++ {
		nl := len(pair[i][0])
		sl := len(pair[i][1])

		sum := 0
		for j := 0 ; j < nl ; j++ {
			k := pair[i][0][j]-'1'
			if sum+length[k] > sl {
				return false
			}
			if len(answer[k]) == 0 {
				answer[k] = pair[i][1][sum:(sum+length[k])]
			} else if answer[k] != pair[i][1][sum:(sum+length[k])] {
				return false
			}
			sum += length[k]
		}
		if sum != sl {
			return false
		}
	}
	return true
}

func dfs(idx int, length []int, pair [][]string) bool {
	if idx == 9 {
		if (isOK(length, pair)) {
			return true
		} else {
			return false
		}
	}

	for l := 1 ; l <= 3 ; l++ {
		length[idx] = l
		if dfs(idx+1, length, pair) {
			return true
		}
	}
	return false
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