package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"io"
)

var depth []int
var rdepth []int
var graph [][]int
const INF = 114514

func dfs(now, par, dep int) int {
	depth[now] = dep
	min := INF
	for i := 0 ; i < len(graph[now]) ; i++ {
		to := graph[now][i]
		if to == par {
			continue
		}
		cn := dfs(to, now, dep+1) + 1
		if min > cn {
			min = cn
		}
	}
	if min == INF {
		min = 0
	}
	rdepth[now] = min
	return min
}

func dfs2(now, par int) {
	if par >= 0 && rdepth[now] > rdepth[par] + 1 {
		rdepth[now] = rdepth[par] + 1
	}
	for i := 0 ; i < len(graph[now]) ; i++ {
		to := graph[now][i]
		if to == par {
			continue
		}
		dfs2(to, now)
	}
}

func main() {
	n := nextInt()
	edges := make([][]int, n-1)
	deg := make([]int, n)
	for i := 0 ; i < n-1 ; i++ {
		edges[i] = make([]int, 2)
		a := nextInt()-1
		b := nextInt()-1
		deg[a]++
		deg[b]++
		edges[i][0] = a
		edges[i][1] = b
	}
	graph = make([][]int, n)
	for i := 0 ; i < n ; i++ {
		graph[i] = make([]int, deg[i])
	}
	for i := 0 ; i < n-1 ; i++ {
		a := edges[i][0]
		b := edges[i][1]
		deg[a]--
		deg[b]--
		graph[a][deg[a]] = b
		graph[b][deg[b]] = a
	}

	depth = make([]int, n)
	rdepth = make([]int, n)
	dfs(0, -1, 0)
	dfs2(0, -1)

	for i := 0 ; i < n ; i++ {
		if depth[i] < rdepth[i] {
			fmt.Println(depth[i])
		} else {
			fmt.Println(rdepth[i])
		}
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