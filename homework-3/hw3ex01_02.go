package main

import "fmt"

type Machine struct {
	Brande      string
	Year        uint32
	TrunkValue  float64
	ValueEngine float64
}

func main() {
	Q7 := Machine{"Audi", 2007, 2.075, 3.6}
	Cayenne := Machine{Brande: "Porsche", Year: 2012, TrunkValue: 0.645, ValueEngine: 3}

	fmt.Printf("%s Q7. Year of manifacture - %d. Value of engine - %.1f liters. Value of trunk - %.3f liters\n\n", Q7.Brande, Q7.Year, Q7.ValueEngine, Q7.TrunkValue)
	fmt.Printf("%#v\n", Cayenne)
}
