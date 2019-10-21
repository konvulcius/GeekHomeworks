
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"io"
	"strings"
)

func main() {
	words := `Hello*"world, and other galaxy"
I*hope*you*enjoys*"your,life"*"very
well*"
Bye*`

	solve := csv.NewReader(strings.NewReader(words))
	solve.Comma = '*'

	for {
		record, err := solve.Read()
		if err == io.EOF {
			log.Fatalln(err)
			break
		} else {
			fmt.Println(record)
		}
	}
}