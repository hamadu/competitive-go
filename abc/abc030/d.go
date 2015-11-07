package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"math/big"
	"io"
)

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

const INF = 1000000000

func main() {
	n := nextInt()
	a := nextInt() - 1

	k := nextString(100010)
	b := make([]int, n)
	for i := 0 ; i < n  ; i++ {
		b[i] = nextInt() - 1
	}
	M := big.NewInt(0)
	for i := 0 ; i < len(k) ; i++ {
		M.Mul(M, big.NewInt(10))
		M.Add(M, big.NewInt(int64(k[i] - '0')))
	}
	M.Add(M, big.NewInt(-1))

	var time int64 = 0
	var cycle int64 = 0
	firstVisit := make([]int64, n)
	part := make([]int, n+10)

	for i := 0 ; i < n ; i++ {
		firstVisit[i] = INF
	}


	now := b[a]
	for {
		if firstVisit[now] != INF {
			cycle = time - firstVisit[now]
			break
		}
		part[time] = now
		firstVisit[now] = time
		now = b[now]
		time++
	}

	answer := 0
	if (M.Cmp(big.NewInt(firstVisit[now])) < 0) {
		answer = part[M.Int64()]
	} else {
		answer = part[firstVisit[now] + M.Mod(M.Sub(M, big.NewInt(firstVisit[now])), big.NewInt(cycle)).Int64()]
	}
	fmt.Println(answer+1)
}
