package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)

	if b > a {
		a, b = b, a
	}
	if c > a {
		a, c = c, a
	}
	squareA := a * a
	squareB := b * b
	squareC := c * c

	if a >= b+c {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		if squareA == squareB+squareC {
			fmt.Println("TRIANGULO RETANGULO")
		}
		if squareA > squareB+squareC {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}
		if squareA < squareB+squareC {
			fmt.Println("TRIANGULO ACUTANGULO")
		}
	}
	if a == b && b == c && c == a {
		fmt.Println("TRIANGULO EQUILATERO")
	} else if a == b || a == c || b == c {
		fmt.Println("TRIANGULO ISOSCELES")
	}

}
