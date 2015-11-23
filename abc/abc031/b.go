package main

import "fmt"

func main()  {
	n, l, h := 0, 0, 0
	fmt.Scan(&l, &h, &n)


	for i := 0; i < n; i++ {
		a := 0
		fmt.Scan(&a)
		if a < l {
			fmt.Println(l-a)
		} else if l <= a && a <= h {
			fmt.Println(0)
		} else {
			fmt.Println(-1)
		}
	}
}
