package main

import "fmt"

func fibonacci(index int) {

	var a uint64 = 1
	var b uint64 = 2

	fmt.Printf("0\n1\n")
	for i := 0; i < index; i++ {
		a = b - a
		fmt.Println(a)
		b = a + b
	}
}

func main() {
	fibonacci(100)
}
