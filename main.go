package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("starting server PORT:8080")

	http.HandleFunc("/favicon.ico", favicon)

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func favicon(w http.ResponseWriter, r *http.Request) {}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Path)
	fmt.Fprintf(w, "Hello world!\n\nVersion: 1.0.0\n\nService: Hello world!")
}