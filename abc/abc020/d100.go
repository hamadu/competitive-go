package main

import "fmt"

const MOD = 1000000007

func main() {
	var n, k = 0, 0
	fmt.Scan(&n, &k)

	if k > 100 {
		panic("i dn know")
	}

	ans := 0
	for mod := 1 ; mod <= k ; mod++ {
		from := mod
		if from > n {
			break
		}
		to := n - n % k + mod
		if to > n {
			to -= k
		}
		div := gcd(mod, k)
		sum := (from + to) * ((to - from) / k + 1) / 2
		ans += (sum * k / div) % MOD
		ans %= MOD
	}
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}