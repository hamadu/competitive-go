package main
import (
	"fmt"
	"io"
	"bufio"
	"os"
	"strconv"
)

func main() {
	r := nextInt()
	c := nextInt()
	k := nextInt()
	n := nextInt()

	rn := make([]int, r)
	cn := make([]int, c)

	ame := make([][]int, n)


	for i := 0 ; i < n ; i++ {
		ri := nextInt()-1
		ci := nextInt()-1
		ame[i] = []int{ri, ci}
		rn[ri]++
		cn[ci]++
	}

	cdeg := make([]int, n+1)
	for i := 0 ; i < c ; i++ {
		cdeg[cn[i]]++
	}

	ans := 0
	for i := 0 ; i < r ; i++ {
		need := k - rn[i]
		if need >= 0 {
			ans += cdeg[need]
		}
	}

	for i := 0 ; i < n ; i++ {
		ri := ame[i][0]
		ci := ame[i][1]
		if rn[ri] + cn[ci] == k {
			ans--
		} else if rn[ri] + cn[ci] == k+1 {
			ans++
		}
	}

	fmt.Println(ans)
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
