package main
import (
	"strconv"
	"os"
	"bufio"
	"io"
	"sort"
	"math"
	"fmt"
)

type Move struct {
	idx int
	abs int
	dir int
}

type Moves []Move

func (m Moves) Len() int {
	return len(m)
}

func (m Moves) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (p Moves) Less(i, j int) bool {
	return p[i].abs > p[j].abs
}

func main() {
	s := nextString(100010)

	n := len(s)

	diff := make([]int, n)

	plus := 0
	minus := 0
	di := 0
	for i := n-1 ; i >= 0 ; i-- {
		if s[i] == 'M' {
			diff[di] = plus - minus
			di++
		} else if s[i] == '+' {
			plus++
		} else {
			minus++
		}
	}

	raw := resize(diff, di)

	var moves Moves = make([]Move, di)
	for i := 0 ; i < di ; i++ {
		abs := int(math.Abs(float64(raw[i])))
		dir := 1
		if abs >= 1 {
			dir = raw[i] / abs
		}
		moves[i] = Move{i, abs, dir}
	}

	sort.Sort(moves)

	pl := 0
	mi := 0
	ans := 0
	for i := 0 ; i < di ; i++ {
		ans += moves[i].abs
		if moves[i].dir == 1 {
			pl++
		} else {
			mi++
		}
	}
	for i := di-1 ; i >= 0 ; i-- {
		if pl == mi {
			break
		}
		if pl > mi {
			if moves[i].dir == 1 {
				pl--
				mi++
				ans -= moves[i].abs * 2
			}
		} else {
			if moves[i].dir == -1 {
				pl++
				mi--
				ans -= moves[i].abs * 2
			}
		}
	}
	fmt.Println(ans)
}

func resize(a []int, size int) []int {
	b := make([]int, size)
	for i := range b {
		b[i] = a[i]
	}
	return b
}

// ===== io library


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
