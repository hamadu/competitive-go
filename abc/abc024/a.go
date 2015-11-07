package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"io"
)

func main() {
	a := nextInt()
	b := nextInt()
	c := nextInt()
	k := nextInt()
	s := nextInt()
	t := nextInt()

	sum := a * s + b * t
	if s + t >= k {
		sum -= (s + t) * c
	}
	fmt.Println(sum)
}

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
