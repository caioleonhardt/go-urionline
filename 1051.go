package urionline

import (
	"fmt"
)

func main() {
	var sal float32
	var t map[int]float32

	fmt.Scanf("%f", &sal)

	t = make(map[int]float32)

	if sal > 0.0 && sal < 2000.00 {
		fmt.Printf("Isento\n")
	} else {
		if sal > 4500.00 {
			t[28] = sal - 4500
			sal -= t[28]
		}

		if sal > 3000.00 && sal <= 4500.00 {
			t[18] = sal - 3000.00
			sal -= t[18]
		}

		if sal > 2000.00 && sal <= 3000.00 {
			t[8] = sal - 2000.00
			sal -= t[8]
		}

		res := t[8] * 0.08
		res += t[18] * 0.18
		res += t[28] * 0.28

		fmt.Printf("R$ %.2f\n", res)
	}

}
