package main

import (
	"fmt"
	"net/http"
)

func main() {
	router()
	// httpServer()
	// helloWorld()
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

	// for files to be served. the url path of the directory needs to be stripped
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":9000", nil)
}

// Extracting variables from the url path
func router() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the library!")
	})

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		title := r.PathValue("title")
		page := r.PathValue("page")

		fmt.Fprintf(w, "You've requested the book: %s and page: %s", title, page)
	})
	http.ListenAndServe("localhost:6182", r)
}
