package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wStart sync.WaitGroup
	var wFinish sync.WaitGroup
	goStart := make(chan int, 3)
	var cars = []string{
		"datsun",
		"maserati",
		"lada",
	}
	for _, auto := range cars {
		wFinish.Add(1)
		wStart.Add(1)
		go func(auto string) {
			go func() {
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				defer wStart.Done()
			}()
			<-goStart
			time.Sleep(time.Duration(rand.Intn(8000)) * time.Millisecond)
			fmt.Println(auto, "it's time", time.Now().Format("15:04:05.000"))
			defer wFinish.Done()
		}(auto)
	}
	wStart.Wait()
	fmt.Println("Start at", time.Now().Format("15:04:05.000"))
	goStart <- 1
	goStart <- 1
	goStart <- 1
	wFinish.Wait()
}
