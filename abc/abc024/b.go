package main
import (
	"io"
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	n := nextInt()
	t := nextInt()
	sum := 0
	shut := 0
	for i := 0 ; i < n ; i++ {
		time := nextInt()
		if time < shut {
			sum += time + t - shut
		} else {
			sum += t
		}
		shut = time + t
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

