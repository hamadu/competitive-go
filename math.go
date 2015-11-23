package competitive_go

import "sort"

func abs(a int) int {
	if a > 0  {
		return a
	} else {
		return -a
	}
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

func makeDivisors(n int) []int {
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

func makePrimes(limit int) []int {
	primes := make([]int, limit)
	isp := make([]bool, limit)
	for i := range isp {
		isp[i] = true
	}
	isp[0] = false
	isp[1] = false
	for i := 2 ; i < len(isp) ; i++ {
		if isp[i] {
			for ii := i * i ; ii < len(isp); ii += i {
				isp[ii] = false
			}
		}
	}
	pi := 0
	for i := range isp {
		if isp[i] {
			primes[pi] = i
			pi++
		}
	}
	return_primes := make([]int, pi)
	for i := 0 ; i < pi ; i++ {
		return_primes[i] = primes[i]
	}
	return return_primes
}
