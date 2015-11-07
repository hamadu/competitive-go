package main

import (
	"bufio"
	"os"
	"strconv"
	"io"
	"fmt"
)

const INF = 500000000

func main() {
	n := nextInt()
	m := nextInt()
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = INF
		}
	}

	for i := 0; i < m; i++ {
		u := nextInt()-1
		v := nextInt()-1
		graph[u][v] = nextInt()
		graph[v][u] = graph[u][v]
	}

	best := INF
	for i := 0 ; i < n ; i++ {
		if (graph[0][i] < INF) {
			temp := graph[0][i]
			graph[0][i] = INF
			graph[i][0] = INF
			cost := findShortestPath(i, 0, graph) + temp
			if best > cost {
				best = cost
			}
			graph[0][i] = temp
		}
	}
	if best >= INF {
		best = -1
	}
	fmt.Println(best)
}




func findShortestPath(from, to int, graph [][]int) int {
	n := len(graph)
	distance := make([]int, n)
	for i := 0 ; i < n ; i++ {
		distance[i] = INF
	}
	distance[from] = 0

	for {
		upd := false
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if distance[j] > distance[i] + graph[i][j] {
					upd = true
					distance[j] = distance[i] + graph[i][j]
				}
			}
		}
		if !upd {
			break
		}
	}
	return distance[to]
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
