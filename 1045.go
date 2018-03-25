package urionline

import (
	"fmt"
)

func main() {
	var a, b, c float32

	fmt.Scanf("%f %f %f", &a, &b, &c)

	// ordering vars
	if a < b {
		a, b = b, a
	}

	if b < c {
		b, c = c, b
	}

	if a < b {
		a, b = b, a
	}

	// is triangule
	if a >= b+c {
		fmt.Printf("NAO FORMA TRIANGULO\n")
	} else {
		if a*a == b*b+c*c {
			fmt.Printf("TRIANGULO RETANGULO\n")
		} else if a*a > b*b+c*c {
			fmt.Printf("TRIANGULO OBTUSANGULO\n")
		} else if a*a < b*b+c*c {
			fmt.Printf("TRIANGULO ACUTANGULO\n")
		}

		if a == b && b == c {
			fmt.Printf("TRIANGULO EQUILATERO\n")
		} else if a == b || a == c || b == c {
			fmt.Printf("TRIANGULO ISOSCELES\n")
		}
	}

}
