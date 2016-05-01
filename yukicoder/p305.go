package main
import "fmt"

func main()  {
	ans := ""
	for i := 0 ; i <= 9 ; i++ {
		best := -1
		bestD := -1
		for d := 0 ; d <= 9 ; d++ {
			query := ""
			for j := 0 ; j <= 9 ; j++ {
				if i == j {
					query += fmt.Sprintf("%d", d)
				} else {
					query += "0"
				}
			}
			fmt.Println(query)

			num := 0
			res := ""
			fmt.Scan(&num, &res)
			if num == 10 {
				return
			}
			if best < num {
				best = num
				bestD = d
			}
		}
		ans += fmt.Sprintf("%d", bestD)
	}

	num := 0
	res := ""
	fmt.Println(ans)
	fmt.Scan(&num, &res)
}
