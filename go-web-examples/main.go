package main

import (
	"fmt"
	"net/http"
)

func main() {
	httpServer()
	helloWorld()
}

func helloWorld() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe("localhost:9000", nil)
}

func httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the website!")
	})

	fs := http.FileServer(http.Dir("./static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":9000", nil)
}
