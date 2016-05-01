package main
import "fmt"

func main()  {
	for i := 0 ; i < 1000 ; i++ {
		fmt.Printf("%03d\n", i)

		s := ""
		fmt.Scan(&s)
		if s == "unlocked" {
			break
		}
	}
}
