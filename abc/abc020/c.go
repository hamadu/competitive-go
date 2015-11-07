package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"io"
	"container/list"
)

var DX = []int {-1, 0, 1, 0}
var DY = []int {0, 1, 0, -1}
const INF = 10000000000

func main() {
	h := nextInt()
	w := nextInt()
	t := nextInt()
	table := make([]string, h)
	for i := 0; i < h; i++ {
		table[i] = nextString(w+10)
	}

	sx, sy := 0, 0
	gx, gy := 0, 0
	for i := 0 ; i < h ; i++ {
		for j := 0 ; j < w ; j++ {
			if table[i][j] == 'S' {
				sy = i
				sx = j
			} else if table[i][j] == 'G' {
				gy = i
				gx = j
			}
		}
	}

	moveCost := build(sy, sx, gy, gx, table)

	min := 0
	max := INF
	for max - min > 1 {
		med := (max + min) / 2
		best := INF
		for i := 1 ; i < len(moveCost) ; i++ {
			cost := moveCost[i] + i * med
			if best > cost {
				best = cost
			}
		}
		if best <= t {
			min = med
		} else {
			max = med
		}
	}
	fmt.Println(min+1)
}

func build(sy, sx, gy, gx int, table []string) []int {
	n := len(table)
	m := len(table[0])
	cost := make([][][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([][]int, m)
		for j := 0 ; j < m ; j++ {
			cost[i][j] = make([]int, n*m)
			for k := 0 ; k < n*m ; k++ {
				cost[i][j][k] = INF
			}
		}
	}


	q := list.New()
	q.PushBack(sy)
	q.PushBack(sx)
	q.PushBack(0)
	q.PushBack(0)
	for q.Len() >= 1 {
		y := q.Remove(q.Front()).(int)
		x := q.Remove(q.Front()).(int)
		z := q.Remove(q.Front()).(int)
		t := q.Remove(q.Front()).(int)
		for d := 0 ; d < 4 ; d++ {
			ty := y + DY[d]
			tx := x + DX[d]
			if ty < 0 || tx < 0 || ty >= n || tx >= m || z+1 >= n*m {
				continue
			}
			tz := z
			tt := t+1
			if table[ty][tx] == '#' {
				tz++
			}
			if cost[ty][tx][tz] > tt {
				cost[ty][tx][tz] = tt
				q.PushBack(ty)
				q.PushBack(tx)
				q.PushBack(tz)
				q.PushBack(tt)
			}
		}
	}
	return cost[gy][gx]
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