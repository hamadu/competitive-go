package main

import (
	"bufio"
	"os"
	"strconv"
	"io"
	"fmt"
	"math/rand"
	"math"
	"sort"
)

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) Rotate(rad float64) {
	v.X, v.Y =
		v.X * math.Cos(rad) + v.Y * math.Sin(rad),
		- v.X * math.Sin(rad) + v.Y * math.Cos(rad)
}

type Polygon []Vertex

func (p Polygon) Len() int {
	return len(p)
}

func (p Polygon) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Polygon) Less(i, j int) bool {
	return p[i].X < p[j].X || (p[i].X == p[j].X && p[i].Y < p[j].Y)
}

func (p Polygon) Rotate(rad float64) {
	n := p.Len()
	for i := 0 ; i < n ; i++ {
		p[i].Rotate(rad)
	}
}

func (p Polygon) ShortestDistance() float64 {
	dist2 := 9e18
	n := p.Len()
	sort.Sort(p)

	for i := 0; i < n ; i++ {
		for j := i+1 ; j < n ; j++ {
			dx := math.Abs(p[i].X - p[j].X)
			if dx * dx > dist2 {
				break
			}
			dy := math.Abs(p[i].Y - p[j].Y)
			d := dx * dx + dy * dy
			if dist2 > d {
				dist2 = d
			}
		}
	}
	return math.Sqrt(dist2)
}

func main() {
	n := nextInt()
	va := make(Polygon, n)
	for i := 0; i < n; i++ {
		va[i] = Vertex{ float64(nextInt()), float64(nextInt()) }
	}
	vb := make(Polygon, n)
	for i := 0; i < n; i++ {
		vb[i] = Vertex{ float64(nextInt()), float64(nextInt()) }
	}
	va.Rotate(rand.Float64() * math.Pi * 2)
	vb.Rotate(rand.Float64() * math.Pi * 2)

	fmt.Println(vb.ShortestDistance() / va.ShortestDistance())
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
