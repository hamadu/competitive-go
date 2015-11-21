package main
import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	fr := n / 2 + 1
	state := 1
	for fr >= 2 {
		if state == 1 {
			fr = fr / 2
		} else {
			fr = (fr + 1) / 2
		}
		state = 1 - state
	}
	if state == 0 {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}