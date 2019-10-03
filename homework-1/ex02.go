package main

import (
	"fmt"
	"math"
)

func main() {
	var	a float64
	var b float64
	var c float64

	fmt.Println("Turn in legs of triangle")
	fmt.Scan(&a, &b)
	c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Printf("Hypotenuse = %.2f\n", c)
	fmt.Printf("Perimetr = %.2f\n", a + b + c)
	fmt.Printf("Square = %.2f\n", (a * b) / 2)
}
