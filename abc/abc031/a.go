package main

import "fmt"

func main()  {
	a, d := 0, 0
	fmt.Scan(&a, &d)

	ad := (a+1)*d
	da := a*(d+1)
	if ad >= da {
		fmt.Println(ad)
	} else {
		fmt.Println(da)
	}
}
