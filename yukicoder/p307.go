package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"io"
)

var dy = []int{0, -1, 0, 1}
var dx = []int{1, 0, -1, 0}

func main()  {
	h := nextInt()
	w := nextInt()

	a := make([][]int, h)
	for i := 0 ; i < h ; i++ {
		a[i] = make([]int, w)
		for j := 0 ; j < w ; j++ {
			a[i][j] = nextInt()
		}
	}

	que := make([]int, 114514)

	all := false
	q := nextInt()
	final := -1
	for i := 0 ; i < q ; i++ {
		r := nextInt()-1
		c := nextInt()-1
		x := nextInt()
		if all {
			final = x
			continue
		}
		if a[r][c] == x {
			continue
		}

		cnt := 1
		a[r][c] = x
		qh := 0
		qt := 0
		que[qh] = r
		qh++
		que[qh] = c
		qh++
		for qt < qh {
			ny := que[qt]
			qt++
			nx := que[qt]
			qt++
			for d := 0 ; d < 4 ; d++ {
				ty := ny + dy[d]
				tx := nx + dx[d]
				if ty < 0 || ty >= h || tx < 0 || tx >= w {
					continue
				}
				if a[ty][tx] == x {
					continue
				}
				a[ty][tx] = x
				que[qh] = ty
				qh++
				que[qh] = tx
				qh++
				cnt++
			}
		}
		if cnt >= h*w {
			all = true
			final = x
		}
		cnt++
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if j >= 1 {
				fmt.Printf(" ")
			}
			out := a[i][j]
			if final != -1 {
				out = final
			}
			fmt.Printf("%d", out)
		}
		fmt.Println()
	}

}


// ====

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
