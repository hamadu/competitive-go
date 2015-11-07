package main

import (
	"fmt"
	"sort"
)

const MOD = 1000000007

func main() {
	var n, k = 0, 0
	fmt.Scan(&n, &k)

	divisors := make_divisors(k)
	table := make([]int, len(divisors))

	ans := 0
	for i := len(divisors)-1 ; i >= 0 ; i-- {
		d := divisors[i]

		// compute sum of 1 <= xd <= n
		to := n - n % d
		table[i] += to * (to / d + 1) / 2

		for j := i-1 ; j >= 0 ; j-- {
			if divisors[i] % divisors[j] == 0 {
				table[j] -= table[i]
			}
		}
		ans += (table[i] / d) % MOD
		ans %= MOD
	}
	ans *= k
	ans %= MOD
	fmt.Println(ans)
}

func make_divisors(n int) []int {
	divisors := make([]int, 10000)
	di := 0
	for i := range divisors {
		divisors[i] = 10000000000
	}
	for i := 1 ; i * i <= n ; i++ {
		if n % i == 0 {
			divisors[di] = i
			di++
			if i * i != n {
				divisors[di] = n / i
				di++
			}
		}
	}
	sort.Ints(divisors)

	return_divisors := make([]int, di)
	for i := range return_divisors {
		return_divisors[i] = divisors[i]
	}
	return return_divisors
}
