package main

import "fmt"

var x []string

func Push(str string) {
	x = append(x, str)
}

func Pop() string {
	if len(x) == 0 {
		return ""
	}
	popElem := x[0]
	x = x[1:]
	return popElem
}

func main() {
	ar := []string{"Hello", "my", "dear", "friend"}

	for i := range ar {
		Push(ar[i])
	}
	fmt.Println(Pop() + " " + Pop())
}
