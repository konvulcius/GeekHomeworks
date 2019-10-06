package main

import (
	"fmt"
)

func main() {

	var array [1000]int
	var new_array []int

	for index := 0; index < len(array); index++ {
		array[index] = index + 2
	}
	for i, p := range array {
		if p != 0 {
			for index := i; index < len(array); index++ {
				if p != array[index] && array[index] % p == 0 {
					array[index] = 0
				}
			}
		}
		if array[i] != 0 && len(new_array) < 101{
			new_array = append(new_array, array[i])
		}
	}
	fmt.Println(new_array)
}
