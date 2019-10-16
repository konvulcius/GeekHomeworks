package main

import (
	"fmt"
)

type firstElem interface {
	WriteFirstElem()
}

type someString struct {
	Elem  string
	Elem2 string
}

func (s someString) WriteFirstElem() {
	fmt.Println(s.Elem)
}

type someInt struct {
	Elem  int
	Elem2 int
}

func (i someInt) WriteFirstElem() {
	fmt.Println(i.Elem)
}

func main() {
	digit := someInt{
		12,
		36,
	}
	word := someString{
		"Dadaya",
		"Freddy",
	}

	firstElem.WriteFirstElem(digit)
	firstElem.WriteFirstElem(word)
}
