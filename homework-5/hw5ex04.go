package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()

	file1, err := os.Open(flag.Arg(0))
	if err != nil {
        log.Fatal(err)
    }
	defer file1.Close()

	file2, err := os.Create(flag.Arg(1))
	if err != nil {
        log.Fatal(err)
    }
	defer file2.Close()

	stat, err := file1.Stat()
    if err != nil {
		log.Fatal(err)
	}
	
	buf := make([]byte, stat.Size())
    _, err = file1.Read(buf)
    if err != nil {
		log.Fatal(err)
	}
	
	file2.Write(buf)
	fmt.Println(string(buf))
}

