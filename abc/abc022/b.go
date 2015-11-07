package main

import (
	"bufio"
	"os"
	"strconv"
	"io"
	"fmt"
)

func main() {
	n := nextInt()
	set := make(map[int]struct{})
	count := 0
	for i := 0 ; i < n ; i++ {
		kind := nextInt()
		_, have := set[kind]
		if have {
			count++
		} else {
			set[kind] = struct{}{}
		}
	}
	fmt.Println(count)
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
