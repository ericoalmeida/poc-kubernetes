package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	USER_NAME := os.Getenv("USER_NAME")

	fmt.Fprintf(w, "<h1> Hello %s!!! <h1>", USER_NAME)
}
