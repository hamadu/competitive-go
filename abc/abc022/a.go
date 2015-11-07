package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"io"
)

func main() {
	n := nextInt()
	a := make([]int, n)
	for i := 0 ; i < n ; i++ {
		a[i] = nextInt()
	}

}

func cont(w, s, t int) int {
	if s <= w && w <= t {
		return 1
	} else {
		return 0
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

