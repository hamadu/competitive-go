package main
import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	sum := 0
	if n % 2 == 0 {
		n /= 2
	}
	for i := 1 ; i * i <= n ; i++ {
		if n % i == 0 {
			if i * i != n {
				sum += n / i
			}
			sum += i
		}
	}
	fmt.Println(sum)
}
