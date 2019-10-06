package main

import "fmt"

func fibonacci(index int) {

	first := []int{0}
	second := []int{1}

	for i := 0; i < index; i++ {
		for val := len(first) - 1; val >= 0; val-- {
			fmt.Printf("%d", first[val])
		}
		fmt.Println("")
		apply := make([]int, len(second))
		copy(apply, second)
		for index, _ := range second {
			if index < len(first) {
				second[index] = first[index] + second[index]
			}
			if second[index] >= 10 {
				if index + 1 == len(second) {
					second[index] %= 10
					second = append(second, 1)
					break
				} else {
					second[index + 1] += second[index] / 10
					second[index] %= 10
				}
			}
		}
		first = apply
	}
}

func main() {
	fibonacci(100)
}
