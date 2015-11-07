package main
import "fmt"

func main() {
	n := 0
	s := ""
	fmt.Scan(&n, &s)

	ac := "b"
	for k := 0 ; k < 100 ; k++ {
		if s == ac {
			fmt.Println(k)
			return
		}
		if k % 3 == 0 {
			ac = "a" + ac + "c"
		} else if k % 3 == 1 {
			ac = "c" + ac + "a"
		} else {
			ac = "b" + ac + "b"
		}
	}
	fmt.Println(-1)
}
