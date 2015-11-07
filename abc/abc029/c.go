package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)
	rec(n, "")
}

func rec(len int, pass string) {
	if len == 0 {
		fmt.Println(pass)
		return
	}
	rec(len-1, pass + "a")
	rec(len-1, pass + "b")
	rec(len-1, pass + "c")
}