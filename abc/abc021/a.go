package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)

	fmt.Println(n)
	for i := 0 ; i < n ; i++ {
		fmt.Println(1)
	}
}
