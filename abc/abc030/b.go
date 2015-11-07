package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m float64
	fmt.Scan(&n, &m)
	hourDeg := 360.0 * math.Mod(n, 12) / 12.0 + 30.0 * m / 60.0
	minuteDeg := 360.0 * m / 60.0

	diff := math.Abs(hourDeg - minuteDeg)
	if diff > 180 {
		diff = 360 - diff
	}
	fmt.Println(diff)
}