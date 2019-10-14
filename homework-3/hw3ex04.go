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
	addressBook["Vsevolod"] = append(addressBook["Vsevolod"], 89125674444)

	b, err := json.Marshal(addressBook)
	if err != nil {
		log.Println(err)
		return
	}
	readed, err2 := ioutil.ReadFile("death_note.txt")
	if err2 != nil {
		log.Println(err)
		return
	}
	for diff := range b {
		if b[diff] != readed[diff] {
			ioutil.WriteFile("death_note.txt", b, 0777)
			break
		}
	}
}
