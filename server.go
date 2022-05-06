package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Hello world!!! </h1>"))
}
