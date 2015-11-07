package main
import (
	"io"
	"bufio"
	"os"
	"strconv"
	"fmt"
)

const MOD = 1000000007

func main() {
	a := nextInt()
	b := nextInt()
	c := nextInt()

	upper_r := b * (c - a + MOD) % MOD
	downr_r := ((a * b) % MOD + (a * c) % MOD - (b * c) % MOD + MOD) % MOD
	r := upper_r * inv(downr_r, MOD) % MOD

	upper_n := (2 * b * c % MOD - a * c % MOD - a * b % MOD + 2 * MOD) % MOD
	downr_n := ((a * b) % MOD + (a * c) % MOD - (b * c) % MOD + MOD) % MOD
	n := upper_n * inv(downr_n, MOD) % MOD

	fmt.Println(n - r, r)
}

// math -----------------------------------

func pow(a, p, mod int) int {
	res := 1
	aa := a
	for p >= 1 {
		if p & 1 == 1 {
			res = (res * aa) % mod
		}
		aa = (aa * aa) % mod
		p /= 2
	}
	return res
}

func inv(a, mod int) int {
	return pow(a, mod-2, mod)
}


// io 00-----------------------------------

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

