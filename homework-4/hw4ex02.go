package main

import (
	"fmt"
	"sort"
)

type SomePeople struct {
	Name   string
	Number int
}

type PhonesBook []SomePeople

type Sort interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (phone PhonesBook) Len() int {
	return len(phone)
}

func (phone PhonesBook) Less(i, j int) bool {
	return phone[i].Name < phone[j].Name
}

func (phone PhonesBook) Swap(i, j int) {
	phone[i], phone[j] = phone[j], phone[i]
}

func (phone SomePeople) String() string {
	return fmt.Sprintf("(%v number: %v)", phone.Name, phone.Number)
}

func main() {
	Moscow := PhonesBook{
		{"Yakov", 89120348690},
		{"German", 89723642123},
		{"Lolita", 88005553535},
		{"Ludvik", 89567890232},
		{"Faust", 89783490712},
		{"Angelina", 89002345123},
	}

	sort.Sort(PhonesBook(Moscow))
	fmt.Println(Moscow)
}
