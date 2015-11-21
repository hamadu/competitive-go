package main
import (
	"fmt"
)

const MOD = 1000000007

func pow(a, p, mod int) int {
	ret := 1
	aa := a
	for p >= 1 {
		if p & 1 == 1 {
			ret *= aa
			ret %= mod
		}
		aa = aa * aa % mod
		p >>= 1
	}
	return ret
}

func inv(a, mod int) int {
	return pow(a, mod-2, mod)
}

func comb(n, k, mod int) int {
	fact := make([]int, n+10)
	invfact := make([]int, n+10)

	fact[0] = 1
	invfact[0] = 1
	for i := 1 ; i <= n ; i++ {
		fact[i] = fact[i-1] * i % mod
		invfact[i] = invfact[i-1] * inv(i, mod) % mod
	}
	return fact[n] * invfact[k] % mod * invfact[n-k] % mod
}

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	fmt.Println(comb(n+k-1, k, MOD))
}