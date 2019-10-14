package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {

	addressBook := make(map[string][]int)

	addressBook["Fedor"] = []int{89123560374}
	addressBook["Elena"] = []int{89563401217}
	addressBook["Elena"] = append(addressBook["Elena"], 83532407809)
	addressBook["Dmitriy"] = []int{89034523123}
	addressBook["Vsevolod"] = []int{909236}
	addressBook["Vsevolod"] = append(addressBook["Vsevolod"], 89125678902)

	b, err := json.Marshal(addressBook)
	if err != nil {
		log.Println(err)
		return
	}
	ioutil.WriteFile("death_note.txt", b, 0777)
}
