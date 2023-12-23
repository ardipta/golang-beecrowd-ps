package main

import (
	"fmt"
)

func main() {
	var x, y int
	var z float64

	fmt.Scanln(&x, &y)

	switch x {
	case 1:
		z = 4.00 * float64(y)
		break
	case 2:
		z = 4.50 * float64(y)
		break
	case 3:
		z = 5.00 * float64(y)
		break
	case 4:
		z = 2.00 * float64(y)
		break
	case 5:
		z = 1.50 * float64(y)
		break
	}

	fmt.Printf("Total: R$ %.2f\n", z)
}
