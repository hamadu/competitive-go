package main
import (
	"fmt"
)

func main() {
	var fr, to string
	var num int
	fmt.Scan(&fr, &num, &to)

	if isPossible(fr, to, num) {
		fmt.Println("FAILURE")
	} else {
		fmt.Println("SUCCESS")
	}
}

func isPossible(fr, to string, num int) bool {
	if count(fr) != count(to) {
		return false
	}
	if num == 0 {
		return fr == to
	} else if num == 1 {
		if string(fr[1]) + string(fr[0]) + string(fr[2]) == to {
			return true
		}
		if string(fr[0]) + string(fr[2]) + string(fr[1]) == to {
			return true
		}
	} else {
		return true
	}
	return false
}

func count(x string) int {
	ctr := 0
	for i := 0 ; i < 3 ; i++ {
		if x[i] == 'o' {
			ctr++
		}
	}
	return ctr
}
