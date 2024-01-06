package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	var numbers []int
	numbers = append(numbers, a)
	numbers = append(numbers, b)
	numbers = append(numbers, c)

	sort.Ints(numbers)

	for _, num := range numbers {
		fmt.Println(num)
	}

	fmt.Println()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
