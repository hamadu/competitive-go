package main
import "fmt"

func main() {
	n, a, b := 0, 0, 0
	fmt.Scan(&n, &a, &b)

	set := map[int]int{}
	set[a] = 1
	set[b] = 1

	k := 0
	fmt.Scan(&k)

	isOK := true
	for i := 0 ; i < k ; i++ {
		a := 0
		fmt.Scan(&a)
		if set[a] == 1 {
			isOK = false
		}
		set[a] = 1
	}
	if isOK {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
