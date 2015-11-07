package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	i, e := strconv.Atoi(nextString())
	if e != nil {
		panic(e)
	}
	return i
}

func nextString() string {
	sc.Split(bufio.ScanWords)
	sc.Scan()
	return sc.Text()
}

func main() {
	n := nextInt()
	m := nextInt()
	x := nextInt()
	y := nextInt()

	a := make([]int, n)
	b := make([]int, m)
	for i := 0 ; i < n ; i++ {
		a[i] = nextInt()
	}
	for i := 0 ; i < m ; i++ {
		b[i] = nextInt()
	}

	time := 0
	aidx, bidx := 0, 0
	counter := 0
	for {
		if counter % 2 == 0 {
			for aidx < n {
				if time <= a[aidx] {
					counter++
					time = a[aidx] + x
					break;
				}
				aidx++
			}
			if counter % 2 == 0 {
				break
			}
		} else {
			for bidx < m {
				if time <= b[bidx] {
					counter++
					time = b[bidx] + y
					break;
				}
				bidx++
			}
			if counter % 2 == 1 {
				break
			}
		}
	}
	fmt.Println(counter / 2)
}
