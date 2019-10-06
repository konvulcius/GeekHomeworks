package main

import (
	"fmt"
	"strconv"
)

func inputInt(msg string) (value string) {
	fmt.Println(msg)
	fmt.Scan(&value)
	return
}

func main() {

	var g int64
	var err error

	for {
		g, err = strconv.ParseInt(inputInt("Enter integer number"), 10, 32)
		if err == nil{
			break
		}
	}
	if g % 3 == 0 {
	fmt.Println("Divided by 3")
	} else {
	fmt.Println("Not divisible by 3")
	}
}
