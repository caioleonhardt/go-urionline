package urionline

import (
	"fmt"
)

func main() {
	for {
		var a int
		fmt.Scanf("%d\n", &a)

		if a == 0 {
			break
		}
		for i := 1; i <= a; i++ {
			if i == 1 {
				fmt.Printf("1")
			} else {
				fmt.Printf(" %d", i)
			}
		}
		fmt.Printf("\n")
	}
}
