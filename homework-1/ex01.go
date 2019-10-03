package main

import "fmt"

func main() {
	const	dollar float64 = 65
	var		value float64
	
	fmt.Println("Turn in value in rubbles")
	fmt.Scan(&value)
	fmt.Printf("You have %.2f$\n", value / dollar)
}
