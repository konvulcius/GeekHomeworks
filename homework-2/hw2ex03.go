package main

import "fmt"

func fibonacci(index int) {

	var a uint64 = 0
	var b uint64 = 1

	for i := 0; i < index; i++ {
		fmt.Println(a)
		a = b - a
		b = a + b
	}
}

func main() {
	fibonacci(100)
}
