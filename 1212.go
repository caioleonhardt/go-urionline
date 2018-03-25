package urionline

/*
 * Aritmética Primária
 */
import (
	"fmt"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	for {
		var a, b int

		fmt.Scanf("%d %d\n", &a, &b)

		if a == 0 && b == 0 {
			break
		}

		strA := reverse(strconv.Itoa(a))
		strB := reverse(strconv.Itoa(b))

		if len(strA) < len(strB) {
			strA, strB = strB, strA
		}

		max := len(strA)

		occurs := 0

		current := 0
		for i := 0; i < max; i++ {
			numa, _ := strconv.Atoi(string(strA[i]))

			numb := 0

			if i < len(strB) {
				numb, _ = strconv.Atoi(string(strB[i]))
			}

			res := numa + numb + current

			//reset current
			current = 0

			if res >= 10 {
				occurs++
				current = int(res / 10)
			}
		}

		if occurs == 0 {
			fmt.Printf("No carry operation.\n")
		} else if occurs == 1 {
			fmt.Printf("1 carry operation.\n")
		} else {
			fmt.Printf("%d carry operations.\n", occurs)
		}
	}
}
