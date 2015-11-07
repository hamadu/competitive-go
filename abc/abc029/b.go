package main

import (
	"fmt"
	"regexp"
)

func main() {
	counter := 0
	for i := 0 ; i < 12 ; i++ {
		s := ""
		fmt.Scan(&s)
		ok, _ := regexp.Match("r", []byte(s))
		if ok {
			counter++
		}
	}
	fmt.Println(counter)
}