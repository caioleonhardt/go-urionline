package urionline

import (
	"fmt"
	"io"
)

func desc(a, b, c int) (int, int, int) {
	if b > a {
		a, b = b, a
	}

	if c > b {
		b, c = c, b
	}

	if b > a {
		a, b = b, a
	}

	return a, b, c
}

func triplaPitagoricaPrimitiva(a, b, c int) (bool, bool) {
	tp := a*a == b*b+c*c

	return tp, isPrimitivo(a, b, c)
}

func isPrimitivo(a, b, c int) bool {
	ra, rb, rc := a, b, c
	mdc := make([]int, 0, 5)

	div := 2
	for div <= ra || div <= rb || div <= rc {
		if ra%div == 0 && rb%div == 0 && rc%div == 0 {
			mdc = append(mdc, div)
		}
		count := 0

		if ra%div == 0 {
			ra = ra / div
		} else {
			count += 1
		}

		if rb%div == 0 {
			rb = rb / div
		} else {
			count += 1
		}

		if rc%div == 0 {
			rc = rc / div
		} else {
			count += 1
		}

		if count == 3 {
			count = 0
			div++
		}
	}

	resMdc := 1
	for _, v := range mdc {
		resMdc *= v
	}
	return resMdc == 1
}

func main() {
	var a, b, c int
	for {
		_, err := fmt.Scanf("%d %d %d\n", &a, &b, &c)

		if err == io.EOF {
			break
		}

		a, b, c = desc(a, b, c)

		tp, primitiva := triplaPitagoricaPrimitiva(a, b, c)

		if tp && primitiva {
			fmt.Printf("tripla pitagorica primitiva\n")
		} else if tp && !primitiva {
			fmt.Printf("tripla pitagorica\n")

		} else {
			fmt.Printf("tripla\n")
		}
	}
}
