// Bhaskara's Formula Beecrowd-1036
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scanln(&a, &b, &c)
	x := (b*b - 4*a*c)
	if x < 0 || a == 0 {
		fmt.Println("Impossivel calcular")
	} else {
		r1 := (-b + math.Sqrt(x)) / (2 * a)
		r2 := (-b - math.Sqrt(x)) / (2 * a)
		fmt.Printf("R1 = %.5f\n", r1)
		fmt.Printf("R2 = %.5f\n", r2)
	}
}
