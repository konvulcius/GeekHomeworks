package main

import (
	"fmt"
	"net/http"
)

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("https://yandex.ru")
	}()
	go func() {
		responses <- request("https://vk.com")
	}()
	go func() {
		responses <- request("https://youtube.com")
	}()
	return <-responses
}

func request(hostname string) string {
	url, _ := http.Get(hostname)
	return url.Request.URL.Host
}

func main() {
	fmt.Println(mirroredQuery(), "is the fastest!")
}
