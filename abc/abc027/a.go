package main
import "fmt"

func main() {
	a, b, c := 0, 0, 0
	fmt.Scan(&a, &b, &c)

	if a == b {
		fmt.Println(c)
	} else if a == c {
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}
}
