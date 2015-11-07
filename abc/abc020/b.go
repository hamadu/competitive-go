package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)
	ans, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	fmt.Println(ans * 2)
}
