package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, r.URL.Query().Get("name"))
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/Users/acthulhu/study/geekbrains/go/src/GoCourse/homework-6", fs)
	http.HandleFunc("/", hello)

	http.ListenAndServe(":4242", nil)
}
