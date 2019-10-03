package main

import (
	"fmt"
	"math"
)

func main() {
	var value float64
	var percent float64

	fmt.Println("Enter the value and percentage of Deposit")
	fmt.Scan(&value, &percent)
	if (percent <= 12) {
	fmt.Printf("Your amount in 5 years:\n%.2f\n", value * math.Pow(percent / 100 + 1, 5))
	} else {
	fmt.Printf("Your amount in 5 years:\n%.2f\n", value * math.Pow((12 + (percent - 12) * 0.65) / 100 + 1, 5))
	}
}
