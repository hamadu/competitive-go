package main
import (
	"io"
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	nextInt()
	d := nextInt()
	k := nextInt()
	moves := make([][]int, d)
	for i := 0 ; i < d ; i++ {
		moves[i] = make([]int, 2)
		moves[i][0] = nextInt()
		moves[i][1] = nextInt()
	}

	for i := 0 ; i < k ; i++ {
		s := nextInt()
		t := nextInt()

		minS := s
		maxS := s
		ans := d
		for i := 0 ; i < d ; i++ {
			if minS <= t && t <= maxS {
				ans = i
				break
			}
			if moves[i][0] <= minS && minS <= moves[i][1] {
				minS = moves[i][0]
			}
			if moves[i][0] <= maxS && maxS <= moves[i][1] {
				maxS = moves[i][1]
			}
		}
		fmt.Println(ans)
	}
}

// -----------------------------------

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

