package main
import (
	"fmt"
)

var DX = []int{0, -1, 0, 1}
var DY = []int{1, 0, -1, 0}

func main() {
	correct := make([][]int, 4)
	table := make([][]int, 4)
	for i := 0 ; i < 4 ; i++ {
		correct[i] = make([]int, 4)
		table[i] = make([]int, 4)
	}
	for i := 0 ; i < 4 ; i++ {
		for j := 0 ; j < 4 ; j++ {
			fmt.Scan(&table[i][j])
			correct[i][j] = (i*4+(j+1))%16
		}
	}

	if check(table, correct) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func check(table, correct [][]int) bool {
	same := true
	for i := 0 ; i < 4 ; i++ {
		for j := 0; j < 4; j++ {
			if table[i][j] != correct[i][j] {
				same = false
			}
		}
	}
	if same {
		return true
	}

	for i := 0 ; i < 4 ; i++ {
		for j := 0; j < 4; j++ {
			if table[i][j] == 0 {
				want := correct[i][j]
				for d := 0 ; d < 4 ; d++ {
					ti := i + DY[d]
					tj := j + DX[d]
					if ti < 0 || ti >= 4 || tj < 0 || tj >= 4 {
						continue
					}
					if table[ti][tj] == want {
						table[i][j] = want
						table[ti][tj] = 0
						return check(table, correct)
					}
				}
			}
		}
	}
	return false
}