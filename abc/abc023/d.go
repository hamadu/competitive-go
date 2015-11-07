package main
import (
	"fmt"
	"io"
	"bufio"
	"os"
	"strconv"
	"math"
	"sort"
)

func main() {
	n := nextInt()

	bl := make([][]int, n)
	time := make([]int, n)
	for i := range bl {
		bl[i] = make([]int, 2)
		bl[i][0] = nextInt()
		bl[i][1] = nextInt()
	}

	min := 0
	max := int(math.Pow10(16))
	for i := range bl {
		if min < bl[i][0] {
			min = bl[i][0]
		}
	}
	min -= 1
	for max - min > 1 {
		med := (max + min) / 2
		if isOK(med, bl, time) {
			max = med
		} else {
			min = med
		}
	}
	fmt.Println(max)
}

func isOK(max int, bal [][]int, time []int) bool {
	n := len(bal)
	for i := 0 ; i < n ; i++ {
		time[i] = (max - bal[i][0]) / bal[i][1]
	}
	sort.Ints(time)
	for i := 0 ; i < n ; i++ {
		if time[i] < i {
			return false
		}
	}
	return true
}

// ------------------


var rdr = bufio.NewReader(os.Stdin)

func nextInt() int {
	i, e := strconv.Atoi(readToken(20))
	if e != nil {
		panic(e)
	}
	return i
}

func nextString(limit int) string {
	return readToken(limit)
}

func readToken(limit int) string {
	buf := make([]byte, 0, limit)

	for {
		byte, err := rdr.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if byte != 10 && byte != 13 && byte != 32 {
			buf = append(buf, byte)
			break
		}
	}
	for {
		byte, err := rdr.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if byte == 10 || byte == 13 || byte == 32 {
			break
		}
		buf = append(buf, byte)
	}
	return string(buf)
}
