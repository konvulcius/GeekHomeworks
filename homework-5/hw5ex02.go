package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
	// getting struct of file
    file, err := os.Open("hw5ex01.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // getting size of file
    stat, err := file.Stat()
    if err != nil {
		log.Fatal(err)
    }

    // reading file
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
		log.Fatal(err)
    }

    fmt.Println(string(bs))
}