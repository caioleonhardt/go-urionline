package urionline

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d %d\n", &b, &c)

	if a < b+c {
		fmt.Printf("Deixa para amanha!\n")
	} else {
		fmt.Printf("Farei hoje!\n")
	}

}
