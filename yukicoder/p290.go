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
	s := nextString(n+10)
	L := -1
	for i := 0 ; i < n-1 ; i++ {
		if s[i] == s[i+1] {
			L = i
			break
		}
	}
	if L == -1 {
		if n <= 3 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	} else {
		fmt.Println("YES")
	}
}

//====

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
