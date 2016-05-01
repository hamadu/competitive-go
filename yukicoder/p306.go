package main
import (
	"fmt"
)

func main()  {
	ax, ay, bx, by := 0.0, 0.0, 0.0, 0.0
	fmt.Scan(&ax, &ay, &bx, &by)

	if ay > by {
		fmt.Printf("%.9f\n", solve(bx, by, ax, ay))
	} else {
		fmt.Printf("%.9f\n", solve(ax, ay, bx, by))
	}
}

func solve(ax, ay, bx, by float64) float64 {
	if ay == by {
		return ay
	}
	return ay + (by - ay) * ax / (ax + bx)
}