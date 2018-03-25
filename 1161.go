package urionline

import (
	"fmt"
	"io"
)

func fatorial(count int) int {
	if count == 0 {
		return 1
	}

	return count * fatorial(count-1)
}

/**
 * There was some error, it was returning 70% wrong answers.
 * I think it's because the fmt.Scanf() needs the \n, because my solution
 * in C language was accepted with the same solution.
 */
func main() {
	for {
		var m, n int
		_, err := fmt.Scanf("%d %d\n", &m, &n)

		if err == io.EOF {
			break
		}

		soma := fatorial(m) + fatorial(n)
		fmt.Printf("%d\n", uint64(soma))
	}
}
