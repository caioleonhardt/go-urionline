package urionline

/*
 * Primo Rápido
 */
import (
	"fmt"
	"math"
)

func isPrimo(primo int) bool {
	if primo == 2 {
		return true
	}

	maxP := int(math.Sqrt(float64(primo)))
	for j := 3; j <= maxP; j += 2 {
		if (primo/j)*j == primo {
			return false
		}
	}

	return true
}
func main() {
	var n int

	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		var primo int
		fmt.Scanf("%d\n", &primo)

		if isPrimo(primo) {
			fmt.Printf("Prime\n")
		} else {
			fmt.Printf("Not Prime\n")
		}
	}
}
